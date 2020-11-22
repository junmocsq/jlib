package jredis

import (
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func TestJredis_LIST(t *testing.T) {
	var r RedisLister = NewRedis()
	key := "list_kkk"
	Convey("LIST", t, func() {
		Convey("BLPOP", func() {
			r.DEL(key)
			r.LPUSH(key, "v1", "v2", "v3")
			So(strings.Join(r.BLPOP(3, key), "|"), ShouldEqual, key+"|v3")
			key2 := "list_kkk2"
			r.LPUSH(key2, "vv1", "vv2")
			So(strings.Join(r.BLPOP(3, key2, key), "|"), ShouldEqual, key2+"|vv2")
		})
	})

}
