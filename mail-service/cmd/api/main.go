package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

// everything needed to run the app or server!
type Config struct {
	Mailer Mail
}

const webPort = "8083"

func main() {

	app := Config{
		Mailer: createMail(),
	}

	log.Println("Starting mail service on port", webPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}

// Creates Mail object dynamically from environment variables
// Mail object is details about the mail server and sender
func createMail() Mail {
	// convert string to int
	port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
	m := Mail{
		Domain:      os.Getenv("MAIL_DOMAIN"),
		Host:        os.Getenv("MAIL_HOST"),
		Port:        port,
		Username:    os.Getenv("MAIL_USERNAME"),
		Password:    os.Getenv("MAIL_PASSWORD"),
		Encryption:  os.Getenv("MAIL_ENCRYPTION"),
		FromName:    os.Getenv("FROM_NAME"),
		FromAddress: os.Getenv("FROM_ADDRESS"),
	}

	return m
}
