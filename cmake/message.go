package cmake

import (
	"encoding/json"
	"fmt"
	"io"
)

type MessageType int

const (
	MessageStart = `[== "CMake Server" ==[`
	MessageEnd   = `]== "CMake Server" ==]`
)

const (
	HelloMsg MessageType = iota + 1
	HandshakeMsg
	ReplyMsg
	ErrorMsg
	ProgressMsg
	MessageMsg
	SignalMsg
	GlobalSettingsMsg
	SetGlobalSettingsMsg
	ConfigureMsg
	ComputeMsg
	CodeModelMsg
	CTestInfoMsg
	CMakeInputsMsg
	CacheMsg
	FileSystemWatchersMsg
)

func TypeFromString(s string) MessageType {
	switch s {
	case HelloMsg.String():
		return HelloMsg
	case HandshakeMsg.String():
		return HandshakeMsg
	case ReplyMsg.String():
		return ReplyMsg
	case ErrorMsg.String():
		return ErrorMsg
	case ProgressMsg.String():
		return ProgressMsg
	case MessageMsg.String():
		return MessageMsg
	case SignalMsg.String():
		return SignalMsg
	case GlobalSettingsMsg.String():
		return GlobalSettingsMsg
	case SetGlobalSettingsMsg.String():
		return SetGlobalSettingsMsg
	case ConfigureMsg.String():
		return ConfigureMsg
	case ComputeMsg.String():
		return ComputeMsg
	case CodeModelMsg.String():
		return CodeModelMsg
	case CTestInfoMsg.String():
		return CTestInfoMsg
	case CMakeInputsMsg.String():
		return CMakeInputsMsg
	case CacheMsg.String():
		return CacheMsg
	case FileSystemWatchersMsg.String():
		return FileSystemWatchersMsg
	default:
		panic(fmt.Errorf("unknown message type %s", s))
	}
}

func (msg MessageType) String() string {
	switch msg {
	case HelloMsg:
		return "hello"
	case HandshakeMsg:
		return "handshake"
	case ReplyMsg:
		return "reply"
	case ErrorMsg:
		return "error"
	case ProgressMsg:
		return "progress"
	case MessageMsg:
		return "message"
	case SignalMsg:
		return "signal"
	case GlobalSettingsMsg:
		return "globalSettings"
	case SetGlobalSettingsMsg:
		return "setGlobalSettings"
	case ConfigureMsg:
		return "configure"
	case ComputeMsg:
		return "compute"
	case CodeModelMsg:
		return "codemodel"
	case CTestInfoMsg:
		return "ctestinfo"
	case CMakeInputsMsg:
		return "cmakeinputs"
	case CacheMsg:
		return "cache"
	case FileSystemWatchersMsg:
		return "fileSystemWatchers"
	default:
		return "unknown"
	}
}

type Message interface {
	Marshal() ([]byte, error)
	Type() MessageType
	String() string
}

func NewMessage(b []byte) (Message, error) {
	m, err := UnmarshalGeneric(b)
	if err != nil {
		return nil, err
	}
	return NewMessageWithType(b, TypeFromString(m.TypeString))
}

func NewMessageWithType(b []byte, msgType MessageType) (Message, error) {
	var m Message
	var err error
	switch msgType {
	case HelloMsg:
		m, err = UnmarshalMessageHello(b)
	case HandshakeMsg:
		m, err = UnmarshalMessageHandshake(b)
	case ReplyMsg:
		m, err = UnmarshalMessageReply(b)
	case GlobalSettingsMsg:
		m, err = UnmarshalMessageGlobalSettings(b)
	case SetGlobalSettingsMsg:
		m, err = UnmarshalMessageSetGlobalSettings(b)
	case ConfigureMsg:
		m, err = UnmarshalMessageConfigure(b)
	case ComputeMsg:
		m, err = UnmarshalMessageCompute(b)
	case MessageMsg:
		m, err = UnmarshalMessageMessage(b)
	case ProgressMsg:
		m, err = UnmarshalMessageProgress(b)
	case ErrorMsg:
		m, err = UnmarshalMessageError(b)
	case CodeModelMsg:
		m, err = UnmarshalMessageCodeModel(b)
	default:
		return nil, fmt.Errorf("no such message type : %d", msgType)
	}
	if err != nil {
		return nil, err
	}
	return m, nil
}

func Write(m Message, w io.Writer) {
	b, _ := m.Marshal()
	fmt.Println("WRITE:", string(b))
	msg := fmt.Sprintf("%s\n%s\n%s\n", MessageStart, string(b), MessageEnd)
	w.Write([]byte(msg))
}

func UnmarshalGeneric(data []byte) (MessageGeneric, error) {
	var r MessageGeneric
	err := json.Unmarshal(data, &r)
	return r, err
}

type MessageGeneric struct {
	TypeString string `json:"type"`
}
