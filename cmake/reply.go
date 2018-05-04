package cmake

import (
	"encoding/json"
	"fmt"
)

func UnmarshalMessageReply(data []byte) (Message, error) {
	var r MessageReply
	err := json.Unmarshal(data, &r)
	if r.TypeString != ReplyMsg.String() {
		return nil, fmt.Errorf("JSON do not contain reply message but '%s' message", r.TypeString)
	}
	return &r, err
}

func (r *MessageReply) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MessageReply struct {
	InReplyTo  string `json:"inReplyTo"`
	TypeString string `json:"type"`
	Cookie     string `json:"cookie,omitempty"`
}

func (m *MessageReply) Type() MessageType {
	return ReplyMsg
}
