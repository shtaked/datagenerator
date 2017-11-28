package handler

//Output handler interface
type Handler interface {
	Write(map[string]interface{}) error
	Close() error
}
