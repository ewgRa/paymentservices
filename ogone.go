package main

import (
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()
	svc := paymentOgoneService{}

	aliasAcceptHandler := httptransport.NewServer(
		ctx,
		makeAliasAcceptEndpoint(svc),
		decodeAliasAcceptRequest,
		encodeResponse,
	)

	aliasDirectLinkHandler := httptransport.NewServer(
		ctx,
		makeAliasDirectLinkEndpoint(svc),
		decodeAliasDirectLinkRequest,
		encodeResponse,
	)

	http.Handle("/alias_accept", aliasAcceptHandler)
	http.Handle("/alias_direct_link", aliasDirectLinkHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type PaymentOgoneService interface {
    AliasAccept(interface{}) (string, error)
    AliasDirectLink(int, int, string, string) (string, error)
}

type paymentOgoneService struct{}

func (paymentOgoneService) AliasAccept(request interface{}) (string, error) {
	// FIXME XXX: implementation
	m := request.(map[string]interface{})
	return m["S"].(string), nil
}

func (paymentOgoneService) AliasDirectLink(orderId int, amount int, currency string, alias string) (string, error) {
	if "" == alias {
		return "", errors.New("Alias required")
	}

	// FIXME XXX: implementation
	return "OK", nil
}

type aliasAcceptResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't define JSON marshaling
}

type aliasDirectLinkRequest struct {
	orderId  int    `json:"orderId"`
	amount   int    `json:","`
	currency string `json:"currency"`
	alias    string `json:"alias"`
}

type aliasDirectLinkResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't define JSON marshaling
}

func makeAliasAcceptEndpoint(svc PaymentOgoneService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		v, err := svc.AliasAccept(request)
		if err != nil {
			return aliasAcceptResponse{v, err.Error()}, nil
		}
		return aliasAcceptResponse{v, ""}, nil
	}
}

func makeAliasDirectLinkEndpoint(svc PaymentOgoneService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(aliasDirectLinkRequest)
		v, err := svc.AliasDirectLink(req.orderId, req.amount, req.currency, req.alias)
		if err != nil {
			return aliasDirectLinkResponse{v, err.Error()}, nil
		}
		return aliasDirectLinkResponse{v, ""}, nil
	}
}

func decodeAliasAcceptRequest(r *http.Request) (interface{}, error) {
	var request interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeAliasDirectLinkRequest(r *http.Request) (interface{}, error) {
	var request aliasDirectLinkRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
