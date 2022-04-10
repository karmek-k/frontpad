package websocket

import (
	"encoding/json"
	"fmt"

	"gopkg.in/olahol/melody.v1"
)

func routeMessages(s *melody.Session, messageBytes []byte) {
	var msg *Message

	err := json.Unmarshal(messageBytes, msg)
	if err != nil {
		s.Write([]byte("bad message format"))
	}

	switch msg.Type {
	case CHAT:
		// chat message
	case CODE_MARKUP:
		// html
	case CODE_STYLING:
		// css
	case CODE_SCRIPT:
		// js
	}
}

func CreateWebSocket() *melody.Melody {
	m := melody.New()

	m.HandleConnect(func(s *melody.Session) {
		fmt.Println("connected")
	})

	m.HandleDisconnect(func(s *melody.Session) {
		fmt.Println("disconnected")
	})

	m.HandleMessage(routeMessages)

	return m
}
