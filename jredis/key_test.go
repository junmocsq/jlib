package jredis

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestJredis_KEY(t *testing.T) {
	var r RedisStringer = NewRedis()
	k1, k2, k3 := "k1", "k2", "k3"

	Convey("KEY", t, func() {
		Convey("DEL", func() {
			r.SET(k1, k1)
			r.SET(k2, k2)
			r.SET(k3, k3)
			n := r.DEL(k1, k2, k3)
			So(n, ShouldEqual, 3)
		})

		Convey("EXPIRE", func() {
			r.SET(k1, k1)
			expire := 600
			So(r.EXPIRE(k1, expire), ShouldEqual, 1)
			r.DEL(k1)
		})

		Convey("TTL", func() {
			expire := 600
			r.SETEX(k1, k1, expire)
			So(r.TTL(k1), ShouldBeGreaterThan, expire-1)
		})

		Convey("EXISTS", func() {
			r.DEL(k1)
			So(r.EXISTS(k1), ShouldEqual, 0)
			r.SET(k1, k1)
			So(r.EXISTS(k1), ShouldEqual, 1)
			r.DEL(k1)
		})

		Convey("RENAME", func() {
			k1 := "oldk1"
			k2 := "newk1"
			expire := 600
			r.SETEX(k1, k1, expire)
			So(r.EXISTS(k1), ShouldEqual, 1)
			r.RENAME(k1, k2)
			So(r.EXISTS(k2), ShouldEqual, 1)
			r.DEL(k2)
		})

		Convey("SORT", func() {
			k1 := "oldk1"
			r.DEL(k1)
			NewRedis().SADD(k1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
			res, _ := r.SORT(k1, 0, 5, true)
			So(res[0], ShouldEqual, 10)
			r.DEL(k1)
		})
	})

}
