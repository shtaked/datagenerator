package column

import (
	"datagenerator/config"
	"errors"
	"sync"
)

//Represents single row of data
type row struct {
	columns []interface{}
}

//Checks if the row is already filled
func (r *row) isFilled() bool {
	for _, c := range r.columns {
		if c == nil {
			return false
		}
	}

	return true
}

//Interface of aggregation for column generators
type Aggregator interface {
	GetValue(colname string, i int) interface{}
}

//Basic aggregator implementation
//it stores columns names and their order
//as well as data of all generated fields
//It puts a row to output channel as soon as a row is generated
type aggregator struct {
	columns   []string
	colnames  map[string]int
	colfilled []int
	data      []row
	out       chan<- map[string]interface{}
	lock      sync.Mutex
}

func NewAggregator(s *config.Settings, out chan<- map[string]interface{}) *aggregator {
	cs := make([]string, len(s.Columns))
	cns := make(map[string]int)
	for i, c := range s.Columns {
		cs[i] = c.Name
		cns[c.Name] = i
	}

	f := make([]int, len(s.Columns))

	d := make([]row, s.RowsCount)
	for i := range d {
		d[i].columns = make([]interface{}, len(s.Columns))
	}

	return &aggregator{
		columns:   cs,
		colnames:  cns,
		colfilled: f,
		data:      d,
		out:       out}
}

// Gets filled row and returns it as map with column name as key
func (a *aggregator) getOutMap(rowidx int) map[string]interface{} {
	out := make(map[string]interface{})

	for i, c := range a.data[rowidx].columns {
		out[a.columns[i]] = c
	}

	return out
}

// Starts goroutine with specified column generation and listens for result
// It uses mutex to prevent r/w accessing of data from multiple goroutines
func (a *aggregator) StartAndListen(cg Generator) error {
	idx, ok := a.colnames[cg.GetName()]
	if !ok {
		return errors.New("column name was not found for aggregator: " + cg.GetName())
	}

	ch := make(chan interface{}, 1000)

	//start
	go cg.Generate(len(a.data), ch)

	// and listen
	a.colfilled[idx] = 0
	for item := range ch {
		a.lock.Lock()

		a.data[a.colfilled[idx]].columns[idx] = item
		a.colfilled[idx]++

		currow := a.colfilled[idx] - 1

		if a.data[currow].isFilled() {
			a.out <- a.getOutMap(currow)
		}

		a.lock.Unlock()
	}

	return nil
}

//Implementation of Aggregator interface
//Method is used to get data for another column generators (uniquefor, uniquewithin)
//That's why it also uses mutex to prevent data races
func (a *aggregator) GetValue(colname string, i int) (v interface{}) {
	colidx, _ := a.colnames[colname]

	a.lock.Lock()
	v = a.data[i].columns[colidx]
	a.lock.Unlock()
	return
}
