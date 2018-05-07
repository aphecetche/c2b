package cmake

import (
	"encoding/json"
	"fmt"
	"strings"
)

func UnmarshalMessageConfigure(data []byte) (Message, error) {
	var r MessageConfigure
	err := json.Unmarshal(data, &r)
	if r.TypeString != ConfigureMsg.String() {
		return nil, fmt.Errorf("JSON do not contain reply configure but '%s' message", r.TypeString)
	}
	return &r, err
}

func (r *MessageConfigure) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MessageConfigure struct {
	TypeString     string   `json:"type"`
	Cookie         string   `json:"cookie,omitempty"`
	CacheArguments []string `json:"cacheArguments"`
}

func (m *MessageConfigure) Type() MessageType {
	return ConfigureMsg
}

func (m *MessageConfigure) String() string {
	return strings.Join(m.CacheArguments, " ")
}

func NewMessageConfigure(cacheArguments []string) Message {
	return &MessageConfigure{
		TypeString:     ConfigureMsg.String(),
		CacheArguments: cacheArguments,
	}
}
