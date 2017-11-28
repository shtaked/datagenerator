package config

import (
	"encoding/json"
	"os"
)

//Reflects column type from json configuration file
//Type could be one of the existing types (str, int, float, bool) or it can be a custom type
//Regexp is only used for string types and subtypes
//Bounds are used only for int and float types and subtypes
//One of uniqueness modifiers can be used at the same time
// - "unique" guarantees uniqueness for total amount of rows independently of other columns
// - "unique_within_column_value" guarantees uniqueness within certain value of another column
//e.g Col1 Col2 (unique_within_column_value=Col1)
//    1    1
//    1    2
//    2    1
//    2    2
//    1    2   <----- is not allowed
// - "unique_for_column_value" guarantees that value is unique only for certain value from another column
//e.g Col1 Col2 (unique_for_column_value=Col1)
//    1    1
//    2    2
//    1    2   <----- is not allowed
type Column struct {
	Name                    string   `json:"name"`
	Type                    string   `json:"type"`
	Regexp                  *string  `json:"regexp,omitempty"`
	LowBound                *float64 `json:"low_bound,omitempty"`
	UpBound                 *float64 `json:"up_bound,omitempty"`
	Unique                  bool     `json:"unique"`
	UniqueForColumnValue    string   `json:"unique_for_column_value"`
	UniqueWithinColumnValue string   `json:"unique_within_column_value"`
}

//Reflects custom type from json configuration file
//Parent should be one of the existing types (str, int, float, bool)
type CustomType struct {
	Name     string   `json:"name"`
	Parent   string   `json:"parent"`
	Regexp   *string  `json:"regexp,omitempty"`
	LowBound *float64 `json:"low_bound,omitempty"`
	UpBound  *float64 `json:"up_bound,omitempty"`
}

//Reflects overall json configuration structure
//Output specifies output file
//Compression is optional and for now only Gzip is supported
type Settings struct {
	Columns     []Column     `json:"columns"`
	CustomTypes []CustomType `json:"custom_types,omitempty"`
	RowsCount   int          `json:"row_count"`
	Output      string       `json:"output"`
	Compression string       `json:"compression"`
}

func Load(name string) (result Settings, err error) {
	f, err := os.Open(name)
	if err != nil {
		return
	}

	d := json.NewDecoder(f)
	err = d.Decode(&result)

	return
}
