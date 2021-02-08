package jdb

import (
	"context"
	"github.com/junmocsq/jlib/jredis"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"time"
)

var (
	expire      = 300
	cacheAccess cacheAccessor
	Empty       = "nil"
	DEBUG       = false

	redisModule = "sql"
	redisClient = jredis.NewRedis(redisModule)

	defaultIdentifier = ""
)

func RegisterCacheAccessor(host, port, auth string, debug ...bool) {
	jredis.RegisterRedisPool(host, port, jredis.ModuleConf(redisModule), jredis.AuthConf(auth), jredis.PrefixConf(redisModule))
	if len(debug) > 0 {
		jredis.SetDebug(debug[0])
	}
	cacheAccess = newCache()
}

func SetDebug(debug bool) {
	DEBUG = debug
}

// isDefault 是否为默认数据库
// ConnMaxLifetime(s) ConnMaxIdleTime(s) MaxOpenConns MaxIdleConns
type conf struct {
	identifier                 string
	isSlave                    bool
	dsn                        string
	isDefault                  bool
	maxLifetime                time.Duration
	maxIdleTime                time.Duration
	maxOpenConns, maxIdleConns int
}

type Option func(*conf)

func ConnMaxLifetime(second int) Option {
	return func(i *conf) {
		i.maxLifetime = time.Duration(second) * time.Second
	}
}

func ConnMaxIdleTime(second int) Option {
	return func(i *conf) {
		i.maxIdleTime = time.Duration(second) * time.Second
	}
}

func ConnMaxOpenConns(conns int) Option {
	return func(i *conf) {
		i.maxOpenConns = conns
	}
}

func ConnMaxIdleConns(conns int) Option {
	return func(i *conf) {
		i.maxIdleConns = conns
	}
}

func IsDefault(b bool) Option {
	return func(i *conf) {
		i.isDefault = b
	}
}

// 注册数据库
// identifier
// isDefault 是否为默认数据库
// ConnMaxLifetime(s) ConnMaxIdleTime(s) MaxOpenConns MaxIdleConns
// // "mysql", "root:123456@tcp(localhost:3306)/test"
func RegisterMasterDb(identifier string, dsn string, options ...Option) {
	c := &conf{
		identifier:   identifier,
		isSlave:      true,
		dsn:          dsn,
		isDefault:    false,
		maxLifetime:  180 * time.Second,
		maxIdleTime:  90 * time.Second,
		maxOpenConns: 30,
		maxIdleConns: 30,
	}

	for _, option := range options {
		option(c)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	dbMap[identifier] = db
	if defaultIdentifier == "" { // 没有默认数据库，则设置此为默认数据库
		defaultIdentifier = identifier
	}

	if c.isDefault {
		defaultIdentifier = identifier
	}
	db, ok := dbMap[identifier]
	if !ok {
		panic("请设置 " + identifier + " 的主库")
	}
	sqlDB, _ := db.DB()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := sqlDB.PingContext(ctx); err != nil {
		logrus.WithField("sql", "ping").Infof("%s  status:down!!!", identifier)
		return
	}

	// ConnMaxLifetime(s) ConnMaxIdleTime(s) MaxOpenConns MaxIdleConns
	// MaxOpenConns MaxIdleConns 建议相同
	sqlDB.SetConnMaxLifetime(c.maxLifetime)
	sqlDB.SetConnMaxIdleTime(c.maxIdleTime)
	sqlDB.SetMaxOpenConns(c.maxOpenConns)
	sqlDB.SetMaxIdleConns(c.maxIdleConns)

}

func RegisterSlaveDb(identifier string, dsns []string) {
	db, ok := dbMap[identifier]
	if !ok {
		panic("请设置 " + identifier + " 的主库")
	}
	var arr []gorm.Dialector
	for _, v := range dsns {
		arr = append(arr, mysql.Open(v))
	}
	db.Use(dbresolver.Register(dbresolver.Config{
		Replicas: arr,
		// sources/replicas 负载均衡策略
		Policy: dbresolver.RandomPolicy{},
	}))
}
