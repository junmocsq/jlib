package jdb

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"math/rand"
	"time"
)

func genRandstr() string {
	n := time.Now().UnixNano()
	rand.Seed(n)
	return fmt.Sprintf("%d%d", n, rand.Intn(100000))
}

func hash(tagNum, sql string, params []interface{}) string {
	s := fmt.Sprintf("%s-%s-%#v", tagNum, sql, params)
	fmt.Println(s)
	data := []byte(s)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func JsonEncode(data interface{}) (string, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	s, e := json.Marshal(&data)
	return string(s), e
}

func JsonDecode(str string, data interface{}) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal([]byte(str), data)
}

func JsonEncode2(data interface{}) (string, error) {
	s, e := json.Marshal(&data)
	return string(s), e
}

func JsonDecode2(str string, data interface{}) error {
	return json.Unmarshal([]byte(str), data)
}
