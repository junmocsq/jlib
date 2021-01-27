package jdb

import (
	"fmt"
	"gopkg.in/ini.v1"
	"testing"
)

func TestMain(m *testing.M) {
	cfg, err := ini.Load("conf.ini", "conf.ini.local")

	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		return
	}
	writeDb := cfg.Section("db").Key("test_write").String()
	readDbs := cfg.Section("db").Key("test_read").Strings(",")

	RegisterCacheAccesser("127.0.0.1", "6379", "default")
	RegisterSqlDb("test", false, writeDb)
	RegisterSqlDb("test", true, readDbs...)
	SetDbPoolParams("test", 180, 90, 5, 5)
	SetDebug(true)
	m.Run()
}
