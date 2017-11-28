package field

import (
	"datagenerator/config"
	"errors"
	"math"
	"math/rand"
)

type floatGenerator struct {
	baseGenerator
	lbound float64
	ubound float64
}

func (fg *floatGenerator) Generate() interface{} {
	return rand.Float64()*(fg.ubound-fg.lbound) + fg.lbound
}

func (fg *floatGenerator) UniqueCount() int {
	//TODO: is not actual count, unsolved for now
	return math.MaxInt64
}

func NewFloatGenerator(c *config.Column) (*floatGenerator, error) {
	lbound := -math.MaxFloat32

	if c.LowBound != nil {
		lbound = *c.LowBound
	}

	ubound := math.MaxFloat32

	if c.UpBound != nil {
		ubound = *c.UpBound
	}

	if ubound < lbound {
		return nil, errors.New("wrong boundary range for column " + c.Name)
	}

	return &floatGenerator{
			baseGenerator: baseGenerator{
				typeName:   c.Type,
				columnName: c.Name},
			lbound: lbound,
			ubound: ubound},
		nil
}
