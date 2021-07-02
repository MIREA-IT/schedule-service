package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
	"therealmone.github.io/mirea-schedule-bot/dto"
	"therealmone.github.io/mirea-schedule-bot/utils"
)

var (
	server *httptest.Server
)

func TestDefaultPort(t *testing.T) {
	port := resolveServerPort()

	if port != defaultPort {
		t.Fatal("Expected port: ", defaultPort)
	}
}

func TestCustomAddress(t *testing.T) {
	expectedPort := 3000

	os.Setenv(serverPostEnvVariableName, strconv.Itoa(expectedPort))

	port := resolveServerPort()

	if port != expectedPort {
		t.Fatal("Expected port: ", expectedPort)
	}
}

func TestHealth(t *testing.T) {
	setup()

	callGet("/api/health", http.StatusOK)

	defer server.Close()
}

func TestGetSchedules(t *testing.T) {
	setup()

	res := callGet("/api/schedules", http.StatusOK)
	var schedules []dto.Schedule

	utils.Decode(res, &schedules)

	if len(schedules) == 0 {
		t.Fatal("Schedules are empty")
	}

	for i := range schedules {
		schedule := schedules[i]

		if schedule.Name == "" {
			t.Fatal("Schedule name is empty")
		}
	}

	defer server.Close()
}

func setup() {
	server = createServer()
}

func createServer() *httptest.Server {
	srv := httptest.NewServer(handlers())

	return srv
}

func callGet(url string, expectedStatus int) *http.Response {
	return call(http.Get, url, expectedStatus)
}

func call(method func(url string) (*http.Response, error), url string, expectedStatus int) *http.Response {
	res, err := method(fmt.Sprintf("%s%s", server.URL, url))

	if err != nil {
		log.Fatal("Error has been occurred: ", err)
	}

	if res.StatusCode != http.StatusOK {
		log.Fatal("Invalid status: %i, expected: %i", res.StatusCode, expectedStatus)
	}

	return res
}
