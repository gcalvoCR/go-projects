package main

import (
	"encoding/json"
	"image/color"
	"net/http"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var client *http.Client

type RandomFact struct {
	Text string `json:"text"`
}

func getRandomFact() (RandomFact, error) {
	var fact RandomFact
	resp, err := client.Get("https://uselessfacts.jsph.pl/random.json?language=en")
	if err != nil {
		return RandomFact{}, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&fact); err != nil {
		return RandomFact{}, err
	}

	return fact, nil

}

func main() {
	client = &http.Client{Timeout: 10 * time.Second}

	a := app.New()
	w := a.NewWindow("Get Useless Fact")
	w.Resize(fyne.NewSize(800, 300))

	title := canvas.NewText("Get your useless facts", color.White)
	title.TextStyle = fyne.TextStyle{
		Bold:      true,
		Italic:    false,
		Monospace: false,
		Symbol:    false,
		TabWidth:  0,
	}
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 24

	factText := widget.NewLabel("")
	factText.Wrapping = fyne.TextWrapWord

	button := widget.NewButton("Get Fact", func() {
		fact, err := getRandomFact()
		if err != nil {
			dialog.ShowError(err, w)
		} else {
			factText.SetText(fact.Text)
		}
	})

	hBox := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), button, layout.NewSpacer())
	vBox := container.New(layout.NewVBoxLayout(), title, hBox, widget.NewSeparator(), factText)

	w.SetContent(vBox)
	w.ShowAndRun()
}
