package jdb

import "jlib/jredis"

var (
	expire      = 300
	cacheAccess cacheAccesser
	dbAccess    DbAccessor
	Empty       = "nil"

	redisModule = "sql"
	redisClient = jredis.NewRedis(redisModule)
)

func RegisterCacheAccesser(host, port, auth string) {
	jredis.RegisterRedisPool(redisModule, "127.0.0.1", "6379", "", redisModule)
	jredis.SetDebug(true)
}

func RegisterDbAccesser() {

}
