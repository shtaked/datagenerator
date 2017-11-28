package column

import "datagenerator/generators/field"

// Guarantees uniqueness within certain value of another column
// It gets dependant column value through Aggregator interface each time it generates new value
type uniqueForGenerator struct {
	baseGenerator
	agr       Aggregator
	depcolumn string
	cvalues   map[interface{}]interface{}
}

func NewUniqueForGenerator(depcolumn string, agr Aggregator, fg field.Generator) *uniqueForGenerator {
	return &uniqueForGenerator{
		baseGenerator: baseGenerator{fg: fg},
		agr:           agr,
		depcolumn:     depcolumn,
		cvalues:       make(map[interface{}]interface{})}
}

func (ug *uniqueForGenerator) getDepValue(i int) (r interface{}) {
	for r == nil {
		r = ug.agr.GetValue(ug.depcolumn, i)
	}

	return
}

func (ug *uniqueForGenerator) Generate(c int, ch chan<- interface{}) {
	for i := 0; i < c; i++ {
		depval := ug.getDepValue(i)

		var result interface{}

		for {
			result = ug.fg.Generate()
			prevval, ok := ug.cvalues[result]
			if !ok {
				//put dep value to the map
				ug.cvalues[result] = depval
				break
			} else {
				//break only if value is the same as before for the same dep value
				if prevval == depval {
					break
				}
			}
		}

		ch <- result
	}

	close(ch)
}
