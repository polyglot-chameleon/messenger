package main

import (
	"messenger/chat"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")

	w.Resize(fyne.NewSize(500, 500))
	w.SetContent(chat.Chat)
	w.ShowAndRun()
}
