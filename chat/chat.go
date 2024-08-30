package chat

import (
	"messenger/controller"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var Chat *fyne.Container
var data binding.UntypedList

func init() {
	messages := controller.MessageController.ReadMessages(1)
	data = binding.NewUntypedList()
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
			update(textArea.Text, false)
			textArea.SetText("")
		},
	}

	Chat = container.NewBorder(nil, form, nil, nil, list)
}

func update(msg string, incoming bool) {
	newMsg := controller.MessageResource{ChatID: 1, TextContent: msg, Incoming: incoming}

	data.Append(newMsg)
	controller.MessageController.WriteMessage(newMsg)
	Server.Send([]byte(newMsg.TextContent), 0)
}
