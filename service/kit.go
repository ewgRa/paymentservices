package service

import (
	"net/http"
	"encoding/json"
)

func KitJsonDecodeFunc(r *http.Request, request interface{}) (interface {}, error) {
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return nil, err
	}

	return request, nil
}

func KitJsonEncodeFunc(w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}