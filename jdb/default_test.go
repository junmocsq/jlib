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

	RegisterCacheAccessor("127.0.0.1", "6379", "default")
	RegisterMasterDb("test", writeDb,
		ConnMaxLifetime(180), ConnMaxIdleTime(90), ConnMaxOpenConns(5), ConnMaxIdleConns(5), IsDefault(true))
	RegisterSlaveDb("test", readDbs)
	//SetDebug(true)
	m.Run()
}
