package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "8080"

type Config struct{}

func main() {
	app := Config{}

	log.Printf("Starting the broker service on port %s\n", webPort)

	//define http Server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	//start the server
	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}