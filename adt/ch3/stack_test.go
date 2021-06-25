package ch3

import (
	"errors"
	"testing"
)

func TestNewStack(t *testing.T) {
	s := NewStack()
	for i := 0; i < StackSize; i++ {
		s.Add(&Element{Val: i})
	}
	if err := s.Add(&Element{Val: 999}); err == nil {
		t.Errorf("满stack还能继续添加元素")
	} else {
		if !errors.Is(err, ErrorStackFull) {
			t.Errorf("满stack添加元素未知错误")
		}
	}
	for i := 0; i < StackSize; i++ {
		s.Delete()
	}
	if _, err := s.Delete(); err == nil {
		t.Errorf("空stack还能继续删除元素")
	} else {
		if !errors.Is(err, ErrorStackEmpty) {
			t.Errorf("空stack还能继续删除元素未知错误")
		}
	}

	//p1_4(3)
}

func BenchmarkNewStack(b *testing.B) {
	s := NewStack()
	for i := 0; i < b.N; i++ {
		s.Add(&Element{Val: i})
		s.Delete()
	}
}
