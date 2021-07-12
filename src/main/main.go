package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"mireait.github.io/mirea-schedule-bot/src/api"
	"mireait.github.io/mirea-schedule-bot/src/server"
	"mireait.github.io/mirea-schedule-bot/src/service"
	"net"
	"os"
	"strconv"
)

const (
	serverPostEnvVariableName = "SERVER_PORT"
	defaultPort               = 50051
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	port := fmt.Sprintf(":%d", resolveServerPort())

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen tcp on port %s: %s", port, err)
	}

	grpcServer := grpc.NewServer()

	scheduleServer := &server.ScheduleServer{
		ScheduleService: service.NewScheduleService("./resources/ИИТ_3к_20-21_весна.xlsx"),
	}

	api.RegisterScheduleServiceServer(grpcServer, scheduleServer)

	log.Printf("Handilg requests on %s\n", listener.Addr())

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("Failed to serve: ", err)
	}
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
