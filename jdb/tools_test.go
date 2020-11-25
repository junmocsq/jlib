package jdb

import "testing"

func TestTools(t *testing.T) {
	t.Log(genRandstr())
}

func BenchmarkGenRandStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		genRandstr()
	}
}
