package jredis

import "fmt"

func (j *jredis) SADD(key string, val ...interface{}) {
	key = GetKey(j.module, key)
	arr := make([]interface{},1)
	arr[0] = key
	arr = append(arr,val...)
	fmt.Println(arr)
	res, err := j.exec("SADD", arr...)
	fmt.Println("sadd",res, err)
}
