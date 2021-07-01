package ch4

import (
	"testing"
)

func TestNewStackLink(t *testing.T) {
	a := []string{"a", "b", "c", "d", "e", "f"}
	s := NewStackLink()
	for _, v := range a {
		s.Push(v)
	}
	length := len(a)
	for i := length - 1; i >= 0; i-- {
		v, _ := s.Pop()
		if v != a[i] {
			t.Errorf("stack pop failed! wanted: %s actual:%s\n ", a[i], v)
		}
	}

}
func TestNewQueueLink(t *testing.T) {
	a := []string{"a", "b", "c", "d", "e", "f"}
	q := NewQueueLink()
	for _, v := range a {
		q.lPush(v)
	}

	length := len(a)
	for i := length - 1; i >= 0; i-- {
		v, _ := q.lPop()
		if v != a[i] {
			t.Errorf("stack pop failed! wanted: %s actual:%s\n ", a[i], v)
		}
	}

	for _, v := range a {
		q.rPush(v)
	}
	for i := length - 1; i >= 0; i-- {
		v, _ := q.rPop()
		if v != a[i] {
			t.Errorf("stack pop failed! wanted: %s actual:%s\n ", a[i], v)
		}
	}

	for _, v := range a {
		q.rPush(v)
	}

	for _, wanted := range a {
		actual, _ := q.lPop()
		if actual != wanted {
			t.Errorf("stack pop failed! wanted: %s actual:%s\n ", wanted, actual)
		}
	}

	for _, v := range a {
		q.lPush(v)
	}

	for _, wanted := range a {
		actual, _ := q.rPop()
		if actual != wanted {
			t.Errorf("stack pop failed! wanted: %s actual:%s\n ", wanted, actual)
		}
	}

	hs := "Able was I ere I saw elbA"
	if !huiWen(hs) {
		t.Error("回文验证失败")
	}
	hs = "Able was I ere I saw elba"
	if huiWen(hs) {
		t.Error("回文验证失败")
	}

	mm := map[string]bool{
		"a+ ((c*d)+(kk*7))+":    true,
		"a+ ((c*d)()+(kk*7)+":   false,
		"a+ ((c*d)()+(kk*7)))+": false,
	}
	for k, v := range mm {
		if checkParentheses(k) != v {
			t.Error("括号校验失败")
		}
	}

}
