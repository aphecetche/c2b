package cmake

import (
	"encoding/json"
	"fmt"
)

func UnmarshalMessageHello(data []byte) (Message, error) {
	var r MessageHello
	err := json.Unmarshal(data, &r)
	if r.TypeString != HelloMsg.String() {
		return nil, fmt.Errorf("JSON do not contain hello message but '%s' message", r.TypeString)
	}
	return &r, err
}

func (r *MessageHello) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MessageHello struct {
	SupportedProtocolVersion []SupportedProtocolVersion `json:"supportedProtocolVersions"`
	TypeString               string                     `json:"type"`
}

func (m *MessageHello) Type() MessageType {
	return HelloMsg
}

func (m *MessageHello) String() string {
	return m.Type().String()
}

type SupportedProtocolVersion struct {
	IsExperimental bool  `json:"isExperimental"`
	Major          int64 `json:"major"`
	Minor          int64 `json:"minor"`
}
