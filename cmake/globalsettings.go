package cmake

import (
	"encoding/json"
	"fmt"
)

func UnmarshalMessageGlobalSettings(data []byte) (Message, error) {
	var r MessageGlobalSettings
	err := json.Unmarshal(data, &r)
	if r.TypeString != GlobalSettingsMsg.String() {
		return nil, fmt.Errorf("JSON do not contain reply globalSettings but '%s' message", r.TypeString)
	}
	return &r, err
}

func (r *MessageGlobalSettings) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MessageGlobalSettings struct {
	TypeString string `json:"type"`
	Cookie     string `json:"cookie,omitempty"`
}

func (m *MessageGlobalSettings) Type() MessageType {
	return GlobalSettingsMsg
}
func (m *MessageGlobalSettings) String() string {
	return m.Type().String()
}

func NewMessageGlobalSettings() Message {
	return &MessageGlobalSettings{
		TypeString: GlobalSettingsMsg.String(),
	}
}
