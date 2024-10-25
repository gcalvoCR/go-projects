package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gcalvocr/go-stripe/internal/cards"
)

type stripePayload struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message,omitempty"`
	Content string `json:"content,omitempty"`
	ID      int    `json:"id,omitempty"`
}

func (app *application) GetPaymentIntent(w http.ResponseWriter, r *http.Request) {
	// Object to store information
	var payload stripePayload

	// Process to populate the object
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		app.errorLog.Println(err)
		return // this should handled better
	}

	// Converting a string to integer
	amount, err := strconv.Atoi(payload.Amount)
	if err != nil {
		app.errorLog.Println(err)
		return // this should handled better
	}

	// Card object to store data
	card := cards.Card{
		Secret:    app.config.stripe.secret,
		Key:       app.config.stripe.key,
		Currencty: payload.Currency,
	}

	// var to store if the payment ok went OK
	okay := true

	// Execute payment intent
	pi, mssg, err := card.Charge(payload.Currency, amount)
	if err != nil {
		okay = false
	}

	// Condition if payment intent went OK
	if okay {
		out, err := json.MarshalIndent(pi, "", "  ")
		if err != nil {
			app.errorLog.Println(err)
			return // this should handled better
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(out)
	} else {
		j := jsonResponse{
			OK:      false,
			Message: mssg,
			Content: "",
		}

		out, err := json.MarshalIndent(j, "", "  ")
		if err != nil {
			app.errorLog.Println(err)
		}

		w.Header().Add("Content-Type", "application/json")
		w.Write(out)
	}

}
