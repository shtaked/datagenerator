package column

import (
	"datagenerator/config"
	"datagenerator/generators/field"
	"math/rand"
	"testing"
	"time"
)

type intValueGetter struct {
	v []int
}

func newIntValueGetter(c int, ubound int) *intValueGetter {
	vg := intValueGetter{make([]int, c)}

	for i := 0; i < c; i++ {
		vg.v[i] = rand.Intn(ubound + 1)
	}

	return &vg
}

func (vg intValueGetter) Get(i int) int {
	return vg.v[i]
}

func (vg intValueGetter) GetValue(colname string, i int) (v interface{}) {
	return vg.Get(i)
}

func TestUniqueWithinGenerator_Generate(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	lbound := 0.0
	upbound := 1000.0
	col := config.Column{
		Name:     "int_column1",
		Type:     "int",
		LowBound: &lbound,
		UpBound:  &upbound}

	rcount := 1000

	fg, _ := field.NewIntGenerator(&col)

	vg := newIntValueGetter(rcount, 5)

	cg := NewUniqueWithinGenerator("int_column2", vg, fg)

	ch := make(chan interface{}, 1000)
	go cg.Generate(rcount, ch)

	out := make(map[int][]int, rcount)
	for i := 0; i < rcount; i++ {
		out[i] = make([]int, 0)
	}

	for i := 0; i < rcount; i++ {
		item := <-ch
		v := item.(int)
		curdepv := vg.Get(i)
		prevdepvs := out[v]
		found := false
		for _, prevdepv := range prevdepvs {
			if curdepv == prevdepv {
				found = true
				break
			}
		}

		if found {
			t.Error("Expected only unique within column values")
			break
		} else {
			out[v] = append(out[v], curdepv)
		}
	}
}
