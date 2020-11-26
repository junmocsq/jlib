package jdb

import (
	"testing"
)

func TestTools(t *testing.T) {
	t.Log(genRandstr())

	prefix := "ahhhshhh"
	sql := "SELECT * FROM user WHERE id=? AND name=?"
	params := []interface{}{100, "lisi"}
	t.Log(hash(prefix, sql, params))
}

func BenchmarkGenRandStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		genRandstr()
	}
}

func BenchmarkHash(b *testing.B) {
	prefix := "ahhhshhh"
	sql := "SELECT * FROM user WHERE id=? AND name=?"
	params := []interface{}{100, "lisi"}
	for i := 0; i < b.N; i++ {
		hash(prefix, sql, params)
	}
}

func BenchmarkJsonEncode(b *testing.B) {
	s := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		Reg  bool   `json:"reg"`
	}{
		"lisi", 18, false,
	}
	for i := 0; i < b.N; i++ {
		JsonEncode(s)
	}
}

func BenchmarkJsonEncode2(b *testing.B) {
	s := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		Reg  bool   `json:"reg"`
	}{
		"lisi", 18, false,
	}
	for i := 0; i < b.N; i++ {
		JsonEncode2(s)
	}
}

// {"name":"lisi","age":18,"reg":false}

func BenchmarkJsonDecode(b *testing.B) {
	var s struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		Reg  bool   `json:"reg"`
	}

	ss := `{"name":"lisi","age":18,"reg":false}`
	for i := 0; i < b.N; i++ {
		JsonDecode(ss, &s)
	}
}

func BenchmarkJsonDecode2(b *testing.B) {
	var s struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		Reg  bool   `json:"reg"`
	}

	ss := `{"name":"lisi","age":18,"reg":false}`
	for i := 0; i < b.N; i++ {
		JsonDecode2(ss, &s)
	}
}
