package messages

type MessageType int

const (
	CHAT MessageType = iota
	CODE_MARKUP
	CODE_STYLING
	CODE_SCRIPT
)