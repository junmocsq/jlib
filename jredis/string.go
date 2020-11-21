package jredis

import (
	"github.com/gomodule/redigo/redis"
)

type RedisStringer interface {
	SET(key string, value interface{}) bool
	SETEX(key string, value interface{}, expire int) bool
	SETNX(key string, value interface{}, expire int) bool
	GET(key string) string
	APPEND(key string, value interface{}) int
	DECR(key string) int
	DECRBY(key string, decrement interface{}) int
	GETBIT(key string, offset interface{}) int
	SETBIT(key string, offset, value interface{}) int
	GETRANGE(key string, start, end interface{}) string
	GETSET(key string, value interface{}) string
	INCR(key string) int
	INCRBY(key string, decrement interface{}) int
	INCRBYFLOAT(key string, decrement interface{}) float64
	MGET(keys ...string) []string
	MSET(mapKv map[string]interface{}) bool
	MSETNX(mapKv map[string]interface{}) int
	SETRANGE(key string, offset, value interface{}) int
	STRLEN(key string) int
	RedisKeyer
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

// GET key
// Time complexity: O(1)
// @return Bulk string reply: the value of key, or nil when key does not exist.
func (j *jredis) GET(key string) string {
	res, err := j.exec("GET", j.getKey(key))
	s, _ := redis.String(res, err)
	return s
}

//	If key already exists and is a string, this command appends the value at the end of the string.
//	If key does not exist it is created and set as an empty string, so APPEND will be similar to SET in this special case.
// 	APPEND key value
// @return Integer reply: the length of the string after the append operation.
func (j *jredis) APPEND(key string, value interface{}) int {
	res, err := j.exec("APPEND", j.getKey(key), value)
	n, _ := redis.Int(res, err)
	return n
}

// DECR key
// Time complexity: O(1)
// @return Integer reply: the value of key after the decrement
func (j *jredis) DECR(key string) int {
	res, err := j.exec("DECR", j.getKey(key))
	n, _ := redis.Int(res, err)
	return n
}

// DECRBY key decrement
// Time complexity: O(1)
// Decrements the number stored at key by decrement. If the key does not exist, it is set to 0 before performing the operation.
// An error is returned if the key contains a value of the wrong type or contains a string that can not be represented as integer.
// This operation is limited to 64 bit signed integers.
// @return Integer reply: the value of key after the decrement
func (j *jredis) DECRBY(key string, decrement interface{}) int {
	res, err := j.exec("DECRBY", j.getKey(key), decrement)
	n, _ := redis.Int(res, err)
	return n
}

// GETBIT key offset
// Time complexity: O(1)
// @return Integer reply: the bit value stored at offset.
func (j *jredis) GETBIT(key string, offset interface{}) int {
	res, err := j.exec("GETBIT", j.getKey(key), offset)
	n, _ := redis.Int(res, err)
	return n
}

// SETBIT key offset value
// Time complexity: O(1)
// Sets or clears the bit at offset in the string value stored at key.
// The bit is either set or cleared depending on value, which can be either 0 or 1.
func (j *jredis) SETBIT(key string, offset, value interface{}) int {
	res, err := j.exec("SETBIT", j.getKey(key), offset, value)
	n, _ := redis.Int(res, err)
	return n
}

// GETRANGE key start end
// Time complexity: O(N) where N is the length of the returned string. The complexity is ultimately determined by the returned length,
// but because creating a substring from an existing string is very cheap, it can be considered O(1) for small strings.
// @return Bulk string reply
func (j *jredis) GETRANGE(key string, start, end interface{}) string {
	res, err := j.exec("GETRANGE", j.getKey(key), start, end)
	s, _ := redis.String(res, err)
	return s
}

// GETSET key value
// Time complexity: O(1)
// Atomically sets key to value and returns the old value stored at key.
// Returns an error when key exists but does not hold a string value.
func (j *jredis) GETSET(key string, value interface{}) string {
	res, err := j.exec("GETSET", j.getKey(key), value)
	s, _ := redis.String(res, err)
	return s
}

// INCR key
// Time complexity: O(1)
func (j *jredis) INCR(key string) int {
	res, err := j.exec("INCR", j.getKey(key))
	n, _ := redis.Int(res, err)
	return n
}

// INCRBY key increment
// Time complexity: O(1)
func (j *jredis) INCRBY(key string, decrement interface{}) int {
	res, err := j.exec("INCRBY", j.getKey(key), decrement)
	n, _ := redis.Int(res, err)
	return n
}

// INCRBYFLOAT key increment
// Time complexity: O(1)
func (j *jredis) INCRBYFLOAT(key string, decrement interface{}) float64 {
	res, err := j.exec("INCRBYFLOAT", j.getKey(key), decrement)
	n, _ := redis.Float64(res, err)
	return n
}

// MGET key [key ...]
// Time complexity: O(N) where N is the number of keys to retrieve.
func (j *jredis) MGET(keys ...string) []string {
	arr := make([]interface{}, len(keys))
	for k, v := range keys {
		arr[k] = j.getKey(v)
	}
	res, err := j.exec("MGET", arr...)
	ret, _ := redis.Strings(res, err)
	return ret
}

// MSET key value [key value ...]
// Time complexity: O(N) where N is the number of keys to retrieve.
func (j *jredis) MSET(mapKv map[string]interface{}) bool {
	arr := make([]interface{}, 2*len(mapKv))
	index := 0
	for k, v := range mapKv {
		arr[index] = j.getKey(k)
		index++
		arr[index] = v
		index++
	}
	res, err := j.exec("MSET", arr...)
	return j.isOk(res, err)
}

// MSET key value [key value ...]
// Time complexity: O(N) where N is the number of keys to retrieve.
func (j *jredis) MSETNX(mapKv map[string]interface{}) int {
	arr := make([]interface{}, 2*len(mapKv))
	index := 0
	for k, v := range mapKv {
		arr[index] = j.getKey(k)
		index++
		arr[index] = v
		index++
	}
	res, err := j.exec("MSETNX", arr...)
	n, _ := redis.Int(res, err)
	return n
}

// SETRANGE key offset value
// @return Integer reply: the length of the string after it was modified by the command.
func (j *jredis) SETRANGE(key string, offset, value interface{}) int {
	res, err := j.exec("SETRANGE", j.getKey(key), offset, value)
	n, _ := redis.Int(res, err)
	return n
}

// STRLEN key
// @return Integer reply: the length of the string at key, or 0 when key does not exist.
func (j *jredis) STRLEN(key string) int {
	res, err := j.exec("STRLEN", j.getKey(key))
	n, _ := redis.Int(res, err)
	return n
}
