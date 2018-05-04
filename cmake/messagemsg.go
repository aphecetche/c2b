package cmake

import (
	"encoding/json"
	"fmt"
)

func UnmarshalMessageMessage(data []byte) (Message, error) {
	var r MessageMessage
	err := json.Unmarshal(data, &r)
	if r.TypeString != MessageMsg.String() {
		return nil, fmt.Errorf("JSON do not contain reply message but '%s' message", r.TypeString)
	}
	return &r, err
}

func (m *MessageMessage) String() string {
	return m.Message
}

func (r *MessageMessage) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MessageMessage struct {
	Message    string `json:"message"`
	TypeString string `json:"type"`
	Cookie     string `json:"cookie,omitempty"`
}

func (m *MessageMessage) Type() MessageType {
	return MessageMsg
}
