package jdb

import "github.com/junmocsq/jlib/jredis"

var (
	expire      = 300
	cacheAccess cacheAccesser
	Empty       = "nil"
	DEBUG       = false

	redisModule = "sql"
	redisClient = jredis.NewRedis(redisModule)
)

func RegisterCacheAccesser(host, port, auth string) {
	jredis.RegisterRedisPool(redisModule, "127.0.0.1", "6379", "", redisModule)
	jredis.SetDebug(true)
	cacheAccess = newCache()
}

func SetDebug(debug bool) {
	DEBUG = debug
}

func RegisterDbAccesser() {

}
