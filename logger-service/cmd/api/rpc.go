package main

import (
	"fmt"
	"golang.org/x/net/context"
	"log"
	"log-service/cmd/data"
	"net"
	"net/rpc"
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

	*resp = "Logged payload via rpc!!"
	return nil
}

func (app *Config) rpcListen() error {
	log.Println("Starting rpc service on port", rpcPort)
	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", rpcPort))
	if err != nil {
		println("Error in listening: ", err)
		return err
	}
	defer listen.Close()

	for {
		rpcConn, err := listen.Accept()
		if err != nil {
			println("Error in accepting connection: ", err)
			continue
		}
		go rpc.ServeConn(rpcConn)
	}
}
