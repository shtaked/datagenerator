package field

import (
	"datagenerator/config"
	"math"
	"testing"
)

func newIntColumn(lbound int, ubound int) *config.Column {
	flbound := float64(lbound)
	fubound := float64(ubound)
	return &config.Column{
		Name:     "test_column",
		Type:     "int",
		LowBound: &flbound,
		UpBound:  &fubound}
}

func TestNewIntGenerator(t *testing.T) {
	g, err := NewIntGenerator(newIntColumn(2, 1))
	if err == nil {
		t.Error("Expected error for lbound > ubound")
	}

	g, err = NewIntGenerator(newIntColumn(1, 2))
	if err != nil {
		t.Errorf("Expected no error for ubound > lbound but found \"%v\"", err)
	}
	if g == nil {
		t.Error("Expected valid generator for ubound > lbound")
	}

	g, err = NewIntGenerator(newIntColumn(1, 1))
	if err != nil {
		t.Errorf("Expected no error for ubound == lbound but found \"%v\"", err)
	}
	if g == nil {
		t.Error("Expected valid generator for ubound == lbound")
	}

	g, err = NewIntGenerator(newIntColumn(-2, -1))
	if err != nil {
		t.Errorf("Expected no error for negative ubound > lbound but found \"%v\"", err)
	}
	if g == nil {
		t.Error("Expected valid generator for negative ubound > lbound")
	}

	g, err = NewIntGenerator(newIntColumn(math.MinInt64, math.MaxInt64))
	if err != nil {
		t.Errorf("Expected no error for extreme ubound > lbound but found \"%v\"", err)
	}
	if g == nil {
		t.Error("Expected valid generator for extreme ubound > lbound")
	}

	g, err = NewIntGenerator(
		&config.Column{
			Name: "test_column",
			Type: "int"})

	if err != nil {
		t.Errorf("Expected no error when bounds are not set but found \"%v\"", err)
	}
	if g == nil {
		t.Error("Expected valid generator when bounds are not set")
	}
}

func testIntRange(t *testing.T, lbound int, ubound int) {
	g, _ := NewIntGenerator(newIntColumn(lbound, ubound))
	for i := 0; i < 100; i++ {
		v := g.Generate()
		iv, ok := v.(int)
		if !ok {
			t.Errorf("Expected to return int but got %T", iv)
			return
		}
		if iv < lbound || iv > ubound {
			t.Errorf("Expected value to be in a specified range (%v, %v) but got %v",
				lbound, ubound, iv)
			return
		}
	}
}

func TestIntGenerator_Generate(t *testing.T) {
	testIntRange(t, 1, 2)
	testIntRange(t, 1, 1)
	testIntRange(t, -2, -1)
	testIntRange(t, math.MinInt64, math.MaxInt64)
}

func TestIntGenerator_UniqueCount(t *testing.T) {
	g, _ := NewIntGenerator(newIntColumn(1, 2))
	c := g.UniqueCount()
	if c != 2 {
		t.Errorf("Expected 2 unique values but got %v", c)
	}

	g, _ = NewIntGenerator(newIntColumn(1, 1))
	c = g.UniqueCount()
	if c != 1 {
		t.Errorf("Expected 1 unique values but got %v", c)
	}
}
