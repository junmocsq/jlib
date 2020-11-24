package jredis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
	"strconv"
)

type RedisZSetter interface {
	ZADD(key string, args ...interface{}) int
	ZCARD(key string) int
	ZCOUNT(key string, min, max float64) int
	ZINCRBY(key string, increment float64, member interface{}) float64
	ZLEXCOUNT(key string, min, max interface{}) int
	ZRANGE(key string, start, stop int) (eles []string, scores []float64)
	ZRANGEBYLEX(key string, min, max interface{}, limit ...interface{}) []string
	ZRANGEBYSCORE(key string, min, max interface{}, limit ...interface{}) (eles []string, scores []float64)
	ZRANK(key string, member interface{}) int
	ZREM(key string, member ...interface{}) int
	ZREMRANGEBYLEX(key string, min, max interface{}) int
	ZREMRANGEBYRANK(key string, start, stop int) int
	ZREMRANGEBYSCORE(key string, min, max interface{}) int
	ZREVRANGE(key string, start, stop int) (eles []string, scores []float64)
	ZREVRANGEBYLEX(key string, max, min interface{}, limit ...interface{}) []string
	ZREVRANGEBYSCORE(key string, max, min interface{}, limit ...interface{}) (eles []string, scores []float64)
	ZREVRANK(key string, member interface{}) int
	ZSCAN(key string, count int, pattern ...string) map[string]string
	ZSCORE(key string, member interface{}) (float64, error)
	RedisKeyer
}

// ZADD key [NX|XX] [GT|LT] [CH] score member [score member ...]
// Time complexity: O(log(N)) for each item added, where N is the number of elements in the sorted set.
// XX: Only update elements that already exist. Never add elements.
// NX: Don't update already existing elements. Always add new elements.
// history
// 		>= 2.4: Accepts multiple elements. In Redis versions older than 2.4 it was possible to add or update a single member per call.
//		>= 3.0.2: Added the XX, NX, CH and INCR options.
//		>=6.2: Added the GT and LT options.
// @return Integer reply The number of elements added to the sorted set, not including elements already existing for which the score was updated.
func (j *jredis) ZADD(key string, args ...interface{}) int {
	arr := make([]interface{}, 1)
	arr[0] = j.getKey(key)
	arr = append(arr, args...)
	n, _ := redis.Int(j.exec("ZADD", arr...))
	return n
}

// ZCARD key
// @description Returns the sorted set cardinality (number of elements) of the sorted set stored at key.
// @return Integer reply: the cardinality (number of elements) of the sorted set, or 0 if key does not exist.
func (j *jredis) ZCARD(key string) int {
	n, _ := redis.Int(j.exec("ZCARD", j.getKey(key)))
	return n
}

// ZCOUNT key min max
// @description Returns the number of elements in the sorted set at key with a score between min and max.
// @return Integer reply: the number of elements in the specified score range.
func (j *jredis) ZCOUNT(key string, min, max float64) int {
	n, _ := redis.Int(j.exec("ZCOUNT", j.getKey(key), min, max))
	return n
}

// ZINCRBY key increment member
// @description Increments the score of member in the sorted set stored at key by increment.
// 				If member does not exist in the sorted set, it is added with increment as its score (as if its previous score was 0.0).
// 				If key does not exist, a new sorted set with the specified member as its sole member is created.
// @return Bulk string reply: the new score of member (a double precision floating point number), represented as string.
func (j *jredis) ZINCRBY(key string, increment float64, member interface{}) float64 {
	res, err := j.exec("ZINCRBY", j.getKey(key), increment, member)
	n, _ := redis.Float64(res, err)
	return n
}

// ZLEXCOUNT key min max
// Time complexity: O(log(N)) with N being the number of elements in the sorted set.
// @description When all the elements in a sorted set are inserted with the same score,
// 				in order to force lexicographical ordering, this command returns the number of elements
// 				in the sorted set at key with a value between min and max.
// @return Integer reply: the number of elements in the specified score range.
func (j *jredis) ZLEXCOUNT(key string, min, max interface{}) int {
	n, _ := redis.Int(j.exec("ZLEXCOUNT", j.getKey(key), min, max))
	return n
}

// ZRANGE key start stop [WITHSCORES]
// @description Returns the specified range of elements in the sorted set stored at key.
// 				The elements are considered to be ordered from the lowest to the highest score.
// 				Lexicographical order is used for elements with equal score.
func (j *jredis) ZRANGE(key string, start, stop int) (eles []string, scores []float64) {
	n, _ := redis.Strings(j.exec("ZRANGE", j.getKey(key), start, stop, "WITHSCORES"))
	length := len(n)
	for i := 0; i < length; i += 2 {
		eles = append(eles, n[i])
		score, err := strconv.ParseFloat(n[i+1], 64)
		if err != nil {
			logrus.WithField("redis", "ZRANGE").Errorf("strconv float64 err:%s", err.Error())
			return nil, nil
		}
		scores = append(scores, score)
	}
	return
}

// ZRANGEBYLEX key min max [LIMIT offset count]
// @description When all the elements in a sorted set are inserted with the same score,
// 				in order to force lexicographical ordering, this command returns all the elements
// 				in the sorted set at key with a value between min and max.
// min max: Valid start and stop must start with ( or [, in order to specify if the range item is respectively exclusive or inclusive.
// 			The special values of + or - for start and stop have the special meaning or positively infinite and negatively infinite strings,
// 			so for instance the command ZRANGEBYLEX myzset - + is guaranteed to return all the elements in the sorted set,
// 			if all the elements have the same score.
// @return Array reply: list of elements in the specified score range.
func (j *jredis) ZRANGEBYLEX(key string, min, max interface{}, limit ...interface{}) []string {
	var res interface{}
	var err error
	if len(limit) == 3 {
		res, err = j.exec("ZRANGEBYLEX", j.getKey(key), min, max, limit[0], limit[1], limit[2])
	} else {
		res, err = j.exec("ZRANGEBYLEX", j.getKey(key), min, max)
	}
	n, _ := redis.Strings(res, err)
	return n
}

// ZRANGEBYSCORE key min max [WITHSCORES] [LIMIT offset count]
// @description Returns all the elements in the sorted set at key with a score between min and max
// 				(including elements with score equal to min or max). The elements are considered to be ordered from low to high scores.
// @usage ZRANGEBYSCORE zset (5 (10		ZRANGEBYSCORE zset -inf +inf
// @return Array reply: list of elements in the specified score range (optionally with their scores).
func (j *jredis) ZRANGEBYSCORE(key string, min, max interface{}, limit ...interface{}) (eles []string, scores []float64) {
	var res interface{}
	var err error
	if len(limit) == 3 {
		res, err = j.exec("ZRANGEBYSCORE", j.getKey(key), min, max, "WITHSCORES", limit[0], limit[1], limit[2])
	} else {
		res, err = j.exec("ZRANGEBYSCORE", j.getKey(key), min, max, "WITHSCORES")
	}
	n, _ := redis.Strings(res, err)
	length := len(n)
	for i := 0; i < length; i += 2 {
		eles = append(eles, n[i])
		score, err := strconv.ParseFloat(n[i+1], 64)
		if err != nil {
			logrus.WithField("redis", "ZRANGEBYSCORE").Errorf("strconv float64 err:%s", err.Error())
			return nil, nil
		}
		scores = append(scores, score)
	}
	return
}

// ZRANK key member
// @description Returns the rank of member in the sorted set stored at key, with the scores ordered from low to high.
// 				The rank (or index) is 0-based, which means that the member with the lowest score has rank 0.
// @return
//			If member exists in the sorted set, Integer reply: the rank of member.
//			If member does not exist in the sorted set or key does not exist, Bulk string reply: nil.
func (j *jredis) ZRANK(key string, member interface{}) int {
	ret, err := redis.Int(j.exec("ZRANK", j.getKey(key), member))
	fmt.Println("ZRANK", ret, err)
	if err == redis.ErrNil {
		// 不存在返回-1
		return -1
	}
	return ret
}

// ZREM key member [member ...]
// Time complexity: O(M*log(N)) with N being the number of elements in the sorted set and M the number of elements to be removed.
// @description Removes the specified members from the sorted set stored at key. Non existing members are ignored.
// @return Integer reply, The number of members removed from the sorted set, not including non existing members.
func (j *jredis) ZREM(key string, member ...interface{}) int {
	arr := make([]interface{}, 1)
	arr[0] = j.getKey(key)
	arr = append(arr, member...)
	ret, _ := redis.Int(j.exec("ZREM", arr...))
	return ret
}

// ZREMRANGEBYLEX key min max
// @description When all the elements in a sorted set are inserted with the same score,
// 				in order to force lexicographical ordering, this command removes all elements
// 				in the sorted set stored at key between the lexicographical range specified by min and max.
// @usage ZREMRANGEBYLEX myzset [alpha [omega
func (j *jredis) ZREMRANGEBYLEX(key string, min, max interface{}) int {
	ret, _ := redis.Int(j.exec("ZREMRANGEBYLEX", j.getKey(key), min, max))
	return ret
}

// ZREMRANGEBYRANK key start stop
// @description Removes all elements in the sorted set stored at key with rank between start and stop.
// 				Both start and stop are 0 -based indexes with 0 being the element with the lowest score.
// 				These indexes can be negative numbers, where they indicate offsets starting at the element
// 				with the highest score. For example: -1 is the element with the highest score,
// 				-2 the element with the second highest score and so forth.
func (j *jredis) ZREMRANGEBYRANK(key string, start, stop int) int {
	ret, _ := redis.Int(j.exec("ZREMRANGEBYRANK", j.getKey(key), start, stop))
	return ret
}

// ZREMRANGEBYSCORE key min max
// @description Removes all elements in the sorted set stored at key with a score between min and max (inclusive).
func (j *jredis) ZREMRANGEBYSCORE(key string, min, max interface{}) int {
	ret, _ := redis.Int(j.exec("ZREMRANGEBYSCORE", j.getKey(key), min, max))
	return ret
}

// ZREVRANGE key start stop [WITHSCORES]
// @description Returns the specified range of elements in the sorted set stored at key.
// 				The elements are considered to be ordered from the highest to the lowest score.
// 				Descending lexicographical order is used for elements with equal score.
func (j *jredis) ZREVRANGE(key string, start, stop int) (eles []string, scores []float64) {
	n, _ := redis.Strings(j.exec("ZREVRANGE", j.getKey(key), start, stop, "WITHSCORES"))
	length := len(n)
	for i := 0; i < length; i += 2 {
		eles = append(eles, n[i])
		score, err := strconv.ParseFloat(n[i+1], 64)
		if err != nil {
			logrus.WithField("redis", "ZREVRANGE").Errorf("strconv float64 err:%s", err.Error())
			return nil, nil
		}
		scores = append(scores, score)
	}
	return
}

// ZREVRANGEBYLEX key max min [LIMIT offset count]
func (j *jredis) ZREVRANGEBYLEX(key string, max, min interface{}, limit ...interface{}) []string {
	var res interface{}
	var err error
	if len(limit) == 3 {
		res, err = j.exec("ZREVRANGEBYLEX", j.getKey(key), max, min, limit[0], limit[1], limit[2])
	} else {
		res, err = j.exec("ZREVRANGEBYLEX", j.getKey(key), max, min)
	}
	n, _ := redis.Strings(res, err)
	return n
}

// ZREVRANGEBYSCORE key max min [WITHSCORES] [LIMIT offset count]
func (j *jredis) ZREVRANGEBYSCORE(key string, max, min interface{}, limit ...interface{}) (eles []string, scores []float64) {
	var res interface{}
	var err error
	if len(limit) == 3 {
		res, err = j.exec("ZREVRANGEBYSCORE", j.getKey(key), max, min, "WITHSCORES", limit[0], limit[1], limit[2])
	} else {
		res, err = j.exec("ZREVRANGEBYSCORE", j.getKey(key), max, min, "WITHSCORES")
	}
	n, _ := redis.Strings(res, err)
	length := len(n)
	for i := 0; i < length; i += 2 {
		eles = append(eles, n[i])
		score, err := strconv.ParseFloat(n[i+1], 64)
		if err != nil {
			logrus.WithField("redis", "ZREVRANGEBYSCORE").Errorf("strconv float64 err:%s", err.Error())
			return nil, nil
		}
		scores = append(scores, score)
	}
	return
}

// ZREVRANK key member
// @description Returns the rank of member in the sorted set stored at key, with the scores ordered from high to low.
// 				The rank (or index) is 0-based, which means that the member with the highest score has rank 0.
func (j *jredis) ZREVRANK(key string, member interface{}) int {
	ret, err := redis.Int(j.exec("ZREVRANK", j.getKey(key), member))
	if err == redis.ErrNil {
		// 不存在返回-1
		return -1
	}
	return ret
}

// ZSCAN key cursor [MATCH pattern] [COUNT count]
func (j *jredis) ZSCAN(key string, count int, pattern ...string) map[string]string {
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
			res, err = j.exec("ZSCAN", j.getKey(key), seed, "COUNT", count, "MATCH", patt)
		} else {
			res, err = j.exec("ZSCAN", j.getKey(key), seed, "COUNT", count)
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

// ZSCORE key member
// @description Returns the score of member in the sorted set at key.
// @return Bulk string reply: the score of member (a double precision floating point number), represented as string.
func (j *jredis) ZSCORE(key string, member interface{}) (float64, error) {
	return redis.Float64(j.exec("ZSCORE", j.getKey(key), member))
}
