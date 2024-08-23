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
	data := binding.NewUntypedList()
	data.Set(messages)

	textArea := widget.NewMultiLineEntry()

	list := widget.NewListWithData(data,

		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			message, _ := i.(binding.Untyped).Get()
			msg := message.(controller.MessageResource)

			if msg.Incoming {
				o.(*widget.Label).Alignment = fyne.TextAlignLeading
			} else {
				o.(*widget.Label).Alignment = fyne.TextAlignTrailing
			}

			text_content := binding.NewString()
			text_content.Set(msg.TextContent)

			o.(*widget.Label).Bind(text_content)
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
