package websocket

type Message struct {
	Type    MessageType `json:"type"`
	Content string      `json:"content"`
}