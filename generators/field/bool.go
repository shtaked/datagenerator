package field

import (
	"datagenerator/config"
	"math/rand"
)

type boolGenerator struct {
	baseGenerator
}

func (bg *boolGenerator) Generate() interface{} {
	return rand.Float32() < 0.5
}

func (bg *boolGenerator) UniqueCount() int {
	return 2
}

func NewBoolGenerator(c *config.Column) (*boolGenerator, error) {
	return &boolGenerator{
			baseGenerator: baseGenerator{
				typeName:   c.Type,
				columnName: c.Name}},
		nil
}
