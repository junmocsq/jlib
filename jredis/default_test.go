package jredis

import (
	"jlib/jtools"
	"testing"
)

func TestMain(m *testing.M) {
	RegisterRedisPool("default", "127.0.0.1", "6379", "", "test")
	SetDefaultModule("default")
	SetDebug(true)
	jtools.Logs("/Users/junmo/go/src/jlib/logs")
	m.Run()
}
