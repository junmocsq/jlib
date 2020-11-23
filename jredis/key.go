package jredis

import (
	"github.com/gomodule/redigo/redis"
)

type RedisKeyer interface {
	DEL(keys ...string) int
	EXPIRE(key string, expire int) int
	TTL(key string) int
	EXISTS(key string) int
	RENAME(key, newKey string) bool
	SCAN(count int, pattern ...string) []string
	SORT(key string, start, size int, isReverse ...bool) ([]float64, error)
}

func (j *jredis) DEL(keys ...string) int {
	length := len(keys)
	if length == 0 {
		return 0
	}
	karr := make([]interface{}, length)
	for index, key := range keys {
		karr[index] = j.getKey(key)
	}
	res, err := j.exec("DEL", karr...)
	n, _ := redis.Int(res, err)
	return n
}

func (j *jredis) EXPIRE(key string, expire int) int {
	key = j.getKey(key)
	res, err := j.exec("EXPIRE", key, expire)
	n, _ := redis.Int(res, err)
	return n
}

// -1 永不过期
// -2 key不存在
// 大于0 剩余时间
func (j *jredis) TTL(key string) int {
	key = j.getKey(key)
	res, err := j.exec("TTL", key)
	n, _ := redis.Int(res, err)
	return n
}

func (j *jredis) EXISTS(key string) int {
	key = j.getKey(key)
	res, err := j.exec("EXISTS", key)
	n, _ := redis.Int(res, err)
	return n
}

func (j *jredis) RENAME(key, newKey string) bool {
	res, err := j.exec("RENAME", j.getKey(key), j.getKey(newKey))
	return j.isOk(res, err)
}

func (j *jredis) SCAN(count int, pattern ...string) []string {
	patt := ""
	if len(pattern) >= 1 {
		patt = pattern[0]
	}
	var list []string
	seed := "0"
	for {
		var res interface{}
		var err error
		if patt != "" {
			res, err = j.exec("SCAN", seed, "COUNT", count, "MATCH", patt)
		} else {
			res, err = j.exec("SCAN", seed, "COUNT", count)
		}
		arr, err := redis.Values(res, err)
		if err != nil {
			break
		}
		if s, ok := arr[0].([]uint8); ok {
			seed = string(s)
		}
		// Starting an iteration with a cursor value of 0, and calling SCAN
		// until the returned cursor is 0 again is called a full iteration.
		if seed == "0" {
			break
		}
		if arrVal, ok := arr[1].([]interface{}); ok {
			for _, v := range arrVal {
				if val, ok := v.([]uint8); ok {
					list = append(list, string(val))
				}
			}
		}
	}
	return list
}

func (j *jredis) SORT(key string, start, size int, isReverse ...bool) ([]float64, error) {
	reverse := "ASC"
	if len(isReverse) >= 1 && isReverse[0] {
		reverse = "DESC"
	}
	res, err := j.exec("SORT", j.getKey(key), "LIMIT", start, size, reverse)
	r, e := redis.Float64s(res, err)
	if e != nil {
		return nil, e
	}
	return r, nil
}
