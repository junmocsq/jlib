package jredis

import (
	"github.com/gomodule/redigo/redis"
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func TestJredis_ZSET(t *testing.T) {
	var r RedisZSetter = NewRedis()
	key := "zset_kkk"

	Convey("ZSET", t, func() {
		Convey("ZADD", func() {
			r.DEL(key)
			So(r.ZADD(key, "CH", 10, "junmo", 12.1, "lx", -18.9, "csq"), ShouldEqual, 3)
			Print(r.ZRANGE(key, 0, -1))
			// NX 只添加，不更新
			So(r.ZADD(key, "CH", "NX", 22, "junmo", 11.1, "lx", 9.7, "csq"), ShouldEqual, 0)
			// XX 只更新，不添加
			So(r.ZADD(key, "CH", "XX", 22, "junmo", 11.1, "lx", 9.7, "csq", 0.8, "lxq"), ShouldEqual, 3)
			// 无lxq
			Print(r.ZRANGEBYSCORE(key, "-inf", "+inf"))
			r.DEL(key)
		})

		Convey("ZCARD", func() {
			r.DEL(key)
			l := r.ZADD(key, "CH", 10, "junmo", 12.1, "lx", -18.9, "csq")
			So(r.ZCARD(key), ShouldEqual, l)
			r.DEL(key)
		})

		Convey("ZCOUNT", func() {
			r.DEL(key)
			r.ZADD(key, "CH", 10, "junmo", 12.1, "lx", -18.9, "csq")
			So(r.ZCOUNT(key, 0, 13), ShouldEqual, 2)
			r.DEL(key)
		})

		Convey("ZINCRBY", func() {
			r.DEL(key)
			r.ZADD(key, "CH", 10, "junmo", 12.1, "lx", -18.9, "csq")
			So(r.ZINCRBY(key, 3.4, "lx"), ShouldEqual, 15.5)
			r.DEL(key)
		})

		Convey("ZLEXCOUNT ZRANGEBYLEX", func() {
			r.DEL(key)
			r.ZADD(key, "CH", 10, "a", 10, "b", 10, "c", 10, "d", 10, "cc", 10, "ab")
			So(r.ZLEXCOUNT(key, "[a", "[b"), ShouldEqual, 3)
			So(strings.Join(r.ZRANGEBYLEX(key, "[a", "[c", "LIMIT", 0, 3), "|"), ShouldEqual, "a|ab|b")
			r.DEL(key)
		})

		Convey("ZRANGE ZRANGEBYSCORE", func() {
			r.DEL(key)
			r.ZADD(key, "CH", 10, "junmo", 12.1, "lx", -18.9, "csq")
			eles, _ := r.ZRANGE(key, 0, 1)
			So(strings.Join(eles, "|"), ShouldEqual, "csq|junmo")
			eles, _ = r.ZRANGEBYSCORE(key, "(10", "(20")
			So(strings.Join(eles, "|"), ShouldEqual, "lx")
			r.DEL(key)
		})

		Convey("ZRANK", func() {
			r.DEL(key)
			r.ZADD(key, "CH", 10, "junmo", 12.1, "lx", -18.9, "csq")
			So(r.ZRANK(key, "junmo"), ShouldEqual, 1)
			So(r.ZRANK(key, "junmocsq"), ShouldEqual, -1)
			r.DEL(key)
		})

		Convey("ZREM", func() {
			r.DEL(key)
			r.ZADD(key, "CH", 10, "junmo", 12.1, "lx", -18.9, "csq")
			So(r.ZREM(key, "junmo", "lx"), ShouldEqual, 2)
			So(r.ZRANK(key, "junmo"), ShouldEqual, -1)
			r.DEL(key)
		})

		Convey("ZREMRANGEBYLEX", func() {
			r.DEL(key)
			r.ZADD(key, "CH", 10, "a", 10, "b", 10, "c", 10, "d", 10, "cc", 10, "ab")
			So(r.ZLEXCOUNT(key, "[a", "[b"), ShouldEqual, 3)
			So(r.ZREMRANGEBYLEX(key, "[a", "[b"), ShouldEqual, 3)
			Print(r.ZRANGE(key, 0, -1))
			r.DEL(key)
		})

		Convey("ZREMRANGEBYRANK", func() {
			r.DEL(key)
			r.ZADD(key, "CH", 10, "a", 10, "b", 10, "c", 10, "d", 10, "cc", 10, "ab")
			So(r.ZREMRANGEBYRANK(key, 0, 3), ShouldEqual, 4)
			eles, _ := r.ZRANGE(key, 0, 1)
			So(strings.Join(eles, "|"), ShouldEqual, "cc|d")
			r.DEL(key)
		})

		Convey("ZREMRANGEBYSCORE", func() {
			r.DEL(key)
			r.ZADD(key, "CH", 10, "junmo", 12.1, "lx", -18.9, "csq")
			So(r.ZREMRANGEBYSCORE(key, "(10", "12.1"), ShouldEqual, 1)
			eles, _ := r.ZRANGE(key, 0, 1)
			So(strings.Join(eles, "|"), ShouldEqual, "csq|junmo")
			r.DEL(key)
		})

		// ZREVRANGE
		Convey("ZREVRANGE", func() {
			r.DEL(key)
			r.ZADD(key, "CH", 10, "junmo", 12.1, "lx", -18.9, "csq")
			eles, _ := r.ZREVRANGE(key, 0, 1)
			So(strings.Join(eles, "|"), ShouldEqual, "lx|junmo")
			r.DEL(key)
		})

		Convey("ZREVRANGEBYLEX", func() {
			r.DEL(key)
			r.ZADD(key, "CH", 10, "a", 10, "b", 10, "c", 10, "d", 10, "cc", 10, "ab")
			eles := r.ZREVRANGEBYLEX(key, "(d", "(a")
			So(strings.Join(eles, "|"), ShouldEqual, "cc|c|b|ab")
			r.DEL(key)
		})

		Convey("ZREVRANGEBYSCORE", func() {
			r.DEL(key)
			r.ZADD(key, "CH", 10, "junmo", 12.1, "lx", -18.9, "csq")
			eles, _ := r.ZREVRANGEBYSCORE(key, "12.1", "10")
			So(strings.Join(eles, "|"), ShouldEqual, "lx|junmo")
			r.DEL(key)
		})

		Convey("ZREVRANK", func() {
			r.DEL(key)
			r.ZADD(key, "CH", 10, "junmo", 12.1, "lx", -18.9, "csq")
			So(r.ZREVRANK(key, "csq"), ShouldEqual, 2)
			r.DEL(key)
		})

		Convey("ZSCAN", func() {
			r.DEL(key)
			r.ZADD(key, "CH", 10, "junmo", 12.1, "lx", -18.9, "csq")
			m := r.ZSCAN(key, 100)
			Print(m)
			r.DEL(key)
		})

		Convey("ZSCORE", func() {
			r.DEL(key)
			r.ZADD(key, "CH", 10, "junmo", 12.1, "lx", -18.9, "csq")
			res, _ := r.ZSCORE(key, "lx")
			So(res, ShouldEqual, 12.1)
			_, err := r.ZSCORE(key, "junmocsq")
			So(err.Error(), ShouldEqual, redis.ErrNil.Error())
			r.DEL(key)
		})
	})
}
