package jredis

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	RegisterRedisPool("default", "127.0.0.1", "6379", "", "test")
	SetDefaultModule("default")
	//SetDebug(true)
	m.Run()
}

func ExampleNewRedis() {
	r := NewRedis("test")
	fmt.Println(r.module)
	// output: test
}
