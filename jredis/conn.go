package jredis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

type GRPool struct {
}

func initRedis(module, host, port, auth string) *redis.Pool {
	fmt.Println("init redis ", module, "pool")
	pool := &redis.Pool{
		MaxIdle:     256,               // 最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态。
		MaxActive:   2000,              // 最大的连接数，表示同时最多有N个连接。0表示不限制。
		IdleTimeout: 240 * time.Second, // 当连接空闲超过这个时间就回收
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				"tcp",
				host+":"+port,
				redis.DialReadTimeout(time.Second),
				redis.DialWriteTimeout(time.Second),
				redis.DialConnectTimeout(time.Second),
				redis.DialDatabase(0),
				redis.DialPassword(auth),
			)
		},
	}
	return pool
}

func getPool(module string) *redis.Pool {
	if c, ok := redisPool[module]; ok {
		return c
	}
	panic(module + " 不存在")
}

func exec(cmd string, module string, key string, args ...interface{}) (interface{}, error) {
	con := getPool(module).Get()
	if err := con.Err(); err != nil {
		return nil, err
	}
	defer con.Close()
	params := make([]interface{}, 0)
	params = append(params, key)
	if len(args) > 0 {
		for _, v := range args {
			params = append(params, v)
		}
	}
	return con.Do(cmd, params...)
}

func getClient(module string) (redis.Conn, error) {
	con := getPool(module).Get()
	if err := con.Err(); err != nil {
		return nil, err
	}
	return con, nil
}
