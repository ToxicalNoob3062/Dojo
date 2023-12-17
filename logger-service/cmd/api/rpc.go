package main

import (
	"golang.org/x/net/context"
	"log"
	"log-service/cmd/data"
	"time"
)

type RPCServer struct{}

type RpcPayload struct {
	Name string
	Data string
}

func (r *RPCServer) LogInfo(payload *RpcPayload, resp *string) error {
	collection := client.Database("logs").Collection("logs")
	_, err := collection.InsertOne(context.TODO(), data.LogEntry{
		Name:      payload.Name,
		Data:      payload.Data,
		CreatedAt: time.Now(),
	})
	if err != nil {
		log.Panicln("Error in inserting log entry: ", err)
		return err
	}

	*resp = "Processed Payload via rpc!!" + payload.Name
	return nil
}
