package models

type MessageType int

const (
	CHAT MessageType = iota
	CODE_MARKUP
	CODE_STYLING
	CODE_SCRIPT
	USER_CONNECT
)

type Message struct {
	Type      MessageType `json:"type"`
	SessionId string      `json:"sessionId"`
	Content   string      `json:"content"`
}