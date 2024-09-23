package main

import (
	"github.com/rivo/tview"
)

var app *tview.Application

func main() {
	app = tview.NewApplication()

	if err := app.SetRoot(nil, true).Run(); err != nil {
		panic(err)
	}
}
