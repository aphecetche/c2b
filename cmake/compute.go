package cmake

import (
	"encoding/json"
	"fmt"
)

func UnmarshalMessageCompute(data []byte) (Message, error) {
	var r MessageCompute
	err := json.Unmarshal(data, &r)
	if r.TypeString != ComputeMsg.String() {
		return nil, fmt.Errorf("JSON do not contain compute message but '%s' message", r.TypeString)
	}
	return &r, err
}

func (r *MessageCompute) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MessageCompute struct {
	TypeString string `json:"type"`
}

func (m *MessageCompute) Type() MessageType {
	return ComputeMsg
}

func (m *MessageCompute) String() string {
	return m.Type().String()
}

func NewMessageCompute() Message {
	return &MessageCompute{
		TypeString: ComputeMsg.String(),
	}
}
