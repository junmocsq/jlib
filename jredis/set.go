package jredis

import (
	"github.com/gomodule/redigo/redis"
)

type RedisSetter interface {
	SADD(key string, val ...interface{}) int
	SCARD(key string) int
	SDIFF(keys ...string) []string
	SDIFFSTORE(destination string, keys ...string) int
	SINTER(keys ...string) []string
	SINTERSTORE(destination string, keys ...string) int
	SISMEMBER(key string, member interface{}) int
	SMEMBERS(key string) []string
	SMOVE(source, destination string, member interface{}) int
	SPOP(key string, count ...int) []string
	SRANDMEMBER(key string, count ...int) []string
	SREM(key string, members ...interface{}) int
	SSCAN(key string, count int, pattern ...string) []string
	SUNION(keys ...string) []string
	SUNIONSTORE(destination string, keys ...string) int
	RedisKeyer
}

// SADD key member [member ...]
// Time complexity: O(1) for each element added, so O(N) to add N elements when the command is called with multiple arguments.
// @return Integer reply: the number of elements that were added to the set, not including all the elements already present into the set.
func (j *jredis) SADD(key string, val ...interface{}) int {
	key = j.getKey(key)
	arr := make([]interface{}, 1)
	arr[0] = key
	arr = append(arr, val...)
	n, _ := redis.Int(j.exec("SADD", arr...))
	return n
}

// SCARD key
// @description Returns the set cardinality (number of elements) of the set stored at key.
// Integer reply: the cardinality (number of elements) of the set, or 0 if key does not exist.
func (j *jredis) SCARD(key string) int {
	n, _ := redis.Int(j.exec("SCARD", j.getKey(key)))
	return n
}

// SDIFF key [key ...]
// Time complexity: O(N) where N is the total number of elements in all given sets.
// @description Returns the members of the set resulting from the difference between the first set and all the successive sets.
// @return Integer reply: the number of elements in the resulting set.
// 类似补集
func (j *jredis) SDIFF(keys ...string) []string {
	arr := make([]interface{}, len(keys))
	for i, v := range keys {
		arr[i] = j.getKey(v)
	}
	ret, _ := redis.Strings(j.exec("SDIFF", arr...))
	return ret
}

// SDIFFSTORE destination key [key ...]
// Time complexity: O(N) where N is the total number of elements in all given sets.
// @description This command is equal to SDIFF, but instead of returning the resulting set, it is stored in destination.
// 				If destination already exists, it is overwritten.
// @return Integer reply: the number of elements in the resulting set.
func (j *jredis) SDIFFSTORE(destination string, keys ...string) int {
	arr := make([]interface{}, len(keys)+1)
	arr[0] = j.getKey(destination)
	for i, v := range keys {
		arr[i+1] = j.getKey(v)
	}
	ret, _ := redis.Int(j.exec("SDIFFSTORE", arr...))
	return ret
}

// SINTER key [key ...]
// Time complexity: O(N*M) worst case where N is the cardinality of the smallest set and M is the number of sets.
// @description Returns the members of the set resulting from the intersection of all the given sets.
// @return Integer reply: the number of elements in the resulting set.
// 交集
func (j *jredis) SINTER(keys ...string) []string {
	arr := make([]interface{}, len(keys))
	for i, v := range keys {
		arr[i] = j.getKey(v)
	}
	ret, _ := redis.Strings(j.exec("SINTER", arr...))
	return ret
}

// SINTERSTORE destination key [key ...]
// Time complexity: O(N*M) worst case where N is the cardinality of the smallest set and M is the number of sets.
// @description This command is equal to SINTER, but instead of returning the resulting set, it is stored in destination.
// 				If destination already exists, it is overwritten.
// @return Integer reply: the number of elements in the resulting set.
// 交集
func (j *jredis) SINTERSTORE(destination string, keys ...string) int {
	arr := make([]interface{}, len(keys)+1)
	arr[0] = j.getKey(destination)
	for i, v := range keys {
		arr[i+1] = j.getKey(v)
	}
	ret, _ := redis.Int(j.exec("SINTERSTORE", arr...))
	return ret
}

// SISMEMBER key member
// Time complexity: O(1)
// @description Returns if member is a member of the set stored at key.
// @return Integer reply, specifically:
//			1 if the element is a member of the set.
//			0 if the element is not a member of the set, or if key does not exist.
func (j *jredis) SISMEMBER(key string, member interface{}) int {
	n, _ := redis.Int(j.exec("SISMEMBER", j.getKey(key), member))
	return n
}

// SMEMBERS key
// Time complexity: O(N) where N is the set cardinality.
// @description Returns all the members of the set value stored at key.This has the same effect as running SINTER with one argument key.
// @return Array reply: all elements of the set.
func (j *jredis) SMEMBERS(key string) []string {
	ret, _ := redis.Strings(j.exec("SMEMBERS", j.getKey(key)))
	return ret
}

// SMOVE source destination member
// Time complexity: O(1)
// @description Move member from the set at source to the set at destination.This operation is atomic.
// @return Integer reply, specifically:
//			1 if the element is moved.
//			0 if the element is not a member of source and no operation was performed.
func (j *jredis) SMOVE(source, destination string, member interface{}) int {
	ret, _ := redis.Int(j.exec("SMOVE", j.getKey(source), j.getKey(destination), member))
	return ret
}

// SPOP key [count]
// Time complexity: O(1)
// @description Removes and returns one or more random elements from the set value store at key.
// @return Bulk string reply: the removed element, or nil when key does not exist.
func (j *jredis) SPOP(key string, count ...int) []string {
	c := 1
	if len(count) >= 1 {
		c = count[0]
	}
	ret, _ := redis.Strings(j.exec("SPOP", j.getKey(key), c))
	return ret
}

// SRANDMEMBER key [count]
// Time complexity: Without the count argument O(1), otherwise O(N) where N is the absolute value of the passed count.
// @description When called with just the key argument, return a random element from the set value stored at key.
// @return Return value
//			Bulk string reply: without the additional count argument the command returns a Bulk Reply with the randomly selected element, or nil when key does not exist.
//			Array reply: when the additional count argument is passed the command returns an array of elements, or an empty array when key does not exist.
func (j *jredis) SRANDMEMBER(key string, count ...int) []string {
	c := 1
	if len(count) >= 1 {
		c = count[0]
	}
	ret, _ := redis.Strings(j.exec("SRANDMEMBER", j.getKey(key), c))
	return ret
}

// SREM key member [member ...]
// Time complexity: O(N) where N is the number of members to be removed.
// @return Integer reply: the number of members that were removed from the set, not including non existing members.
func (j *jredis) SREM(key string, members ...interface{}) int {
	arr := make([]interface{}, 1)
	arr[0] = j.getKey(key)
	arr = append(arr, members...)
	ret, _ := redis.Int(j.exec("SREM", arr...))
	return ret
}

// SSCAN key cursor [MATCH pattern] [COUNT count]
// Available since 2.8.0.
// Time complexity: O(1) for every call. O(N) for a complete iteration, including enough command calls for the cursor to return back to 0. N is the number of elements inside the collection..
func (j *jredis) SSCAN(key string, count int, pattern ...string) []string {
	var list []string
	seed := "0"
	patt := ""
	if len(pattern) >= 1 {
		patt = pattern[0]
	}
	for {
		var res interface{}
		var err error
		if patt != "" {
			res, err = j.exec("SSCAN", j.getKey(key), seed, "COUNT", count, "MATCH", patt)
		} else {
			res, err = j.exec("SSCAN", j.getKey(key), seed, "COUNT", count)
		}
		arr, err := redis.Values(res, err)
		if err != nil {
			break
		}
		if s, ok := arr[0].([]uint8); ok {
			seed = string(s)
		}

		tempList, _ := redis.Strings(arr[1], nil)
		list = append(list, tempList...)

		// Starting an iteration with a cursor value of 0, and calling SSCAN
		// until the returned cursor is 0 again is called a full iteration.
		if seed == "0" {
			break
		}
	}
	return list
}

// SUNION key [key ...]
// Time complexity: O(N) where N is the total number of elements in all given sets.
// @description Returns the members of the set resulting from the union of all the given sets.
// @return Array reply: list with members of the resulting set.
func (j *jredis) SUNION(keys ...string) []string {
	arr := make([]interface{}, len(keys))
	for i, v := range keys {
		arr[i] = j.getKey(v)
	}
	ret, _ := redis.Strings(j.exec("SUNION", arr...))
	return ret
}

// SUNIONSTORE destination key [key ...]
// Time complexity: O(N) where N is the total number of elements in all given sets
// @description This command is equal to SUNION, but instead of returning the resulting set, it is stored in destination.
// 				If destination already exists, it is overwritten.
// @return Integer reply: the number of elements in the resulting set.
// 交集
func (j *jredis) SUNIONSTORE(destination string, keys ...string) int {
	arr := make([]interface{}, len(keys)+1)
	arr[0] = j.getKey(destination)
	for i, v := range keys {
		arr[i+1] = j.getKey(v)
	}
	ret, _ := redis.Int(j.exec("SUNIONSTORE", arr...))
	return ret
}
