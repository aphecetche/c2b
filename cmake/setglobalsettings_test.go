package cmake

import (
	"testing"
)

func TestSetGlobalSettings(t *testing.T) {
	m := NewMessageSetGlobalSettingsWarmStart("/tmp/build")
	b, _ := m.Marshal()
	expected := `{"type":"setGlobalSettings","buildDirectory":"/tmp/build"}`
	if string(b) != expected {
		t.Errorf("Expected %s but got %s\n", expected, string(b))
	}
}
