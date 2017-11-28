package tsv

import (
	"testing"
	//"bytes"
)

func TestMarshaller_GetHeader(t *testing.T) {
	m, err := NewMarshaller([]string{"First Name", "Last Name", "Age"})
	o, err := m.GetHeader()
	if err != nil {
		t.Errorf("Expected no error for normal input but found \"%v\"", err)
		return
	}

	if string(o) != "First Name\tLast Name\tAge\n" {
		t.Errorf("Expected \"First Name\tLast Name\tAge\n\" output instead of \"%v\"", string(o))
	}

	m, err = NewMarshaller([]string{"First Name"})
	o, err = m.GetHeader()
	if err != nil {
		t.Errorf("Expected no error for single-column input but found \"%v\"", err)
		return
	}

	if string(o) != "First Name\n" {
		t.Errorf("Expected \"First Name\n\" output instead of \"%v\"", string(o))
	}

	m, err = NewMarshaller([]string{})
	if err == nil {
		t.Error("Expected error for empty list of columns")
	}
}

func TestMarshaller_Marshal(t *testing.T) {
	m, err := NewMarshaller([]string{"First Name", "Last Name", "Age"})

	in := map[string]interface{}{
		"Last Name": "Brown", "First Name": "John",
	}

	o, err := m.Marshal(in)
	if err == nil {
		t.Error("Expected error for less columns input")
	}

	in = map[string]interface{}{
		"Last Name": "Brown", "First Name": "John", "Age": 55, "Sex": "m",
	}

	o, err = m.Marshal(in)
	if err == nil {
		t.Error("Expected error for more columns input")
		return
	}

	in = map[string]interface{}{}

	o, err = m.Marshal(in)
	if err == nil {
		t.Error("Expected error for empty columns input")
		return
	}

	in = map[string]interface{}{
		"Last": "Brown", "First": "John", "age": 55,
	}

	o, err = m.Marshal(in)
	if err == nil {
		t.Error("Expected error for wrong columns input")
		return
	}

	in = map[string]interface{}{
		"Last Name": "Brown", "First Name": "John", "Age": 55,
	}

	o, err = m.Marshal(in)
	if err != nil {
		t.Errorf("Expected no error for normal input but found \"%v\"", err)
		return
	}

	if string(o) != "John\tBrown\t55\n" {
		t.Errorf("Expected \"Brown\tJohn\t55\n\" output instead of \"%v\"", string(o))
	}
}
