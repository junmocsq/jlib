package jredis

import "testing"

func TestJredis_SADD(t *testing.T) {
	r := NewRedis()
	r.SADD("abc", 1, 2, 3, 100, 200, 121, 122, 12, 4, 1111, 2222, 3333, 4444, 5555, "abc")
}
