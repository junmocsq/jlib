package jdb

import (
	"errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Dao interface {
	SetTag(tag string) Dao
	SetKey(key string) Dao
	PrepareSql(sql string, args ...interface{}) Dao
	ClearCache() bool
	FetchOne(ret interface{}) (err error)
	FetchAll(ret interface{}) (err error)
	RowsAffected() (int64, error) // update,delete,replace into
	Insert(model interface{}, fields ...string) error
	Begin() (err error)
	RollBack() (err error)
	Commit() (err error)
}

type dao struct {
	dbname  string
	sql     string
	params  []interface{}
	tag     string
	key     string
	tagKey  string
	isSlave bool
	isCache bool
	db      DbAccessor
}

func NewDao(dbname string, isSlave ...bool) Dao {
	d := &dao{}
	d.dbname = dbname
	slave := false
	if len(isSlave) > 0 {
		slave = isSlave[0]
	}
	d.isSlave = slave
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
	d.db.setSqlAndParams(d.sql, d.params)
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
		d.db = newDb(d.dbname, d.isSlave)
	}
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
		d.db = newDb(d.dbname, d.isSlave)
	}
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
		d.db = newDb(d.dbname, d.isSlave)
	}
	n, e := d.db.RowsAffected()
	if d.isCache {
		cacheAccess.Delete(d.getCacheKey())
	}
	return n, e
}

func (d *dao) Insert(model interface{}, fields ...string) error {
	defer d.clear()
	if d.db == nil {
		d.db = newDb(d.dbname, d.isSlave)
	}
	e := d.db.Insert(model, fields...)
	if e == nil && d.isCache {
		cacheAccess.Delete(d.getCacheKey())
	}
	return e
}

func (d *dao) Begin() (err error) {
	if d.db == nil {
		d.db = newDb(d.dbname, d.isSlave)
	}
	return d.db.Begin()
}

func (d *dao) RollBack() (err error) {
	return d.db.RollBack()
}

func (d *dao) Commit() (err error) {
	return d.db.Commit()
}

func (d *dao) clear() {
	d.dbname = ""
	d.sql = ""
	d.params = nil
	d.tag = ""
	d.key = ""
	d.tagKey = ""
	d.isSlave = false
	d.isCache = false
}
