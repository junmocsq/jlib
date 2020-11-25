package jredis

import (
	. "github.com/smartystreets/goconvey/convey"
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestJredis_STRING(t *testing.T) {
	key := "jredis_string"
	rand.Seed(time.Now().UnixNano())
	val := rand.Int()
	var r RedisStringer = NewRedis()

	Convey("STRING", t, func() {
		Convey("SET GET", func() {
			So(r.SET(key, val), ShouldBeTrue)
			So(r.GET(key), ShouldEqual, strconv.Itoa(val))
		})

		Convey("SETEX", func() {
			r.DEL(key)
			So(r.SETEX(key, val, 600), ShouldBeTrue)
			So(r.TTL(key), ShouldBeGreaterThan, 600-2)
		})

		Convey("SETNX", func() {
			r.DEL(key)
			So(r.SETNX(key, val, 600), ShouldBeTrue)
			So(r.SETNX(key, val, 600), ShouldBeFalse)
		})

		Convey("APPEND", func() {
			r.DEL(key)
			value := "hello"
			So(r.APPEND(key, value), ShouldEqual, len(value))
			So(r.APPEND(key, value), ShouldEqual, 2*len(value))
			r.DEL(key)
		})

		Convey("DECR", func() {
			r.DEL(key)
			value := 100
			r.SET(key, value)
			So(r.DECR(key), ShouldEqual, value-1)
			r.DEL(key)
		})

		Convey("DECRBY", func() {
			r.DEL(key)
			value := 100
			r.SET(key, value)
			decrement := 10
			So(r.DECRBY(key, decrement), ShouldEqual, value-decrement)
			r.DEL(key)
		})

		Convey("SETBIT GETBIT", func() {
			r.DEL(key)
			r.SETBIT(key, 10, 1)
			So(r.GETBIT(key, 10), ShouldEqual, 1)
			r.DEL(key)
		})

		Convey("GETRANGE", func() {
			r.DEL(key)
			r.SET(key, "hello,world")
			So(r.GETRANGE(key, 0, 0), ShouldEqual, "h")
			So(r.GETRANGE(key, 0, -1), ShouldEqual, "hello,world")
			So(r.GETRANGE(key, 0, 3), ShouldEqual, "hell")
			Print(r.GETRANGE(key, 0, 3))
			// OutPut: hello
			r.DEL(key)
		})

		Convey("GETSET", func() {
			r.DEL(key)
			value1 := "hello,world 1"
			value2 := "hello,world 2"
			So(r.GETSET(key, value1), ShouldEqual, "")
			So(r.GETSET(key, value2), ShouldEqual, value1)
			r.DEL(key)
		})

		Convey("INCR INCRBY", func() {
			r.DEL(key)
			value1 := 100
			r.SET(key, value1)
			So(r.INCR(key), ShouldEqual, value1+1)
			So(r.INCRBY(key, 10), ShouldEqual, value1+11)
			r.DEL(key)
		})

		Convey("INCRBYFLOAT", func() {
			r.DEL(key)
			value1 := 100.2
			r.SET(key, value1)
			So(r.INCRBYFLOAT(key, 10.2), ShouldEqual, 110.4)
			So(r.INCRBYFLOAT(key, -11.6), ShouldEqual, 98.8)
			r.DEL(key)
		})

		Convey("MSET MGET", func() {
			k1, k2, k3 := "tt_k1", "tt_k2", "tt_k3"
			v1, v2, v3 := "tt_v1", "tt_v2", "tt_v3"
			r.DEL(k1, k2, k3)
			mapKv := make(map[string]interface{})
			mapKv[k1] = v1
			mapKv[k2] = v2
			mapKv[k3] = v3
			So(r.MSET(mapKv), ShouldBeTrue)
			So(strings.Join(r.MGET(k1, k2, k3), "|"), ShouldEqual, v1+"|"+v2+"|"+v3)
			r.DEL(key)
		})

		Convey("MSETNX", func() {
			k1, k2, k3 := "tt_k1", "tt_k2", "tt_k3"
			v1, v2, v3 := "tt_v1", "tt_v2", "tt_v3"
			r.DEL(k1, k2, k3)
			mapKv := make(map[string]interface{})
			mapKv[k1] = v1
			mapKv[k2] = v2

			So(r.MSETNX(mapKv), ShouldEqual, 1)
			delete(mapKv, k1)
			mapKv[k3] = v3
			So(r.MSETNX(mapKv), ShouldEqual, 0)
			So(strings.Join(r.MGET(k1, k2, k3), "|"), ShouldEqual, v1+"|"+v2+"|")
			Printf("%#v", r.MGET(k1, k2, k3))
			r.DEL(key)
		})

		Convey("SETRANGE", func() {
			r.DEL(key)
			value1 := "hello"
			value2 := "world"
			r.SET(key, value1)
			So(r.SETRANGE(key, 10, value2), ShouldEqual, 15)
			Print(r.STRLEN(key), []byte(r.GET(key)))
			So(r.SETRANGE(key, 3, value2), ShouldEqual, 15)
			Print(r.STRLEN(key), []byte(r.GET(key)))
			r.DEL(key)
		})

		Convey("STRLEN", func() {
			r.DEL(key)
			value1 := "hello,world"
			r.SET(key, value1)
			So(r.STRLEN(key), ShouldEqual, len(value1))
			r.DEL(key)
		})
	})
}

func BenchmarkJredis_String(b *testing.B) {
	r := NewRedis()
	k := "Benchmark_String"
	for i := 0; i < b.N; i++ {
		r.SET(k, i)
	}
}
