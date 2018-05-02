package csp

import "encoding/json"

func UnmarshalHello(data []byte) (Hello, error) {
	var r Hello
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Hello) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Hello struct {
	SupportedProtocolVersions []SupportedProtocolVersion `json:"supportedProtocolVersions"`
	Type                      string                     `json:"type"`
}

type SupportedProtocolVersion struct {
	IsExperimental bool  `json:"isExperimental"`
	Major          int64 `json:"major"`
	Minor          int64 `json:"minor"`
}
