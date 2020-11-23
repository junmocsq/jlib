package jredis

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestJredis_HASH(t *testing.T) {
	var r = NewRedis()
	key := "hash_kkk"

	Convey("HASH", t, func() {
		Convey("HDEL HEXISTS", func() {
			r.DEL(key)
			r.HMSET(key, map[interface{}]interface{}{
				"junmo": "lxq",
				"csq":   "lx",
				"clq":   "jr",
			})

			So(r.HDEL(key, "junmo"), ShouldEqual, 1)
			So(r.HEXISTS(key, "junmo"), ShouldEqual, 0)
			r.HSET(key, "junmo", "")
			So(r.HEXISTS(key, "junmo"), ShouldEqual, 1)
			r.DEL(key)
		})

		Convey("HGET", func() {
			r.DEL(key)
			r.HMSET(key, map[interface{}]interface{}{
				"junmo": "lxq",
				"csq":   "lx",
				"clq":   "jr",
			})

			So(r.HGET(key, "junmo"), ShouldEqual, "lxq")
			r.DEL(key)
		})

		Convey("HINCRBY", func() {
			r.DEL(key)
			r.HSET(key, "lx", 10)

			So(r.HINCRBY(key, "lx", 10), ShouldEqual, 20)
			r.DEL(key)
		})

		Convey("HINCRBYFLOAT", func() {
			r.DEL(key)
			r.HSET(key, "lx", 10.2)

			So(r.HINCRBYFLOAT(key, "lx", 3.5), ShouldEqual, 13.7)
			r.DEL(key)
		})

		Convey("HKEYS", func() {
			r.DEL(key)
			r.HMSET(key, map[interface{}]interface{}{
				"junmo": "lxq",
				"csq":   "lx",
				"clq":   "jr",
			})
			Println(r.HKEYS(key))
			So(r.HKEYS(key), ShouldContain, "junmo")
			r.DEL(key)
		})

		Convey("HLEN", func() {
			r.DEL(key)
			r.HMSET(key, map[interface{}]interface{}{
				"junmo": "lxq",
				"csq":   "lx",
				"clq":   "jr",
			})
			So(r.HLEN(key), ShouldEqual, 3)
			r.DEL(key)
		})

		Convey("HMSET HMGET", func() {
			r.DEL(key)
			r.HMSET(key, map[interface{}]interface{}{
				"junmo": "lxq",
				"csq":   "lx",
				"clq":   "jr",
			})
			So(r.HMSET(key, map[interface{}]interface{}{
				"junmo": "lxq",
				"csq":   "lx",
				"clq":   "jr",
			}), ShouldBeTrue)
			res := r.HMGET(key, "clq", "csq", "lx")
			So(res[0], ShouldEqual, "jr")
			So(res[2], ShouldEqual, "")
			r.DEL(key)
		})

		Convey("HSCAN HGETALL", func() {
			r.DEL(key)
			r.HMSET(key, map[interface{}]interface{}{
				"junmo": "lxq",
				"csq":   "lx",
				"clq":   "jr",
			})
			Println(r.HGETALL(key))
			Println(r.HSCAN(key, 1))
			r.DEL(key)
		})

		Convey("HSET", func() {
			r.DEL(key)
			So(r.HSET(key, "junmo", "lx"), ShouldEqual, 1)
			So(r.HGET(key, "junmo"), ShouldEqual, "lx")
			r.DEL(key)
		})

		Convey("HSETNX", func() {
			r.DEL(key)
			So(r.HSETNX(key, "junmo", "lx"), ShouldEqual, 1)
			So(r.HGET(key, "junmo"), ShouldEqual, "lx")
			So(r.HSETNX(key, "junmo", "lxq"), ShouldEqual, 0)
			So(r.HGET(key, "junmo"), ShouldEqual, "lx")
			r.DEL(key)
		})

		Convey("HSTRLEN", func() {
			r.DEL(key)
			So(r.HSET(key, "junmo", "lx"), ShouldEqual, 1)
			So(r.HSTRLEN(key, "junmo"), ShouldEqual, 2)
			r.DEL(key)
		})

		Convey("HVALS", func() {
			r.DEL(key)
			r.HMSET(key, map[interface{}]interface{}{
				"junmo": "lxq",
				"csq":   "lx",
				"clq":   "jr",
			})
			So(r.HVALS(key), ShouldContain, "lxq")
			Print(r.HVALS(key))
			r.DEL(key)
		})

	})
}
