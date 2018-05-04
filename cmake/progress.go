package cmake

import (
	"encoding/json"
	"fmt"
)

func UnmarshalMessageProgress(data []byte) (Message, error) {
	var r MessageProgress
	err := json.Unmarshal(data, &r)
	if r.TypeString != ProgressMsg.String() {
		return nil, fmt.Errorf("JSON do not contain progress message but '%s' message", r.TypeString)
	}
	return &r, err
}

func (r *MessageProgress) String() string {
	return fmt.Sprintf("Progress of %s %5.0f %%", r.ProgressMessage, 100.0*float64(r.ProgressCurrent)/float64(r.ProgressMaximum))
}
func (r *MessageProgress) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MessageProgress struct {
	Cookie          string `json:"cookie"`
	InReplyTo       string `json:"inReplyTo"`
	ProgressCurrent int64  `json:"progressCurrent"`
	ProgressMaximum int64  `json:"progressMaximum"`
	ProgressMessage string `json:"progressMessage"`
	ProgressMinimum int64  `json:"progressMinimum"`
	TypeString      string `json:"type"`
}

func (m *MessageProgress) Type() MessageType {
	return ProgressMsg
}
