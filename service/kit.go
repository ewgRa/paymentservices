package service

import (
	"encoding/json"
	"net/http"
)

// KitJSONDecodeFunc decode json http request to endpoint request
func KitJSONDecodeFunc(r *http.Request, request interface{}) (interface{}, error) {
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return nil, err
	}

	return request, nil
}

// KitJSONEncodeFunc encode endpoint response to json
func KitJSONEncodeFunc(w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
