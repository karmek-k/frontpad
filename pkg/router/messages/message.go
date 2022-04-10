package messages

type Message struct {
	Type    MessageType `json:"type"`
	Content string      `json:"content"`
}