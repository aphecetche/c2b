package cmake

import (
	"encoding/json"
	"fmt"
)

func UnmarshalMessageHandshake(data []byte) (Message, error) {
	var r MessageHandshake
	err := json.Unmarshal(data, &r)
	if r.TypeString != HandshakeMsg.String() {
		return nil, fmt.Errorf("JSON do not contain handshake message but '%s' message", r.TypeString)
	}
	return &r, err
}

func (r *MessageHandshake) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MessageHandshake struct {
	Cookie          string          `json:"cookie,omitempty"`
	TypeString      string          `json:"type"`
	ProtocolVersion ProtocolVersion `json:"protocolVersion,omitempty"`
	SourceDirectory string          `json:"sourceDirectory,omitempty"`
	BuildDirectory  string          `json:"buildDirectory"`
	Generator       string          `json:"generator,omitempty"`
	ExtraGenerator  string          `json:"extraGenerator,omitempty"`
	Platform        string          `json:"platform,omitempty"`
	Toolset         string          `json:"toolset,omitempty"`
}

type ProtocolVersion struct {
	Major int64 `json:"major"`
	Minor int64 `json:"minor,omitempty"`
}

func (m *MessageHandshake) Type() MessageType {
	return HandshakeMsg
}

func (m *MessageHandshake) String() string {
	return fmt.Sprintf("source:%s build:%s generator:%s",
		m.SourceDirectory, m.BuildDirectory, m.Generator)
}

func NewMessageHandshake(major int64, source string, build string, generator string) Message {
	return &MessageHandshake{
		TypeString:      HandshakeMsg.String(),
		ProtocolVersion: ProtocolVersion{major, 0},
		SourceDirectory: source,
		BuildDirectory:  build,
		Generator:       generator,
	}
}
