package models

type MessageType int

const (
	CHAT MessageType = iota
	CODE_MARKUP
	CODE_STYLING
	CODE_SCRIPT
)

type Message struct {
	Type    MessageType `json:"type"`
	Content string      `json:"content"`
}