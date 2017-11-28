package tsv

import (
	"datagenerator/output/encoding"
	"datagenerator/output/encoding/tsv"
	"datagenerator/output/file"
	"io"
	"os"
)

type fileHandler struct {
	name        string
	file        *os.File
	comprwriter io.WriteCloser
	marshaller  encoding.Marshaller
}

func NewFileHandler(name string, compr file.Compressor, columns []string) (h *fileHandler, err error) {
	f, err := os.Create(name)
	if err != nil {
		return nil, err
	}

	m, err := tsv.NewMarshaller(columns)
	if err != nil {
		return
	}

	header, err := m.GetHeader()
	if err != nil {
		return
	}

	h = &fileHandler{name, f, nil, m}

	if compr != nil {
		h.comprwriter = compr.NewWriter(f)
		h.comprwriter.Write(header)
	} else {
		h.file.Write(header)
	}

	return
}

func (h *fileHandler) Write(item map[string]interface{}) (err error) {
	b, err := h.marshaller.Marshal(item)
	if err != nil {
		return
	}

	if h.comprwriter != nil {
		_, err = h.comprwriter.Write(b)
	} else {
		_, err = h.file.Write(b)
	}

	return
}

func (h *fileHandler) Close() (err error) {
	if h.comprwriter != nil {
		err = h.comprwriter.Close()
	}

	h.file.Close()
	return
}
