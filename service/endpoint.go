package service

type Endpoint interface {
	Response(r interface{}) (interface{}, error)
}