package websocket

import (
	"encoding/json"
	"fmt"

	"github.com/karmek-k/frontpad/pkg/websocket/messages"
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
		break  // implement later
	case CODE_MARKUP:
		messages.HandleCodeMarkupMessage(s, msg.Content)
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
