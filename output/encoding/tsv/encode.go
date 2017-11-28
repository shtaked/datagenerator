package tsv

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
)

//Tab separated value serializer
//Stores column names for header
type marshaller struct {
	columns []string
}

func NewMarshaller(columns []string) (*marshaller, error) {
	if len(columns) == 0 {
		return nil, errors.New("empty list of columns for marshaller")
	}

	return &marshaller{columns}, nil
}

func (m *marshaller) GetHeader() (out []byte, err error) {
	var b bytes.Buffer
	w := csv.NewWriter(bufio.NewWriter(&b))
	w.Comma = '\t'

	err = w.Write(m.columns)
	w.Flush()

	out = b.Bytes()

	return
}

func (m *marshaller) Marshal(row map[string]interface{}) (out []byte, err error) {
	var b bytes.Buffer
	w := csv.NewWriter(bufio.NewWriter(&b))
	w.Comma = '\t'

	if len(m.columns) != len(row) {
		return nil, errors.New("columns count mismatch for TSV marshalling")
	}

	buf := make([]string, len(m.columns))
	for i, column := range m.columns {
		item, ok := row[column]
		if !ok {
			return nil, errors.New("wrong column for TSV marshalling: " + column)
		}

		buf[i] = fmt.Sprintf("%v", item)
	}

	err = w.Write(buf)
	w.Flush()

	out = b.Bytes()

	return
}
