package cmake

import (
	"fmt"
	"testing"
)

func TestMessageGeneric(t *testing.T) {
	m, err := UnmarshalGeneric([]byte(`{"type":"toto","message":"zon"}`))
	fmt.Println(m, err)
}
