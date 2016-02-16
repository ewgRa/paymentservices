package kit_service

import (
	kithttp "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
	"net/http"
	"encoding/json"
	"../."
)

type Transport interface {
	DecodeFunc() kithttp.DecodeRequestFunc
	EncodeFunc(w http.ResponseWriter, response interface{}) error
}

type JsonTransport struct {
	CreateRequest func() interface{}
}

func NewJsonTransport(CreateRequest func() interface{}) *JsonTransport {
	return &JsonTransport{CreateRequest: CreateRequest}
}

func (t *JsonTransport) DecodeFunc() kithttp.DecodeRequestFunc {
	return func(r *http.Request) (interface{}, error) {
		request := t.CreateRequest()

		if err := json.NewDecoder(r.Body).Decode(request); err != nil {
			return nil, err
		}

		return request, nil
	}
}

func (t *JsonTransport) EncodeFunc(w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func NewHandler(ctx context.Context, ep service.Endpoint, transport Transport) *kithttp.Server {
	return kithttp.NewServer(
		ctx,
		func(ctx context.Context, request interface{}) (response interface{}, err error) {
			return ep.Response(request)
		},
		transport.DecodeFunc(),
		transport.EncodeFunc,
	)
}
