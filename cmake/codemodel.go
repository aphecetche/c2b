package cmake

import (
	"encoding/json"
	"fmt"
)

func UnmarshalMessageCodeModel(data []byte) (Message, error) {
	var r MessageCodeModel
	err := json.Unmarshal(data, &r)
	if r.TypeString != CodeModelMsg.String() {
		return nil, fmt.Errorf("JSON do not contain compute codemodel but '%s' message", r.TypeString)
	}
	return &r, err
}

func (r *MessageCodeModel) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MessageCodeModel struct {
	TypeString string `json:"type"`
}

func (m *MessageCodeModel) Type() MessageType {
	return CodeModelMsg
}

func (m *MessageCodeModel) String() string {
	return m.Type().String()
}

func NewMessageCodeModel() Message {
	return &MessageCodeModel{
		TypeString: CodeModelMsg.String(),
	}
}
