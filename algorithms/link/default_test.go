package link

import "testing"

func TestEqual(t *testing.T) {
	v1, v2 := 1, true
	Equal(v1, v2)
}
