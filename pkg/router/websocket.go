package router

import (
	"fmt"

	"gopkg.in/olahol/melody.v1"
)

func CreateWebSocket() *melody.Melody {
	m := melody.New()

	m.HandleConnect(func(s *melody.Session) {
		fmt.Println("connected")
	})

	m.HandleDisconnect(func(s *melody.Session) {
		fmt.Println("disconnected")
	})

	return m
}