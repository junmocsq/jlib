package jdb

type cacheAccesser interface {
	Set(key, value string, expire int) bool
	Get(key string) string
	Delete(key string) bool
	Expire(key string, expire int) bool
}

func newCache() cacheAccesser {
	return &cache{}
}

type cache struct {
}

func (c *cache) Set(key, value string, expire int) bool {
	if expire == 0 {
		return redisClient.SET(key, value)
	} else {
		return redisClient.SETEX(key, value, expire)
	}

}

func (c *cache) Get(key string) string {
	return redisClient.GET(key)
}

func (c *cache) Delete(key string) bool {
	res := redisClient.DEL(key)
	if res == 0 {
		return false
	}
	return true
}

func (c *cache) Expire(key string, expire int) bool {
	res := redisClient.EXPIRE(key, expire)
	if res == 0 {
		return false
	}
	return true
}
