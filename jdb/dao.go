package jdb

import (
	"errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Dao interface {
	Debug() Dao
	SetSlave() Dao
	SetTag(tag string) Dao
	SetKey(key string) Dao
	PrepareSql(sql string, args ...interface{}) Dao
	ClearCache() bool
	FetchOne(ret interface{}) (err error)
	FetchAll(ret interface{}) (err error)
	RowsAffected() (int64, error) // update,delete,replace into
	Insert(model interface{}, fields ...string) error
	Begin() (err error)
	Rollback() (err error)
	Commit() (err error)
}

type dao struct {
	identifier string
	sql        string
	params     []interface{}
	tag        string
	key        string
	tagKey     string
	isSlave    bool
	isCache    bool
	isDebug    bool
	db         DbAccessor
}

func NewDao(identifier ...string) Dao {
	d := &dao{}
	if len(identifier) > 0 {
		d.identifier = identifier[0]
	} else {
		d.identifier = defaultIdentifier
	}
	d.isDebug = DEBUG
	return d
}

func (d *dao) SetSlave() Dao {
	d.isSlave = true
	return d
}
func (d *dao) Debug() Dao {
	d.isDebug = true
	return d
}

func (d *dao) SetTag(tag string) Dao {
	d.isCache = true
	d.tag = tag
	return d
}

func (d *dao) SetKey(key string) Dao {
	d.key = key
	d.tagKey = d.tag + "_" + d.key
	return d
}

func (d *dao) PrepareSql(sql string, args ...interface{}) Dao {
	d.sql = sql
	d.params = args
	if d.isCache {
		d.tagKey = d.genSqlCacheKey()
	}
	return d
}

func (d *dao) getCacheKey() string {
	// 设置key 返回key 否则返回tag
	str := d.tag
	if str == "" {
		panic("请设置cache tag")
	}
	if d.key != "" {
		return str + "_" + d.key
	}
	return str
}

func (d *dao) ClearCache() bool {
	defer d.clear()
	return cacheAccess.Delete(d.getCacheKey())
}

func (d *dao) genSqlCacheKey() string {
	tagNum := cacheAccess.Get(d.tag)
	if tagNum == "" {
		tagNum = genRandstr()
		cacheAccess.Set(d.tag, tagNum, expire)
	}
	return hash(tagNum, d.sql, d.params)
}

func (d *dao) FetchOne(ret interface{}) (err error) {
	defer d.clear()
	if d.isCache {
		// 查询cache
		cacheRes := cacheAccess.Get(d.tagKey)
		if cacheRes == Empty {
			return
		} else if cacheRes != "" {
			e := JsonDecode(cacheRes, ret)
			if e != nil {
				logrus.WithField("dao", "JsonDecode").Errorf("data:%#v err:%s", ret, e.Error())
			}
			return
		}
	}
	if d.db == nil {
		d.db = newDb(d.identifier, d.isDebug, d.isSlave)
	}
	d.db.setSqlAndParams(d.sql, d.params)
	err = d.db.FetchOne(ret)
	if d.isCache {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			cacheAccess.Set(d.tagKey, Empty, expire)
		}
		if ret != nil {
			jres, e := JsonEncode(ret)
			if e == nil {
				cacheAccess.Set(d.tagKey, jres, expire)
			} else {
				logrus.WithField("dao", "JsonEncode").Errorf("data:%#v err:%s", ret, e.Error())
			}
		}
	}
	return
}

func (d *dao) FetchAll(ret interface{}) (err error) {
	defer d.clear()
	if d.isCache {
		// 查询cache
		cacheRes := cacheAccess.Get(d.tagKey)
		if cacheRes == Empty {
			return
		} else if cacheRes != "" {
			e := JsonDecode(cacheRes, ret)
			if e != nil {
				logrus.WithField("dao", "JsonDecode").Errorf("data:%#v err:%s", ret, e.Error())
			}
			return
		}
	}
	if d.db == nil {
		d.db = newDb(d.identifier, d.isDebug, d.isSlave)
	}
	d.db.setSqlAndParams(d.sql, d.params)
	err = d.db.FetchAll(ret)
	if d.isCache {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			cacheAccess.Set(d.tagKey, Empty, expire)
		}
		if ret != nil {
			jres, e := JsonEncode(ret)
			if e == nil {
				cacheAccess.Set(d.tagKey, jres, expire)
			} else {
				logrus.WithField("dao", "JsonEncode").Errorf("data:%#v err:%s", ret, e.Error())
			}
		}
	}
	return
}

func (d *dao) RowsAffected() (int64, error) {
	defer d.clear()
	if d.db == nil {
		d.db = newDb(d.identifier, d.isDebug, d.isSlave)
	}
	d.db.setSqlAndParams(d.sql, d.params)
	n, e := d.db.RowsAffected()
	if d.isCache {
		cacheAccess.Delete(d.getCacheKey())
	}
	return n, e
}

func (d *dao) Insert(model interface{}, fields ...string) error {
	defer d.clear()
	if d.db == nil {
		d.db = newDb(d.identifier, d.isDebug, d.isSlave)
	}
	d.db.setSqlAndParams(d.sql, d.params)
	e := d.db.Insert(model, fields...)
	if e == nil && d.isCache {
		cacheAccess.Delete(d.getCacheKey())
	}
	return e
}

func (d *dao) Begin() (err error) {
	if d.db == nil {
		d.db = newDb(d.identifier, d.isDebug, d.isSlave)
	}
	return d.db.Begin()
}

func (d *dao) Rollback() (err error) {
	return d.db.Rollback()
}

func (d *dao) Commit() (err error) {
	return d.db.Commit()
}

func (d *dao) clear() {
	d.identifier = ""
	d.sql = ""
	d.params = nil
	d.tag = ""
	d.key = ""
	d.tagKey = ""
	d.isSlave = false
	d.isCache = false
}
