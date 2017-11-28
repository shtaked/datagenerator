package column

import (
	"datagenerator/config"
	"datagenerator/generators/field"
	"math/rand"
	"testing"
	"time"
)

type boolValueGetter struct {
	v []bool
}

func newBoolValueGetter(c int) *boolValueGetter {
	agr := boolValueGetter{make([]bool, c)}

	for i := 0; i < c; i++ {
		agr.v[i] = rand.Float32() < 0.5
	}

	return &agr
}

func (vg boolValueGetter) Get(i int) bool {
	return vg.v[i]
}

func (vg boolValueGetter) GetValue(colname string, i int) (v interface{}) {
	return vg.Get(i)
}

func TestUniqueForGenerator_Generate(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	lbound := 0.0
	upbound := 10.0
	col := config.Column{
		Name:     "int_column",
		Type:     "int",
		LowBound: &lbound,
		UpBound:  &upbound}

	rcount := 1000

	fg, _ := field.NewIntGenerator(&col)

	vg := newBoolValueGetter(rcount)

	cg := NewUniqueForGenerator("bool_column", vg, fg)

	ch := make(chan interface{}, 1000)
	go cg.Generate(rcount, ch)

	out := make(map[int]bool, rcount)
	for i := 0; i < rcount; i++ {
		item := <-ch
		v := item.(int)
		curb := vg.Get(i)
		prevb, ok := out[v]
		if ok {
			if curb != prevb {
				t.Error("Expected only unique for column values")
				break
			}
		} else {
			out[v] = curb
		}
	}
}
