CREATE TABLE IF NOT EXISTS chats (
    id INTEGER PRIMARY KEY,
    chat_name VARCHAR NOT NULL
);
CREATE TABLE IF NOT EXISTS messages (
    chat_id INTEGER,
    text_content TEXT VARCHAR NOT NULL,
    PRIMARY KEY (chat_id, text_content),
    FOREIGN KEY (chat_id) REFERENCES Chats (id) ON DELETE CASCADE
);