// Chuck Norris fortune teller cli using tview.
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	app      *tview.Application
	textView *tview.TextView
	delay    = 10
)

// Payload contails the joke as a string
type Payload struct {
	Value string
}

func getJoke() string {
	result, err := http.Get("https://api.chucknorris.io/jokes/random?category=science")
	if err != nil {
		panic(err)
	}

	payloadBytes, err := io.ReadAll(result.Body)
	if err != nil {
		panic(err)
	}

	payload := &Payload{}

	err = json.Unmarshal(payloadBytes, payload)
	if err != nil {
		panic(err)
	}

	return payload.Value
}

func showJoke(joke string) {
	textView.Clear()

	timeStr := fmt.Sprintf("[red]%s", time.Now().Format(time.RFC1123))

	fmt.Fprintln(textView, timeStr)
	fmt.Fprintln(textView)
	fmt.Fprintln(textView)

	fmt.Fprintln(textView, "[white]"+joke)
}

func refreshScreen() {
	tick := time.NewTicker(time.Second * time.Duration(delay))

	for { //nolint:all // false positive for for+select loop
		select {
		case <-tick.C:
			joke := getJoke()
			showJoke(joke)
			app.Draw()
		}
	}
}

func init() {
	textView = tview.NewTextView().SetDynamicColors(true).
		SetWrap(true).
		SetWordWrap(true).
		SetTextAlign(tview.AlignCenter).
		SetTextColor(tcell.ColorLime)
}

func main() {
	app = tview.NewApplication()

	showJoke(getJoke())

	go refreshScreen()

	if err := app.SetRoot(textView, true).Run(); err != nil {
		panic(err)
	}
}
