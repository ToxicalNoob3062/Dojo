package main

import (
	"context"
	"fmt"
	"log"
	"log-service/cmd/data"
	"log-service/logs"
	"net"

	"google.golang.org/grpc"
)

type LogServer struct {
	logs.UnimplementedLogServiceServer
	Models data.Models
}

func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
	input := req.GetLogEntry()

	// Create a new log entry
	logEntry := data.LogEntry{
		Name: input.Name,
		Data: input.Data,
	}
	err := l.Models.LogEntry.Insert(logEntry)
	if err != nil {
		res := &logs.LogResponse{
			Result: "failure",
		}
		return res, err
	}

	res := &logs.LogResponse{
		Result: "success",
	}
	return res, nil
}

func (app *Config) gRPCListen() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		log.Fatalf("Error in listening: %v", err)
	}
	s := grpc.NewServer()

	logs.RegisterLogServiceServer(s, &LogServer{
		Models: app.Models,
	})

	log.Printf("Starting gRPC service on port %s", grpcPort)

	if err := s.Serve(listen); err != nil {
		log.Fatalf("Error in serving: %v", err)
	}
}
