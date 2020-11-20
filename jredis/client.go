package jredis

import "github.com/gomodule/redigo/redis"

type jredis struct {
	module string
}

func NewRedis(module ...string) *jredis {
	if len(module) >= 1 {
		return &jredis{module: module[0]}
	}
	return &jredis{module: defaultModule}
}

func (j *jredis) Client() (redis.Conn, error) {
	return getClient(j.module)
}

func (j *jredis) GetKey(key string) string {
	return j.getKey(key)
}

func (j *jredis) EXEC(conn redis.Conn, cmd string, args ...interface{}) (interface{}, error) {
	return conn.Do(cmd, args...)
}

func (j *jredis) exec(cmd string, args ...interface{}) (interface{}, error) {
	con := getPool(j.module).Get()
	if err := con.Err(); err != nil {
		return nil, err
	}
	defer con.Close()
	return con.Do(cmd, args...)
}

// 格式化redis结果
func (j *jredis) isOk(res interface{}, err error) bool {
	if ok, err := redis.String(res, err); err == nil && ok == "OK" {
		return true
	}
	return false
}

func (j *jredis) getKey(key string) string {
	return getKey(j.module, key)
}
