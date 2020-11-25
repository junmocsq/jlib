package jdb

import (
	"crypto/md5"
	"fmt"
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
	data := []byte(s)
	return fmt.Sprintf("%x", md5.Sum(data))
}
