#!/bin/zsh

protoc --go_out=. --go-grpc_out=. api/proto/scheduleService.proto