package ch3

import "testing"

func TestNewQueue(t *testing.T) {
	q := NewQueue()
	for i := 100; i < 195; i++ {
		q.PushRight(&Element{Val: i})
		q.PopLeft()
	}
	q.print()
	for i := 100; i < 110; i++ {
		q.PushRight(&Element{Val: i})
	}
	q.print()
	q.PopRight()
	q.print()
}

func BenchmarkNewQueue(b *testing.B) {
	s := NewQueue()
	for i := 0; i < b.N; i++ {
		s.PushRight(&Element{Val: i})
		s.PopLeft()
	}
}
