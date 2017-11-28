package column

import "datagenerator/generators/field"

// Guarantees uniqueness within certain value of another column
// It gets dependant column value through Aggregator interface each time it generates new value
type uniqueWithinGenerator struct {
	baseGenerator
	agr       Aggregator
	depcolumn string
	cvalues   map[interface{}][]interface{}
}

func NewUniqueWithinGenerator(depcolumn string, agr Aggregator, fg field.Generator) *uniqueWithinGenerator {
	return &uniqueWithinGenerator{
		baseGenerator: baseGenerator{fg: fg},
		agr:           agr,
		depcolumn:     depcolumn,
		cvalues:       make(map[interface{}][]interface{})}
}

func (ug *uniqueWithinGenerator) getDepValue(i int) (r interface{}) {
	for r == nil {
		r = ug.agr.GetValue(ug.depcolumn, i)
	}

	return
}

func (ug *uniqueWithinGenerator) Generate(c int, ch chan<- interface{}) {
	for i := 0; i < c; i++ {
		depval := ug.getDepValue(i)

		var result interface{}

		for {
			result = ug.fg.Generate()
			prevvals, ok := ug.cvalues[result]
			if !ok {
				//put dep value to the map
				ug.cvalues[result] = make([]interface{}, 0)
				ug.cvalues[result] = append(ug.cvalues[result], depval)
				break
			} else {
				found := false
				for _, prevval := range prevvals {
					if prevval == depval {
						found = true
						break
					}
				}

				//break only if new value is not found among old values
				if !found {
					ug.cvalues[result] = append(ug.cvalues[result], depval)
					break
				}
			}
		}

		ch <- result
	}

	close(ch)
}
