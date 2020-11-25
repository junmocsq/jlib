package jredis

import (
	"fmt"
	"github.com/junmocsq/jlib/jtools"
	"testing"
)

func TestMain(m *testing.M) {
	RegisterRedisPool("default", "127.0.0.1", "6379", "", "test")
	SetDefaultModule("default")
	SetDebug(true)
	jtools.Logs("/Users/junmo/go/src/jlib/logs")
	m.Run()
}

func ExampleNewRedis() {
	r := NewRedis("test")
	fmt.Println(r.module)
	// output: test
}
