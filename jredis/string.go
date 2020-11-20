package jredis

import (
	"github.com/gomodule/redigo/redis"
)

type RedisStringer interface {
}

func (j *jredis) SET(key string, value interface{}) bool {
	res, err := j.exec("SET", j.getKey(key), value)
	return j.isOk(res, err)
}

// 设置过期时间
func (j *jredis) SETEX(key string, value interface{}, expire int) bool {
	res, err := j.exec("SET", j.getKey(key), value, "EX", expire)
	return j.isOk(res, err)
}

// 加锁
func (j *jredis) SETNX(key string, value interface{}, expire int) bool {
	res, err := j.exec("SET", j.getKey(key), value, "NX", "EX", expire)
	return j.isOk(res, err)
}

func (j *jredis) GET(key string) (string, error) {
	res, err := j.exec("GET", j.getKey(key))
	return redis.String(res, err)
}
