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
		Convey("BLPOP BRPOP", func() {
			key2 := "list_kkk2"
			r.DEL(key, key2)
			r.LPUSH(key, "v1", "v2", "v3")
			So(strings.Join(r.BLPOP(3, key), "|"), ShouldEqual, key+"|v3")
			r.LPUSH(key2, "vv1", "vv2")
			So(strings.Join(r.BLPOP(3, key2, key), "|"), ShouldEqual, key2+"|vv2")
			r.DEL(key, key2)
			r.LPUSH(key, "v1", "v2", "v3")
			So(strings.Join(r.BRPOP(3, key), "|"), ShouldEqual, key+"|v1")
			r.LPUSH(key2, "vv1", "vv2")
			So(strings.Join(r.BRPOP(3, key2, key), "|"), ShouldEqual, key2+"|vv1")
			r.DEL(key, key2)
		})

		Convey("BRPOPLPUSH", func() {
			key2 := "list_kkk2"
			r.DEL(key, key2)
			r.LPUSH(key, "v1", "v2", "v3")
			So(r.BRPOPLPUSH(key, key2, 3), ShouldEqual, "v1")
			r.DEL(key, key2)
		})

		Convey("LINDEX", func() {
			r.DEL(key)
			r.RPUSH(key, "v1", "v2", "v3")
			So(r.LINDEX(key, 0), ShouldEqual, "v1")
			r.LPUSH(key, "lx")
			So(r.LINDEX(key, 0), ShouldEqual, "lx")
			r.DEL(key)
		})

		Convey("LINSERTAFTER LINSERTBEFORE", func() {
			r.DEL(key)
			r.RPUSH(key, "v1", "v2", "v3")
			So(r.LINSERTAFTER(key, "v1", "lx"), ShouldEqual, 4)
			So(r.LINSERTBEFORE(key, "lx", "xq"), ShouldEqual, 5)
			So(r.LINDEX(key, 2), ShouldEqual, "lx")
			So(r.LINDEX(key, 1), ShouldEqual, "xq")
			r.DEL(key)
		})

		Convey("LLEN", func() {
			r.DEL(key)
			r.RPUSH(key, "v1", "v2", "v3")
			So(r.LLEN(key), ShouldEqual, 3)
			r.DEL(key)
		})

		Convey("LPOP RPOP", func() {
			r.DEL(key)
			r.RPUSH(key, "v1", "v2", "v3")
			So(r.LPOP(key), ShouldEqual, "v1")
			So(r.RPOP(key), ShouldEqual, "v3")
			r.DEL(key)
		})

		Convey("LPUSH LPUSHX", func() {
			r.DEL(key)
			So(r.LPUSHX(key, "vv1", "vv2", "vv3"), ShouldEqual, 0)
			So(r.LPUSH(key, "v1", "v2", "v3"), ShouldEqual, 3)
			So(r.LPUSHX(key, "vv1", "vv2", "vv3"), ShouldEqual, 6)
			So(r.LINDEX(key, 0), ShouldEqual, "vv3")
			r.DEL(key)
		})

		Convey("LRANGE", func() {
			r.DEL(key)
			r.LPUSH(key, "v1", "v2", "v3")
			So(strings.Join(r.LRANGE(key, 0, -1), "|"), ShouldEqual, "v3|v2|v1")
			r.DEL(key)
		})

		Convey("LREM", func() {
			r.DEL(key)
			r.RPUSH(key, "v1", "v3", "v2", "v3", "v4", "v3", "v5")
			So(r.LREM(key, -2, "v3"), ShouldEqual, 2)
			So(strings.Join(r.LRANGE(key, 0, -1), "|"), ShouldEqual, "v1|v3|v2|v4|v5")
			r.DEL(key)
			r.RPUSH(key, "v1", "v3", "v2", "v3", "v4", "v3", "v5")
			So(r.LREM(key, 2, "v3"), ShouldEqual, 2)
			So(strings.Join(r.LRANGE(key, 0, -1), "|"), ShouldEqual, "v1|v2|v4|v3|v5")
			r.DEL(key)
			r.RPUSH(key, "v1", "v3", "v2", "v3", "v4", "v3", "v5")
			So(r.LREM(key, 0, "v3"), ShouldEqual, 3)
			So(strings.Join(r.LRANGE(key, 0, -1), "|"), ShouldEqual, "v1|v2|v4|v5")
			r.DEL(key)
		})

		Convey("LSET", func() {
			r.DEL(key)
			r.RPUSH(key, "v1", "v2", "v3")
			So(r.LSET(key, 0, "lx"), ShouldBeTrue)
			So(strings.Join(r.LRANGE(key, 0, -1), "|"), ShouldEqual, "lx|v2|v3")
			r.DEL(key)
		})

		Convey("LTRIM", func() {
			r.DEL(key)
			r.RPUSH(key, "v1", "v2", "v3", "v4", "v5")
			So(r.LTRIM(key, 1, 2), ShouldBeTrue)
			So(strings.Join(r.LRANGE(key, 0, -1), "|"), ShouldEqual, "v2|v3")
			r.DEL(key)
		})

		Convey("RPOPLPUSH", func() {
			key2 := "list_kkk2"
			r.DEL(key, key2)
			So(r.RPOPLPUSH(key, "key2"), ShouldEqual, "")
			r.RPUSH(key, "v1", "v2", "v3", "v4", "v5")
			So(r.RPOPLPUSH(key, "key2"), ShouldEqual, "v5")
			r.DEL(key, key2)
		})

		Convey("RPUSH RPUSHX", func() {
			r.DEL(key)
			So(r.RPUSHX(key, "vv1", "vv2", "vv3"), ShouldEqual, 0)
			So(r.RPUSH(key, "v1", "v2", "v3"), ShouldEqual, 3)
			So(r.RPUSHX(key, "vv1", "vv2", "vv3"), ShouldEqual, 6)
			So(r.LINDEX(key, 5), ShouldEqual, "vv3")
			r.DEL(key)
		})
	})

}
