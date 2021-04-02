package algorithms

import (
	"github.com/junmocsq/jlib/algorithms/link"
	"testing"
)

func TestNewArray(t *testing.T) {
	var arrStack = []Stacker{
		NewArray(), NewLink(1), NewLink(2), NewLink(3), NewLink(4), NewLoopArrayStack(),
	}
	arr := []interface{}{"junmo", "zxf", "lxq", "lmm"}
	for _, a := range arrStack {
		for _, v := range arr {
			a.Push(v)
		}

		for i := len(arr) - 1; i >= 0; i-- {
			if !link.Equal(a.Pop(), arr[i]) {
				t.Error("stack failed!", arr[i])
			}
		}
	}
}
