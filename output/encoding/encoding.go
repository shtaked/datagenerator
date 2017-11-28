package encoding

//Interface for serializing single row of data
type Marshaller interface {
	Marshal(map[string]interface{}) ([]byte, error)
}
