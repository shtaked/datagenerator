package column

import "datagenerator/generators/field"

//The simpliest type of column generator
//Corresponds to column without any uniqueness modifiers
type simpleGenerator struct {
	baseGenerator
}

func NewSimpleGenerator(fg field.Generator) *simpleGenerator {
	return &simpleGenerator{
		baseGenerator: baseGenerator{
			fg: fg}}
}

func (sg *simpleGenerator) Generate(c int, ch chan<- interface{}) {
	for i := 0; i < c; i++ {
		ch <- sg.fg.Generate()
	}

	close(ch)
}
