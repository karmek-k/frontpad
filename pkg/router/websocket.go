package router

import (
	"encoding/json"
	"fmt"

	"github.com/karmek-k/frontpad/pkg/router/messages"
	"gopkg.in/olahol/melody.v1"
)

func routeMessages(s *melody.Session, messageBytes []byte) {
	var msg *messages.Message

	err := json.Unmarshal(messageBytes, msg)
	if err != nil {
		s.Write([]byte("bad message format"))
	}

	switch msg.Type {
	case messages.CHAT:
		// chat message
	case messages.CODE_MARKUP:
		// html
	case messages.CODE_STYLING:
		// css
	case messages.CODE_SCRIPT:
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