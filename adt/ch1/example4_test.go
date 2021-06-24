package ch1

import (
	"testing"
)

func TestPerm(t *testing.T) {
	arr := []string{"a", "b", "c", "d"}
	if len(Perm(arr)) != 24 {
		t.Errorf("%v 全排列错误", arr)
	}
}

func BenchmarkPerm(b *testing.B) {
	b.StopTimer()
	arr := []string{"a", "b", "c", "d"}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Perm(arr)
	}
}
