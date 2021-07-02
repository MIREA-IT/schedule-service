package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
	"therealmone.github.io/mirea-schedule-bot/controller"
)

const (
	get  = "GET"
	post = "POST"
	put  = "PUT"

	serverPostEnvVariableName = "SERVER_PORT"
	defaultPort               = 8080
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	address := fmt.Sprintf(":%d", resolveServerPort())

	log.Printf("Handilg requests on address: %s\n", address)

	log.Fatal(http.ListenAndServe(address, handlers()))
}

func resolveServerPort() int {
	envPort := os.Getenv(serverPostEnvVariableName)

	if envPort != "" {
		port, err := strconv.Atoi(envPort)

		if err != nil {
			log.Fatal("Invalid port: ", envPort)
		}

		return port
	} else {
		return defaultPort
	}
}

func handlers() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/api/health", controller.CheckHealth).Methods(get)
	router.HandleFunc("/api/schedules", controller.GetSchedules).Methods(get)

	return router
}
