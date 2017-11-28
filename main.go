package main

import (
	"datagenerator/config"
	"datagenerator/generators"
	"datagenerator/generators/column"
	"datagenerator/output/file"
	"datagenerator/output/handler/tsv"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	if len(os.Args) < 2 {
		checkErr(errors.New("config file name is not specified"))
	}

	start := time.Now()

	rand.Seed(time.Now().UnixNano())

	fmt.Println("Data generator is started")

	f := os.Args[1]
	s, err := config.Load(f)
	checkErr(err)

	fmt.Println("Initializing field generators")
	fgs, err := generators.GetFieldGenerators(&s)
	checkErr(err)

	fmt.Println("Creating output file")

	cs := make([]string, len(s.Columns))
	for i, c := range s.Columns {
		cs[i] = c.Name
	}

	var compr file.Compressor
	if s.Compression == "gzip" {
		compr = file.GzipCompressor{}
		s.Output += ".gz"
	}
	h, err := tsv.NewFileHandler(s.Output, compr, cs)
	//h, err := tsv.NewConsoleHandler(cs)
	checkErr(err)
	defer h.Close()

	fmt.Println("Generating result file " + s.Output)

	out := make(chan map[string]interface{}, 1000)
	agr := column.NewAggregator(&s, out)

	for i, fg := range fgs {
		var cg column.Generator
		if s.Columns[i].Unique {
			cg = column.NewUniqueGenerator(fg)
		} else if s.Columns[i].UniqueForColumnValue != "" {
			//TODO: check if column exitsts
			cg = column.NewUniqueForGenerator(s.Columns[i].UniqueForColumnValue, agr, fg)
		} else if s.Columns[i].UniqueWithinColumnValue != "" {
			//TODO: check if column exitsts
			cg = column.NewUniqueWithinGenerator(s.Columns[i].UniqueWithinColumnValue, agr, fg)
		} else {
			cg = column.NewSimpleGenerator(fg)
		}

		if !cg.CanGenerate(s.RowsCount) {
			checkErr(errors.New(
				fmt.Sprintf("too many values (%v) to generate values for column %v",
					s.RowsCount, cg.GetName())))
		}

		go agr.StartAndListen(cg)
	}

	for i := 0; i < s.RowsCount; i++ {
		item := <-out
		err = h.Write(item)
		checkErr(err)
	}

	fmt.Println("Finished")

	end := time.Now()
	diff := end.Sub(start)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
