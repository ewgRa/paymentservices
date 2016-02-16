package service
/*
import (
	"testing"
	"fmt"
	"strings"
)

type TestResponse struct {
	V   string `json:"v"`
}

type TestRequest struct {
	Paper string    `json:"paper"`
	Count int    `json:"count"`
}

func TestJsonTransportDecode(t *testing.T) {
	transport := &JsonTransport{
		CreateRequest: func() interface{} { return &TestRequest{} },
	}

	r, _ := transport.Decode(strings.NewReader(`{"paper": "A4", "count": 5}`))

	fmt.Println(r.(*TestRequest).Count)
}*/