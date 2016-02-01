package ogone

import (
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
	"net/http"
	"../service"
)

type AliasAcceptResponseStruct struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't define JSON marshaling
}

func aliasAccept(request interface{}) (string, error) {
	// FIXME XXX: implementation
	m := request.(map[string]interface{})
	return m["S"].(string), nil
}

func makeAliasAcceptEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		v, err := aliasAccept(request)
		if err != nil {
			return AliasAcceptResponseStruct{v, err.Error()}, nil
		}
		return AliasAcceptResponseStruct{v, ""}, nil
	}
}

func decodeAliasAcceptRequest(r *http.Request) (interface{}, error) {
	var request interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func NewAliasAcceptHandler(ctx context.Context) *httptransport.Server {
	return httptransport.NewServer(
		ctx,
		makeAliasAcceptEndpoint(),
		decodeAliasAcceptRequest,
		service.EncodeResponse,
	)
}
