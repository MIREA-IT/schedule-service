package utils

import (
	"encoding/json"
	"net/http"
)

func Encode(response http.ResponseWriter, data interface{}) {
	response.Header().Add("Content-Type", "application/json")
	json.NewEncoder(response).Encode(data)
}

func Decode(response *http.Response, data interface{}) {
	json.NewDecoder(response.Body).Decode(data)
}
