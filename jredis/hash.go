package jredis

import (
	"github.com/gomodule/redigo/redis"
)

type RedisHasher interface {
	HDEL(key string, fields ...interface{}) int
	HEXISTS(key string, field interface{}) int
	HGET(key string, field interface{}) string
	HGETALL(key string) map[string]string
	HINCRBY(key string, field interface{}, increment int) int
	HINCRBYFLOAT(key string, field interface{}, increment float64) float64
	HKEYS(key string) []string
	HLEN(key string) int
	HMGET(key string, fields ...interface{}) []string
	HMSET(key string, mapKv map[interface{}]interface{}) bool
	HSCAN(key string, count int, pattern ...string) map[string]string
	HSET(key string, field, value interface{}) int
	HSETNX(key string, field, value interface{}) int
	HSTRLEN(key string, field interface{}) int
	HVALS(key string) []string
	RedisKeyer
}

// HDEL key field [field ...]
// Time complexity: O(N) where N is the number of fields to be removed.
// @return Integer reply: the number of fields that were removed from the hash, not including specified but non existing fields.
func (j *jredis) HDEL(key string, fields ...interface{}) int {
	arr := make([]interface{}, 1)
	arr[0] = j.getKey(key)
	arr = append(arr, fields...)
	ret, _ := redis.Int(j.exec("HDEL", arr...))
	return ret

}

// HEXISTS key field
// Time complexity: O(1)
// @return Integer reply, specifically:
//
//	1 if the hash contains field.
//	0 if the hash does not contain field, or key does not exist.
func (j *jredis) HEXISTS(key string, field interface{}) int {
	ret, _ := redis.Int(j.exec("HEXISTS", j.getKey(key), field))
	return ret
}

// HGET key field
// @return Bulk string reply: the value associated with field,
// or nil when field is not present in the hash or key does not exist.
func (j *jredis) HGET(key string, field interface{}) string {
	ret, _ := redis.String(j.exec("HGET", j.getKey(key), field))
	return ret
}

// HGETALL key
// Time complexity: O(N) where N is the size of the hash.
// @return Array reply: list of fields and their values stored in the hash, or an empty list when key does not exist.
func (j *jredis) HGETALL(key string) map[string]string {
	ret, _ := redis.StringMap(j.exec("HGETALL", j.getKey(key)))
	return ret
}

// HINCRBY key field increment
// Time complexity: O(1)
// @return Integer reply: the value at field after the increment operation.
func (j *jredis) HINCRBY(key string, field interface{}, increment int) int {
	ret, _ := redis.Int(j.exec("HINCRBY", j.getKey(key), field, increment))
	return ret
}

// HINCRBYFLOAT key field increment
// Time complexity: O(1)
// @return Bulk string reply: the value of field after the increment.
func (j *jredis) HINCRBYFLOAT(key string, field interface{}, increment float64) float64 {
	ret, _ := redis.Float64(j.exec("HINCRBYFLOAT", j.getKey(key), field, increment))
	return ret
}

// HKEYS key
// Time complexity: O(N) where N is the size of the hash
// @return Array reply: list of fields in the hash, or an empty list when key does not exist.
func (j *jredis) HKEYS(key string) []string {
	ret, _ := redis.Strings(j.exec("HKEYS", j.getKey(key)))
	return ret
}

// HLEN key
// Time complexity: O(1)
// @return Integer reply: number of fields in the hash, or 0 when key does not exist.
func (j *jredis) HLEN(key string) int {
	ret, _ := redis.Int(j.exec("HLEN", j.getKey(key)))
	return ret
}

// HMGET key field [field ...]
// Time complexity: O(N) where N is the number of fields being requested.
// @return Array reply: list of values associated with the given fields, in the same order as they are requested.
func (j *jredis) HMGET(key string, fields ...interface{}) []string {
	arr := make([]interface{}, 1)
	arr[0] = j.getKey(key)
	arr = append(arr, fields...)
	ret, _ := redis.Strings(j.exec("HMGET", arr...))
	return ret
}

// HMSET key field value [field value ...]
// Time complexity: O(N) where N is the number of fields being set.
// @return Simple string reply
func (j *jredis) HMSET(key string, mapKv map[interface{}]interface{}) bool {
	arr := make([]interface{}, 2*len(mapKv)+1)
	arr[0] = j.getKey(key)
	index := 1
	for k, v := range mapKv {
		arr[index] = k
		index++
		arr[index] = v
		index++
	}
	return j.isOk(j.exec("HMSET", arr...))
}

// HSCAN key cursor [MATCH pattern] [COUNT count]
// Available since 2.8.0.
// Time complexity: O(1) for every call. O(N) for a complete iteration, including enough command calls for the cursor to return back to 0. N is the number of elements inside the collection..
func (j *jredis) HSCAN(key string, count int, pattern ...string) map[string]string {
	seed := "0"
	patt := ""
	if len(pattern) >= 1 {
		patt = pattern[0]
	}
	ret := make(map[string]string)
	for {
		var res interface{}
		var err error
		if patt != "" {
			res, err = j.exec("HSCAN", j.getKey(key), seed, "COUNT", count, "MATCH", patt)
		} else {
			res, err = j.exec("HSCAN", j.getKey(key), seed, "COUNT", count)
		}
		arr, err := redis.Values(res, err)
		if err != nil {
			break
		}
		if s, ok := arr[0].([]uint8); ok {
			seed = string(s)
		}

		tempMap, _ := redis.StringMap(arr[1], nil)
		for k, v := range tempMap {
			ret[k] = v
		}

		// Starting an iteration with a cursor value of 0, and calling SSCAN
		// until the returned cursor is 0 again is called a full iteration.
		if seed == "0" {
			break
		}
	}
	return ret
}

// HSET key field value [field value ...]
// Time complexity: O(1) for each field/value pair added, so O(N) to add N field/value pairs when the command is called with multiple field/value pairs.
// @description Sets field in the hash stored at key to value. If key does not exist, a new key holding a hash is created.
//
//	If field already exists in the hash, it is overwritten.
//	As of Redis 4.0.0, HSET is variadic and allows for multiple field/value pairs.
//
// @return Integer reply: The number of fields that were added.
func (j *jredis) HSET(key string, field, value interface{}) int {
	ret, _ := redis.Int(j.exec("HSET", j.getKey(key), field, value))
	return ret
}

// HSETNX key field value
// Time complexity: O(1)
// @description Sets field in the hash stored at key to value, only if field does not yet exist.
//
//	If key does not exist, a new key holding a hash is created.
//	If field already exists, this operation has no effect.
//
// @return Integer reply, specifically:
//
//	1 if field is a new field in the hash and value was set.
//	0 if field already exists in the hash and no operation was performed.
func (j *jredis) HSETNX(key string, field, value interface{}) int {
	ret, _ := redis.Int(j.exec("HSETNX", j.getKey(key), field, value))
	return ret
}

// HSTRLEN key field
// Time complexity: O(1)
// @description Returns the string length of the value associated with field in the hash stored at key.
//
//	If the key or the field do not exist, 0 is returned.
//
// @return Integer reply: the string length of the value associated with field,
//
//	or zero when field is not present in the hash or key does not exist at all.
func (j *jredis) HSTRLEN(key string, field interface{}) int {
	ret, _ := redis.Int(j.exec("HSTRLEN", j.getKey(key), field))
	return ret
}

// HVALS key
// Time complexity: O(N) where N is the size of the hash.
// @return Array reply: list of values in the hash, or an empty list when key does not exist.
func (j *jredis) HVALS(key string) []string {
	ret, _ := redis.Strings(j.exec("HVALS", j.getKey(key)))
	return ret
}
