package algorithms

import (
	"github.com/junmocsq/jlib/algorithms/link"
	"testing"
)

func TestNewLink(t *testing.T) {
	var arrStack = []Queue{
		NewArrayQueue(), NewQueueLink(1), NewQueueLink(2), NewQueueLink(3), NewQueueLink(4), NewLoopArrayQueue(),
	}
	arr := []interface{}{"junmo", "zxf", "lxq", "lmm"}
	for _, a := range arrStack {
		for _, v := range arr {
			a.Enqueue(v)
		}

		for _, v := range arr {
			if !link.Equal(a.Dequeue(), v) {
				t.Error("stack failed!", v)
			}
		}
	}

	for _, v := range arr {
		t.Log(arrStack[5].Enqueue(v),
			arrStack[5].Enqueue(v),
			arrStack[5].Enqueue(v))
	}

	t.Log(arrStack[5])
}
