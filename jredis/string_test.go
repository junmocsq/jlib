package jredis

import (
	. "github.com/smartystreets/goconvey/convey"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	RegisterRedisPool("default", "127.0.0.1", "6379", "", "test")
	SetDefaultModule("default")
	m.Run()
}
func TestJredis_STRING(t *testing.T) {
	key := "jredis_string"
	rand.Seed(time.Now().UnixNano())
	val := rand.Int()
	r := NewRedis()

	Convey("STRING", t, func() {
		Convey("SET GET", func() {
			So(r.SET(key, val), ShouldBeTrue)
			res, _ := r.GET(key)
			So(res, ShouldEqual, strconv.Itoa(val))
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
	})
}
