package controller

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type MessageResource struct {
	TextContent string
	Incoming    bool
}

type messageController struct {
	db *sql.DB
}

var MessageController *messageController

func init() {
	MessageController = &messageController{}
	MessageController.db, _ = sql.Open("sqlite3", "db/messenger.db")
}

func (mc *messageController) WriteMessage(chat_id int, text_content string) {
	_, err := mc.db.Exec(fmt.Sprintf("INSERT INTO messages(chat_id, text_content) VALUES (1, '%s')", text_content))
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
