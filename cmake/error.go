package cmake

import (
	"encoding/json"
	"fmt"
)

func UnmarshalMessageError(data []byte) (Message, error) {
	var r MessageError
	err := json.Unmarshal(data, &r)
	if r.TypeString != ErrorMsg.String() {
		return nil, fmt.Errorf("JSON do not contain reply error but '%s' message", r.TypeString)
	}
	return &r, err
}

func (r *MessageError) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func (r *MessageError) String() string {
	return fmt.Sprintf("ERROR from %s : %s", r.InReplyTo, r.ErrorMessage)
}

type MessageError struct {
	InReplyTo    string `json:"inReplyTo"`
	ErrorMessage string `json:"errorMessage"`
	TypeString   string `json:"type"`
	Cookie       string `json:"cookie,omitempty"`
}

func (m *MessageError) Type() MessageType {
	return ErrorMsg
}
