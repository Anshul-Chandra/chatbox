package types

// Message holds the content of a websocket chat message
type Message struct {
	ReceiverID string `json:"receiverID"`
	Data       string `json:"data"`
}

// NewMessage returns a new instance of message type
func NewMessage(receiverID string, data string) *Message {
	return &Message{
		ReceiverID: receiverID,
		Data:       data,
	}
}

// GetReceiver returns the receiver ID associated with the message
func (m *Message) GetReceiver() string {
	if m == nil {
		return ""
	}

	return m.ReceiverID
}
