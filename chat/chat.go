package chat

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var Chat *fyne.Container

var data = []string{"a", "string", "list"}

func init() {
	textArea := widget.NewMultiLineEntry()

	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])
		})

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Widget: textArea}},
		OnSubmit: func() { log.Println("Form submitted:", textArea.Text) },
	}

	Chat = container.NewBorder(nil, form, nil, nil, list)
}
