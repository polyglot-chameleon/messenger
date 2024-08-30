package chat

import (
	"io"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

type server struct {
	conns []*websocket.Conn
}

var Server *server

func init() {
	Server = &server{}
	http.Handle("/", websocket.Handler(Server.listen))
	go http.ListenAndServe(":8080", nil)
}

func (s *server) listen(ws *websocket.Conn) {
	log.Printf("New Connection: %v", ws.RemoteAddr())
	s.conns = append(s.conns, ws)

	buf := make([]byte, 1024)

	for {
		nBytes, err := ws.Read(buf)

		if err != nil {
			log.Print(err)
			if err == io.EOF {
				break
			}

			continue
		}

		update(string(buf[:nBytes]), true)
	}
}

func (s *server) Send(msg []byte, connIdx int) {
	s.conns[connIdx].Write(msg)
}
