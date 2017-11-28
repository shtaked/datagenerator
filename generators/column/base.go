package column

import (
	"datagenerator/generators/field"
)

//interface for generating column of certain type
//puts result to the specified channel
type Generator interface {
	GetName() string
	CanGenerate(c int) bool
	Generate(c int, ch chan<- interface{})
}

type baseGenerator struct {
	fg field.Generator
}

func (bg *baseGenerator) GetName() string {
	return bg.fg.ColumnName()
}

//by default all simple columns can generate any amount of rows
func (bg *baseGenerator) CanGenerate(c int) bool {
	return true
}
