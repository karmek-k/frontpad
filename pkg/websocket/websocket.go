package websocket

import (
	"encoding/json"
	"fmt"

	"github.com/karmek-k/frontpad/pkg/db/models"
	"github.com/karmek-k/frontpad/pkg/websocket/messages"
	"gopkg.in/olahol/melody.v1"
)

func routeMessages(s *melody.Session, messageBytes []byte) {
	var msg *models.Message

	err := json.Unmarshal(messageBytes, msg)
	if err != nil {
		s.CloseWithMsg([]byte("bad message format"))

		return
	}

	switch msg.Type {
	case models.CHAT:
		break  // implement later
	case models.CODE_MARKUP:
		messages.HandleCodeMarkupMessage(s, msg)
	case models.CODE_STYLING:
		// css
	case models.CODE_SCRIPT:
		// js
	case models.USER_CONNECT:
		// user connect
		fmt.Printf("Hello %s\n", msg.Content)
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
