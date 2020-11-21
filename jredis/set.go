package jredis

import "fmt"

type RedisSetter interface {
}

func (j *jredis) SADD(key string, val ...interface{}) {
	key = j.getKey(key)
	arr := make([]interface{}, 1)
	arr[0] = key
	arr = append(arr, val...)
	fmt.Println(arr)
	res, err := j.exec("SADD", arr...)
	fmt.Println("sadd", res, err)
}
