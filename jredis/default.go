package jredis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
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

type Conf struct {
	module          string
	host            string
	port            string
	auth            string
	prefixKey       string
	isDefaultModule bool
}

type Option func(*Conf)

func ModuleConf(module string) Option {
	return func(conf *Conf) {
		conf.module = module
	}
}

func AuthConf(auth string) Option {
	return func(conf *Conf) {
		conf.auth = auth
	}
}
func PrefixConf(prefixKey string) Option {
	return func(conf *Conf) {
		conf.prefixKey = prefixKey
	}
}

func IsDefaultModule(b bool) Option {
	return func(conf *Conf) {
		conf.isDefaultModule = b
	}
}

func RegisterRedisPool(host, port string, options ...Option) {
	conf := &Conf{
		module:          "default",
		host:            host,
		port:            port,
		auth:            "",
		prefixKey:       "",
		isDefaultModule: false,
	}
	for _, option := range options {
		option(conf)
	}
	// 只注册一次
	if _, ok := redisPool[conf.module]; !ok {
		redisPool[conf.module] = initRedis(conf.module, host, port, conf.auth)
		if conf.prefixKey != "" {
			prefixKeyArr[conf.module] = conf.prefixKey
		}
		if defaultModule == "" || conf.isDefaultModule {
			defaultModule = conf.module
		}
		logrus.WithField("redis", "init").
			Infof("init redis module:%s host:%s port:%s auth:%s prefixKey:%s defaultModule:%s",
				conf.module, host, port, conf.auth, prefixKeyArr[conf.module], defaultModule)
	} else {
		panic("redis module :" + conf.module + " is registered！！")
	}
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
