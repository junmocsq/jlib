package jredis

import (
	"github.com/gomodule/redigo/redis"
	"strings"
)

var redisPool map[string]*redis.Pool = make(map[string]*redis.Pool)
var prefixKeyArr = make(map[string]string)
var defaultModule = ""
var debug = false

func init() {
	// 注册redis default模块 必须注册
	//RegisterRedisPool("default","127.0.0.1","6379","","jredis")

	// 注册redis sql模块
	//RegisterRedisPool("sql","127.0.0.1","6379","","sql")
}

func RegisterRedisPool(module, host, port, auth, prefixKey string) {
	// 只注册一次
	if _, ok := redisPool[module]; !ok {
		redisPool[module] = initRedis(module, host, port, auth)
		if prefixKey != "" {
			prefixKeyArr[module] = prefixKey
		}
		if defaultModule == "" {
			SetDefaultModule(module)
		}
	}
}

func SetDefaultModule(module string) {
	defaultModule = module
}

func SetDebug(d ...bool) {
	if len(d) >= 1 {
		debug = d[0]
	}
}

// 获取redis key
func getKey(module, key string) string {
	if prefixKey, ok := prefixKeyArr[module]; ok && prefixKey != "" {
		return prefixKey + "_" + key
	}
	return key
}

// redis key 切去前缀
func trimPrefixKey(module, key string) string {
	if prefixKey, ok := prefixKeyArr[module]; ok && prefixKey != "" {
		return strings.Trim(key, prefixKey+"_")
	}
	return key
}
