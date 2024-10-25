package main

import (
	"log"
	"net/http"
)

func (app *Config) SendMail(w http.ResponseWriter, r *http.Request) {
	type mailMessage struct {
		From    string `json:"from"`
		To      string `json:"to"`
		Subject string `json:"subject"`
		Message string `json:"message"`
	}

	var requestPayload mailMessage

	log.Println("reading json")
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJson(w, err)
	}

	msg := Message{
		From:    requestPayload.From,
		To:      requestPayload.To,
		Subject: requestPayload.Subject,
		Data:    requestPayload.Message,
	}

	log.Println("Sending Message", msg)
	err = app.Mailer.SendSMTPMessage(msg)
	if err != nil {
		app.errorJson(w, err)
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Email sent to " + requestPayload.To,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}
