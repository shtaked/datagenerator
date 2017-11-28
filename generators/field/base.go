package field

//interface for generating a single random field of certain type
type Generator interface {
	TypeName() string
	ColumnName() string
	Generate() interface{}
	UniqueCount() int
}

type baseGenerator struct {
	typeName   string
	columnName string
}

func (bcg *baseGenerator) TypeName() string {
	return bcg.typeName
}

func (bcg *baseGenerator) ColumnName() string {
	return bcg.columnName
}
