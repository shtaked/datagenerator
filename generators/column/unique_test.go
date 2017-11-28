package column

import (
	"datagenerator/config"
	"datagenerator/generators/field"
	"testing"
)

func TestUniqueGenerator_CanGenerate(t *testing.T) {
	fg, _ := field.NewBoolGenerator(
		&config.Column{
			Name: "test_column",
			Type: "bool"})
	cg := NewUniqueGenerator(fg)
	if cg.CanGenerate(3) {
		t.Error("Expected that unique bool column generator can not generate more then 2 values")
	}
}

func TestUniqueGenerator_Generate(t *testing.T) {
	lbound := 0.0
	upbound := 1000.0
	fg, _ := field.NewIntGenerator(
		&config.Column{
			Name:     "test_column",
			Type:     "int",
			LowBound: &lbound,
			UpBound:  &upbound})
	cg := NewUniqueGenerator(fg)

	ch := make(chan interface{}, 1000)
	go cg.Generate(1000, ch)

	out := make(map[int]bool, 1000)
	for i := 0; i < 1000; i++ {
		item := <-ch
		v := item.(int)
		_, ok := out[v]
		if ok {
			t.Error("Expected only unique values")
			break
		}

		out[v] = true
	}
}
