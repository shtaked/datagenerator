package field

import (
	"datagenerator/config"
	"errors"
	"math"
	"math/rand"
)

type intGenerator struct {
	baseGenerator
	lbound int
	ubound int
}

func (ig *intGenerator) Generate() interface{} {
	if ig.ubound == ig.lbound {
		return ig.ubound
	} else {
		return rand.Intn(ig.ubound+1-ig.lbound) + ig.lbound
	}
}

func (ig *intGenerator) UniqueCount() int {
	return ig.ubound - ig.lbound + 1
}

func NewIntGenerator(c *config.Column) (*intGenerator, error) {
	lbound := math.MinInt64

	if c.LowBound != nil {
		lbound = int(*c.LowBound)
	}

	ubound := math.MaxInt64

	if c.UpBound != nil {
		ubound = int(*c.UpBound)
	}

	if ubound < lbound {
		return nil, errors.New("wrong boundary range for column " + c.Name)
	}

	return &intGenerator{
			baseGenerator: baseGenerator{
				typeName:   c.Type,
				columnName: c.Name},
			lbound: lbound,
			ubound: ubound},
		nil
}
