package column

import "datagenerator/generators/field"

//Generates data where each value is unique
type uniqueGenerator struct {
	baseGenerator
	values map[interface{}]bool
}

func NewUniqueGenerator(fg field.Generator) *uniqueGenerator {
	return &uniqueGenerator{
		baseGenerator: baseGenerator{fg: fg},
		values:        make(map[interface{}]bool)}
}

func (ug *uniqueGenerator) CanGenerate(c int) bool {
	return ug.fg.UniqueCount() >= c
}

func (ug *uniqueGenerator) Generate(c int, ch chan<- interface{}) {
	for i := 0; i < c; i++ {
		var result interface{}
		for ok := true; ok; _, ok = ug.values[result] {
			result = ug.fg.Generate()
		}
		ug.values[result] = true
		ch <- result
	}

	close(ch)
}
