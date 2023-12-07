package main

import (
	"log-service/cmd/data"
	"net/http"
)

type JsonPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) writeLog(w http.ResponseWriter, r *http.Request) {
	//read the json into a variable
	var requestPayload JsonPayload
	_ = app.readJson(w, r, &requestPayload)

	//insert data
	event := data.LogEntry{
		Name: requestPayload.Name,
		Data: requestPayload.Data,
	}

	//insert data into mongo
	err := app.Models.LogEntry.Insert(event)
	if err != nil {
		app.errorJson(w, err)
		return
	}

	//return success
	resp := jsonResponse{
		Error:   false,
		Message: "logged!!",
	}

	app.writeJson(w, http.StatusAccepted, resp)
}
