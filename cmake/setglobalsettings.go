package cmake

import (
	"encoding/json"
	"fmt"
)

func UnmarshalMessageSetGlobalSettings(data []byte) (Message, error) {
	var r MessageSetGlobalSettings
	err := json.Unmarshal(data, &r)
	if r.TypeString != SetGlobalSettingsMsg.String() {
		return nil, fmt.Errorf("JSON do not contain setGlobalSettings message but '%s' message", r.TypeString)
	}
	return &r, err
}

func (r *MessageSetGlobalSettings) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MessageSetGlobalSettings struct {
	TypeString      string `json:"type"`
	Cookie          string `json:"cookie,omitempty"`
	DebugOutput     bool   `json:"debugoutput,omitempty"`
	SourceDirectory string `json:"sourceDirectory,omitempty"`
	BuildDirectory  string `json:"buildDirectory"`
	Generator       string `json:"generator,omitempty"`
	ExtraGenerator  string `json:"extraGenerator,omitempty"`
}

func (m *MessageSetGlobalSettings) Type() MessageType {
	return SetGlobalSettingsMsg
}

func NewMessageSetGlobalSettings(source, build, generator string) Message {
	return &MessageSetGlobalSettings{
		TypeString:      SetGlobalSettingsMsg.String(),
		SourceDirectory: source,
		BuildDirectory:  build,
		Generator:       generator,
	}
}
