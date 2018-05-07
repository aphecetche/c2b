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
	Configurations []Configuration `json:"configurations,omitempty"`
	InReplyTo      string          `json:"inReplyTo"`
	TypeString     string          `json:"type"`
	Cookie         string          `json:"cookie,omitempty"`
}

type Configuration struct {
	Name     string    `json:"name"`
	Projects []Project `json:"projects"`
}

type Project struct {
	BuildDirectory  string   `json:"buildDirectory"`
	Name            string   `json:"name"`
	SourceDirectory string   `json:"sourceDirectory"`
	Targets         []Target `json:"targets"`
}

type Target struct {
	Artifacts       []string    `json:"artifacts"`
	BuildDirectory  string      `json:"buildDirectory"`
	FileGroups      []FileGroup `json:"fileGroups"`
	FullName        string      `json:"fullName"`
	LinkerLanguage  string      `json:"linkerLanguage"`
	Name            string      `json:"name"`
	SourceDirectory string      `json:"sourceDirectory"`
	Type            string      `json:"type"`
}

type FileGroup struct {
	CompileFlags string        `json:"compileFlags"`
	Defines      []string      `json:"defines"`
	IncludePath  []IncludePath `json:"includePath"`
	IsGenerated  bool          `json:"isGenerated"`
	Language     string        `json:"language"`
	Sources      []string      `json:"sources"`
}

type IncludePath struct {
	Path string `json:"path"`
}

func (m *MessageReply) Type() MessageType {
	return ReplyMsg
}

func (m *MessageReply) String() string {
	return fmt.Sprintf("Reply To %s", m.InReplyTo)
}
