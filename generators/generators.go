package generators

import (
	"datagenerator/config"
	"datagenerator/generators/field"
	"errors"
)

//returns generator that can generate single random field by it's configuration description
func newFieldGenerator(c *config.Column) (field.Generator, error) {
	//TODO: use reflection instead to construct proper type
	switch c.Type {
	case "bool":
		return field.NewBoolGenerator(c)
	case "int":
		return field.NewIntGenerator(c)
	case "float":
		return field.NewFloatGenerator(c)
	case "string":
		return field.NewStringGenerator(c)
	case "year":
		return field.NewYearGenerator(c)
	case "email":
		return field.NewEmailGenerator(c)
	default:
		return nil, errors.New("unknown column type: " + c.Type)
	}
}

//returns array of field generators where each element corresponds to resulting column
func GetFieldGenerators(s *config.Settings) (result []field.Generator, err error) {
	//map of custom types by names
	cts := make(map[string]config.CustomType)

	for _, t := range s.CustomTypes {
		if len(t.Name) == 0 {
			return nil, errors.New("name is not specified for custom type")
		}

		if len(t.Parent) == 0 {
			return nil, errors.New("parent is not specified for custom type " + t.Name)
		}

		cts[t.Name] = t
	}

	result = make([]field.Generator, len(s.Columns))

	for i, c := range s.Columns {
		ct, ok := cts[c.Type]
		rc := c
		if ok {
			//resolving one of custom types for column
			rc.Type = ct.Parent
			if ct.Regexp != nil {
				rc.Regexp = ct.Regexp
			}

			if ct.LowBound != nil {
				rc.LowBound = ct.LowBound
			}

			if ct.UpBound != nil {
				rc.UpBound = ct.UpBound
			}
		}

		result[i], err = newFieldGenerator(&rc)

		if err != nil {
			return
		}
	}

	return
}
