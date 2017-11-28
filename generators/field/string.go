package field

import (
	"datagenerator/config"
	"errors"
	"github.com/lucasjones/reggen"
	"math"
)

type stringGenerator struct {
	baseGenerator
	rg *reggen.Generator
}

func (sg *stringGenerator) Generate() interface{} {
	return sg.rg.Generate(10)
}

func (sg *stringGenerator) UniqueCount() int {
	//TODO: is not actual count, unsolved
	return math.MaxInt64
}

func NewStringGenerator(c *config.Column) (sg *stringGenerator, err error) {
	if c.Regexp == nil || *c.Regexp == "" {
		return nil, errors.New("regexp is not specified for column " + c.Name)
	}

	rg, err := reggen.NewGenerator(*c.Regexp)
	if err != nil {
		return
	}

	sg = &stringGenerator{
		baseGenerator: baseGenerator{
			typeName:   c.Type,
			columnName: c.Name},
		rg: rg}

	return
}
