package main

import (
	"authentication/cmd/data"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJson(w, r, &requestPayload)

	if err != nil {
		app.errorJson(w, err, http.StatusBadRequest)
		return
	}

	user, err := app.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		log.Println("Error getting user by email", err)
		app.errorJson(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		log.Println("Error in matching password!!", err)
		app.errorJson(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	//log reQuest
	err = app.logRequest("auth", fmt.Sprintf("Logged in user %s", user.Email))
	if err != nil {
		app.errorJson(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", user.Email),
		Data:    user,
	}

	app.writeJson(w, http.StatusAccepted, payload)
}

func (app *Config) logRequest(name, data string) error {
	var entry = struct {
		Name string `json:"name"`
		Data string `json:"data"`
	}{
		Name: name,
		Data: data,
	}

	jsonData, _ := json.MarshalIndent(entry, "", "\t")
	loggerServiceUrl := "http://project_logger-service_1:8082/log"

	request, err := http.NewRequest("POST", loggerServiceUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		return err
	}
	return nil
}

func (app *Config) Register(w http.ResponseWriter, _ *http.Request) {
	user := data.User{
		Email:     "admin@example.com",
		Password:  "verysecret",
		Active:    false,
		FirstName: "Admin",
		LastName:  "User",
	}

	_, exists := app.Models.User.GetByEmail(user.Email)
	if exists == nil {
		app.errorJson(w, errors.New("user already exists"))
		return
	}

	id, err := app.Models.User.Insert(user)

	if err != nil {
		println("Error inserting user", err.Error())
		app.errorJson(w, err)
		return
	}

	response := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("User created with id %d", id),
	}

	app.writeJson(w, http.StatusAccepted, response)
}
