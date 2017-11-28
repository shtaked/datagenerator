package field

import (
	"datagenerator/config"
	"errors"
	"github.com/lucasjones/reggen"
)

func NewYearGenerator(c *config.Column) (*intGenerator, error) {
	const miny int = 0
	const maxy int = 2100

	lbound := miny

	if c.LowBound != nil {
		lbound = int(*c.LowBound)
	}

	ubound := maxy

	if c.UpBound != nil {
		ubound = int(*c.UpBound)
	}

	if ubound < lbound {
		return nil, errors.New("wrong boundary range for column " + c.Name)
	}

	if ubound < miny || ubound > maxy {
		return nil, errors.New("wrong year boundary for column " + c.Name)
	}

	return &intGenerator{
			baseGenerator: baseGenerator{
				typeName:   c.Type,
				columnName: c.Name},
			lbound: lbound, ubound: ubound},
		nil
}

func NewEmailGenerator(c *config.Column) (*stringGenerator, error) {
	rg, err := reggen.NewGenerator("^[a-z]{5,10}@[a-z]{5,10}\\.(com|net|org)$}")

	return &stringGenerator{
			baseGenerator: baseGenerator{
				typeName:   c.Type,
				columnName: c.Name},
			rg: rg},
		err
}
