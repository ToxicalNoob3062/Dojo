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

	//define http Server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	//start the server
	log.Printf("Starting the broker service on port %s\n", webPort)
	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}
