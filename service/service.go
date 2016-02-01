package service

import "net/http"
import "encoding/json"

func EncodeResponse(w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
