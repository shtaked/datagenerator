package field

import (
	"datagenerator/config"
	"math"
	"testing"
)

func newFloatColumn(lbound float64, ubound float64) *config.Column {
	return &config.Column{
		Name:     "test_column",
		Type:     "int",
		LowBound: &lbound,
		UpBound:  &ubound}
}

func TestNewFloatGenerator(t *testing.T) {
	g, err := NewFloatGenerator(newFloatColumn(2.0, 1.0))
	if err == nil {
		t.Error("Expected error for lbound > ubound")
	}

	g, err = NewFloatGenerator(newFloatColumn(1.0, 2.0))
	if err != nil {
		t.Errorf("Expected no error for ubound > lbound but found \"%v\"", err)
	}
	if g == nil {
		t.Error("Expected valid generator for ubound > lbound")
	}

	g, err = NewFloatGenerator(newFloatColumn(1.0, 1.0))
	if err != nil {
		t.Errorf("Expected no error for ubound == lbound but found \"%v\"", err)
	}
	if g == nil {
		t.Error("Expected valid generator for ubound == lbound")
	}

	g, err = NewFloatGenerator(newFloatColumn(-2.0, -1.0))
	if err != nil {
		t.Errorf("Expected no error for negative ubound > lbound but found \"%v\"", err)
	}
	if g == nil {
		t.Error("Expected valid generator for negative ubound > lbound")
	}

	g, err = NewFloatGenerator(newFloatColumn(-math.MaxFloat64, math.MaxFloat64))
	if err != nil {
		t.Errorf("Expected no error for extreme ubound > lbound but found \"%v\"", err)
	}
	if g == nil {
		t.Error("Expected valid generator for extreme ubound > lbound")
	}

	g, err = NewFloatGenerator(
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

func testFloatRange(t *testing.T, lbound float64, ubound float64) {
	g, _ := NewFloatGenerator(newFloatColumn(lbound, ubound))
	for i := 0; i < 100; i++ {
		v := g.Generate()
		fv, ok := v.(float64)
		if !ok {
			t.Errorf("Expected to return float64 but got %T", fv)
			return
		}
		if fv < lbound || fv > ubound {
			t.Errorf("Expected value to be in a specified range (%v, %v) but got %v",
				lbound, ubound, fv)
			return
		}
	}
}

func TestFloatGenerator_Generate(t *testing.T) {
	testFloatRange(t, 1.0, 2.0)
	testFloatRange(t, 1.0, 1.0+math.SmallestNonzeroFloat64)
	testFloatRange(t, 1.0, 1.0)
	testFloatRange(t, -2.0, -1.0)
	testFloatRange(t, -math.MaxFloat32, math.MaxFloat32)
}
