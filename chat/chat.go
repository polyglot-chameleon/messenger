package chat

import (
	"messenger/controller"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var Chat *fyne.Container

func init() {
	messages := controller.MessageController.ReadMessages(1)
	data := binding.BindStringList(&messages)

	textArea := widget.NewMultiLineEntry()
	list := widget.NewListWithData(data,

		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Widget: textArea}},
		OnSubmit: func() {
			data.Append(textArea.Text)
			controller.MessageController.WriteMessage(1, textArea.Text)
			textArea.SetText("")
		},
	}

	Chat = container.NewBorder(nil, form, nil, nil, list)
}
