package controller

import (
	"database/sql"
	"fmt"
	"log"

	"fyne.io/fyne/v2/data/binding"
	_ "github.com/mattn/go-sqlite3"
)

type MessageResource struct {
	ChatID      int
	TextContent string
	Incoming    bool
}

type messageController struct {
	db *sql.DB
}

var MessageController *messageController
var data binding.UntypedList

func init() {
	MessageController = &messageController{}
	MessageController.db, _ = sql.Open("sqlite3", "db/messenger.db")
}

func (mc *messageController) WriteMessage(newMsg MessageResource) {

	var incoming int
	if newMsg.Incoming {
		incoming = 1
	} else {
		incoming = 0
	}

	_, err := mc.db.Exec(fmt.Sprintf("INSERT INTO messages(chat_id, incoming, text_content) VALUES (1, %v, '%s');", incoming, newMsg.TextContent))
	if err != nil {
		log.Fatal(err)
	}
}

func (mc *messageController) ReadMessages(chat_id int) []any {
	rows, _ := mc.db.Query("SELECT text_content, incoming FROM messages;")
	defer rows.Close()

	var messages []any
	message := MessageResource{TextContent: "", Incoming: false}

	for rows.Next() {
		rows.Scan(&message.TextContent, &message.Incoming)
		messages = append(messages, message)
	}

	return messages
}
