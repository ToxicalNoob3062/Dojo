package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"log-service/cmd/data"
	"net/http"
	"net/rpc"
	"time"
)

const (
	webPort  = "8082"
	rpcPort  = "5082"
	grpcPort = "6082"
	mongoUrl = "mongodb://mongo:27017"
)

var client *mongo.Client

type Config struct {
	Models data.Models
}

func main() {
	//connect to mongo
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Panic(err)
	}

	client = mongoClient

	//create a context inorder to disconnect
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	//close connection
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	//create a new config
	app := Config{
		Models: data.New(client),
	}

	//start rpc server in a goroutine
	err = rpc.Register(new(RPCServer))
	if err != nil {
		log.Println("Error in registering rpc server: ", err.Error())
	}
	go app.rpcListen()

	//start grpc server in a goroutine
	go app.gRPCListen()

	//start the server
	app.serve()
}

func (app *Config) serve() {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	log.Println("Starting log service on port", webPort)
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func connectToMongo() (*mongo.Client, error) {
	//create connection options
	clientOptions := options.Client().ApplyURI(mongoUrl)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

	//connect to mongo
	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	return c, nil
}
