package jredis

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestJredis_SET(t *testing.T) {
	var r RedisSetter = NewRedis()
	key := "set_kkk"

	Convey("SET", t, func() {
		Convey("SADD", func() {
			r.DEL(key)
			So(r.SADD(key, 121, 2, 3, 100, 200, "abc"), ShouldEqual, 6)
			r.DEL(key)
		})
		Convey("SCARD", func() {
			r.DEL(key)
			r.SADD(key, 121, 2, 3, 100, 200, "abc")
			So(r.SCARD(key), ShouldEqual, 6)
			r.DEL(key)
		})
		Convey("SDIFF", func() {
			key2 := key + "2"
			r.DEL(key, key2)
			r.SADD(key, 121, 2, 3, 100, 200, "abc")
			r.SADD(key2, 122, 2, 3, 100, 200, "abc")
			So(r.SDIFF(key, key2), ShouldContain, "121")
			r.DEL(key, key2)
		})

		Convey("SDIFFSTORE", func() {
			key2 := key + "2"
			key3 := key + "3"
			r.DEL(key, key2, key3)
			r.SADD(key, 121, 2, 3, 100, 200, "abc")
			r.SADD(key2, 122, 20, 3, 100, 200, "abc")
			So(r.SDIFFSTORE(key3, key, key2), ShouldEqual, 2)
			So(r.SMEMBERS(key3), ShouldContain, "121")
			So(r.SMEMBERS(key3), ShouldContain, "2")
			r.DEL(key, key2, key3)
		})

		Convey("SINTER", func() {
			key2 := key + "2"
			r.DEL(key, key2)
			r.SADD(key, 121, 2, 3, 100, 200, "abc")
			r.SADD(key2, 122, 2, 3, 100, 200, "abc")
			So(r.SINTER(key, key2), ShouldContain, "2")
			r.DEL(key, key2)
		})

		Convey("SINTERSTORE", func() {
			key2 := key + "2"
			key3 := key + "3"
			r.DEL(key, key2, key3)
			r.SADD(key, 121, 2, 3, 100, 200, "abc")
			r.SADD(key2, 122, 20, 3, 100, 200, "abc")
			So(r.SINTERSTORE(key3, key, key2), ShouldEqual, 4)
			So(r.SMEMBERS(key3), ShouldContain, "3")
			So(r.SMEMBERS(key3), ShouldContain, "100")
			r.DEL(key, key2, key3)
		})

		Convey("SISMEMBER", func() {
			r.DEL(key)
			r.SADD(key, 121, 2, 3, 100, 200, "abc")
			So(r.SISMEMBER(key, "121"), ShouldEqual, 1)
			So(r.SISMEMBER(key, "122"), ShouldEqual, 0)
			r.DEL(key)
		})

		Convey("SMEMBERS", func() {
			r.DEL(key)
			r.SADD(key, 121, 2, 3, 100, 200, "abc")
			So(r.SMEMBERS(key), ShouldContain, "3")
			Print(r.SMEMBERS(key))
			r.DEL(key)
		})

		Convey("SMOVE", func() {
			key2 := key + "2"
			r.DEL(key, key2)
			r.SADD(key, 121, 2, 3, 100, 200, "abc")
			So(r.SMOVE(key, key2, "abc"), ShouldEqual, 1)
			So(r.SMOVE(key, key2, "abcd"), ShouldEqual, 0)
			Print(r.SMEMBERS(key2))
			r.DEL(key, key2)
		})

		Convey("SPOP", func() {
			r.DEL(key)
			r.SADD(key, 100, 200, "abc")
			So(r.SPOP(key)[0], ShouldBeIn, []string{"100", "200", "abc"})
			Print(r.SMEMBERS(key))
			r.DEL(key)
		})

		Convey("SRANDMEMBER", func() {
			r.DEL(key)
			r.SADD(key, 100, 200, "abc")
			So(r.SRANDMEMBER(key)[0], ShouldBeIn, []string{"100", "200", "abc"})
			Print(r.SMEMBERS(key))
			r.DEL(key)
		})

		Convey("SREM", func() {
			r.DEL(key)
			r.SADD(key, 100, 200, "abc")
			So(r.SREM(key, "100", "200"), ShouldEqual, 2)
			Print(r.SMEMBERS(key))
			r.DEL(key)
		})

		Convey("SSCAN", func() {
			r.DEL(key)
			r.SADD(key, 100, 200, "abc")
			So(r.SSCAN(key, 1)[0], ShouldBeIn, []string{"100", "200", "abc"})
			Print(r.SMEMBERS(key))
			r.DEL(key)
		})

		Convey("SUNION", func() {
			key2 := key + "2"
			r.DEL(key, key2)
			r.SADD(key, "lx", "abc")
			r.SADD(key2, "csq", "abc")
			So(r.SUNION(key, key2), ShouldContain, "csq")
			r.DEL(key, key2)
		})

		Convey("SUNIONSTORE", func() {
			key2 := key + "2"
			key3 := key + "3"
			r.DEL(key, key2, key3)
			r.SADD(key, "lx", "abc")
			r.SADD(key2, "csq", "abc")
			So(r.SUNIONSTORE(key3, key, key2), ShouldEqual, 3)
			So(r.SMEMBERS(key3), ShouldContain, "csq")
			Print(r.SMEMBERS(key3))
			// OutPut: [abc lx csq]
			r.DEL(key, key2, key3)
		})
	})
}
