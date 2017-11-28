package tsv

import (
	"datagenerator/output/encoding"
	"datagenerator/output/encoding/tsv"
	"os"
)

type consoleHandler struct {
	marshaller encoding.Marshaller
}

func NewConsoleHandler(columns []string) (h *consoleHandler, err error) {
	m, err := tsv.NewMarshaller(columns)
	if err != nil {
		return
	}

	header, err := m.GetHeader()
	if err != nil {
		return
	}

	h = &consoleHandler{m}
	os.Stdout.Write(header)

	return
}

func (h *consoleHandler) Write(item map[string]interface{}) (err error) {
	b, err := h.marshaller.Marshal(item)
	if err != nil {
		return
	}

	os.Stdout.Write(b)

	return
}

func (h *consoleHandler) Close() (err error) {
	return
}
