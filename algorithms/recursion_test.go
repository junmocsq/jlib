package algorithms

import (
	"testing"
	"time"
)

func TestRecursion(t *testing.T) {
	f := NewF()
	key := 20
	t1 := time.Now()
	t.Log(f.f(key), time.Since(t1))
	t2 := time.Now()
	t.Log(f.fUnRepeat(key), time.Since(t2))

	// recursion_test.go:12: 165580141 567.029561ms
	// recursion_test.go:14: 165580141 64.569µs
}

func BenchmarkNewF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f := NewF()
		key := 20
		f.f(key)
	}
}

func BenchmarkNewF_UnRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f := NewF()
		key := 20
		f.fUnRepeat(key)
	}
}
