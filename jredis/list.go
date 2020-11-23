package jredis

import (
	"github.com/gomodule/redigo/redis"
)

type RedisLister interface {
	BLPOP(timeout int, keys ...string) []string
	BRPOP(timeout int, keys ...string) []string
	BRPOPLPUSH(source, destination string, timeout int) string
	LINDEX(key string, index int) string
	LINSERTAFTER(key string, privot, value interface{}) int
	LINSERTBEFORE(key string, privot, value interface{}) int
	LLEN(key string) int
	LPOP(key string) string
	LPUSH(key string, values ...interface{}) int
	LPUSHX(key string, values ...interface{}) int
	LRANGE(key string, start, stop int) []string
	LREM(key string, count int, value interface{}) int
	LSET(key string, index int, value interface{}) bool
	LTRIM(key string, start, stop int) bool
	RPOP(key string) string
	RPOPLPUSH(source, destination string) string
	RPUSH(key string, values ...interface{}) int
	RPUSHX(key string, values ...interface{}) int
	RedisKeyer
}

// BLPOP key [key ...] timeout
func (j *jredis) BLPOP(timeout int, keys ...string) []string {
	arr := make([]interface{}, len(keys)+1)
	for index, key := range keys {
		arr[index] = j.getKey(key)
	}
	arr[len(keys)] = timeout
	res, err := j.exec("BLPOP", arr...)
	ret, err := redis.Strings(res, err)
	if len(ret) == 2 {
		ret[0] = j.trimPrefixKey(ret[0])
	}
	return ret
}

// BRPOP key [key ...] timeout
func (j *jredis) BRPOP(timeout int, keys ...string) []string {

	arr := make([]interface{}, len(keys)+1)
	for index, key := range keys {
		arr[index] = j.getKey(key)
	}
	arr[len(keys)] = timeout
	res, err := j.exec("BRPOP", arr...)
	ret, err := redis.Strings(res, err)
	if len(ret) == 2 {
		ret[0] = j.trimPrefixKey(ret[0])
	}
	return ret
}

// BRPOPLPUSH source destination timeout
// 从source右探出一个元素从左进入destination，阻塞时间为timeout，0为永久阻塞
func (j *jredis) BRPOPLPUSH(source, destination string, timeout int) string {
	res, err := j.exec("BRPOPLPUSH", j.getKey(source), j.getKey(destination), timeout)
	str, err := redis.String(res, err)
	return str
}

// LINDEX key index
// @return Bulk string reply: the requested element, or nil when index is out of range.
func (j *jredis) LINDEX(key string, index int) string {
	res, err := j.exec("LINDEX", j.getKey(key), index)
	str, err := redis.String(res, err)
	return str
}

// LINSERT key BEFORE|AFTER pivot element
// @return Integer reply: the length of the list after the insert operation, or -1 when the value pivot was not found.
func (j *jredis) LINSERTAFTER(key string, privot, value interface{}) int {
	return j.lINSERT(key, "AFTER", privot, value)
}

func (j *jredis) LINSERTBEFORE(key string, privot, value interface{}) int {
	return j.lINSERT(key, "BEFORE", privot, value)
}

func (j *jredis) lINSERT(key string, position, privot, value interface{}) int {
	res, err := j.exec("lINSERT", j.getKey(key), position, privot, value)
	n, _ := redis.Int(res, err)
	return n
}

// @return integer reply: the length of the list at key.
func (j *jredis) LLEN(key string) int {
	res, err := j.exec("LLEN", j.getKey(key))
	n, _ := redis.Int(res, err)
	return n
}

// @return Bulk string reply: the value of the first element, or nil when key does not exist.
func (j *jredis) LPOP(key string) string {
	res, err := j.exec("LPOP", j.getKey(key))
	str, _ := redis.String(res, err)
	return str
}

// LPUSH key element [element ...]
// @return Integer reply: the length of the list after the push operations.
func (j *jredis) LPUSH(key string, values ...interface{}) int {
	arr := make([]interface{}, 1)
	arr[0] = j.getKey(key)
	arr = append(arr, values...)
	res, err := j.exec("LPUSH", arr...)
	n, _ := redis.Int(res, err)
	return n
}

// LPUSHX key element [element ...]
// @description Inserts specified values at the head of the list stored at key, only if key already exists and holds a list.
// @return Bulk string reply: the value of the first element, or nil when key does not exist.
func (j *jredis) LPUSHX(key string, values ...interface{}) int {
	arr := make([]interface{}, 1)
	arr[0] = j.getKey(key)
	arr = append(arr, values...)
	res, err := j.exec("LPUSHX", arr...)
	n, _ := redis.Int(res, err)
	return n
}

// LRANGE key start stop
// @return Array reply: list of elements in the specified range.
func (j *jredis) LRANGE(key string, start, stop int) []string {
	res, err := j.exec("LRANGE", j.getKey(key), start, stop)
	arr, _ := redis.Strings(res, err)
	return arr
}

// LREM key count element
// @return Integer reply: the number of removed elements
// count > 0 : 从表头开始向表尾搜索，移除与 VALUE 相等的元素，数量为 COUNT 。
// count < 0 : 从表尾开始向表头搜索，移除与 VALUE 相等的元素，数量为 COUNT 的绝对值。
// count = 0 : 移除表中所有与 VALUE 相等的值。
func (j *jredis) LREM(key string, count int, value interface{}) int {
	res, err := j.exec("LREM", j.getKey(key), count, value)
	n, _ := redis.Int(res, err)
	return n
}

// LSET key index element
// @return An error is returned for out of range indexes.
func (j *jredis) LSET(key string, index int, value interface{}) bool {
	res, err := j.exec("LSET", j.getKey(key), index, value)
	return j.isOk(res, err)
}

// LTRIM key start stop
// @description Trim an existing list so that it will contain only the specified range of elements specified.
// 列表剪裁
func (j *jredis) LTRIM(key string, start, stop int) bool {
	res, err := j.exec("LTRIM", j.getKey(key), start, stop)
	return j.isOk(res, err)
}

// @return Bulk string reply: the value of the first element, or nil when key does not exist.
func (j *jredis) RPOP(key string) string {
	res, err := j.exec("RPOP", j.getKey(key))
	str, _ := redis.String(res, err)
	return str
}

// RPOPLPUSH source destination
// 从source右弹出一个元素从左进入destination
// @return Bulk string reply: the element being popped and pushed.
func (j *jredis) RPOPLPUSH(source, destination string) string {
	res, err := j.exec("RPOPLPUSH", j.getKey(source), j.getKey(destination))
	str, err := redis.String(res, err)
	return str
}

// RPUSH key element [element ...]
// @return Integer reply: the length of the list after the push operations.
func (j *jredis) RPUSH(key string, values ...interface{}) int {
	arr := make([]interface{}, 1)
	arr[0] = j.getKey(key)
	arr = append(arr, values...)
	res, err := j.exec("RPUSH", arr...)
	n, _ := redis.Int(res, err)
	return n
}

// LPUSHX key element [element ...]
// @description Inserts specified values at the head of the list stored at key, only if key already exists and holds a list.
// @return Bulk string reply: the value of the first element, or nil when key does not exist.
func (j *jredis) RPUSHX(key string, values ...interface{}) int {
	arr := make([]interface{}, 1)
	arr[0] = j.getKey(key)
	arr = append(arr, values...)
	res, err := j.exec("RPUSHX", arr...)
	n, _ := redis.Int(res, err)
	return n
}
