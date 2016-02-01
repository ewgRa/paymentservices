package ogone

import (
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
	"net/http"
	"../service"
)

type AliasDirectLinkRequestStruct struct {
	orderId  int    `json:"orderId"`
	amount   int    `json:","`
	currency string `json:"currency"`
	alias    string `json:"alias"`
}

type AliasDirectLinkResponseStruct struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't define JSON marshaling
}

func aliasDirectLink(orderId int, amount int, currency string, alias string) (string, error) {
	if "" == alias {
		return "", errors.New("Alias required")
	}

	// FIXME XXX: implementation
	return "OK", nil
}

func makeAliasDirectLinkEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AliasDirectLinkRequestStruct)
		v, err := aliasDirectLink(req.orderId, req.amount, req.currency, req.alias)
		if err != nil {
			return AliasDirectLinkResponseStruct{v, err.Error()}, nil
		}
		return AliasDirectLinkResponseStruct{v, ""}, nil
	}
}

func decodeAliasDirectLinkRequest(r *http.Request) (interface{}, error) {
	var request AliasDirectLinkRequestStruct
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func NewAliasDirectLinkHandler(ctx context.Context) *httptransport.Server {
	return httptransport.NewServer(
		ctx,
		makeAliasDirectLinkEndpoint(),
		decodeAliasDirectLinkRequest,
		service.EncodeResponse,
	)
}
