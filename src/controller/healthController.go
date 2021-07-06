package controller

import (
	"log"
	"net/http"
)

func CheckHealth(response http.ResponseWriter, request *http.Request) {
	log.Println("Handling /health request")

	response.WriteHeader(http.StatusOK)
}
