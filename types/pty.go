package types

type MessgaeType string

const (
	Ping   MessgaeType = "ping"
	Pong   MessgaeType = "pong"
	Input  MessgaeType = "input"
	Output MessgaeType = "output"
	// Should be sent in the string format <rows>,<cols>
	Resize MessgaeType = "resize"
)

type SocketEvent struct {
	Type    MessgaeType `json:"type"`
	Content string      `json:"content"`
}
