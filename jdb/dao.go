package jdb

type Dao interface {
	SetDb(dbname string) Dao
	SetTag(tag string) Dao
	SetKey(key string) Dao
	PrepareSql(sql string, args ...interface{}) Dao
	ClearCache() bool
	FetchOne()
	FetchAll()
	Insert()
	Update()
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
}

func NewDao() *dao {
	return &dao{}
}
func (d *dao) SetDb(dbname string) Dao {
	d.dbname = dbname
	return d
}

func (d *dao) SetSlave() Dao {
	d.isSlave = true
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
	d.tagKey = d.genSqlCacheKey()
	return d
}

func (d *dao) getCacheKey() string {
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

func (d *dao) FetchOne() {
	defer d.clear()
	if d.isCache {
		// 查询cache
		cacheRes := cacheAccess.Get(d.tagKey)
		if cacheRes == Empty {
			return
		} else if cacheRes != "" {
			return
		}
	}
	dbAccess.FetchOne()
	if d.isCache {
		cacheAccess.Set(d.tagKey, Empty, expire)
	}
	return
}

func (d *dao) FetchAll() {
	defer d.clear()
	if d.isCache {
		// 查询cache
		cacheRes := cacheAccess.Get(d.tagKey)
		if cacheRes == Empty {
			return
		} else if cacheRes != "" {
			return
		}
	}
	dbAccess.FetchAll()
	if d.isCache {
		cacheAccess.Set(d.tagKey, Empty, expire)
	}
	return
}

func (d *dao) Insert() {
	defer d.clear()
	dbAccess.Insert()
	if d.isCache {
		cacheAccess.Delete(d.getCacheKey())
	}
}

func (d *dao) Update() {
	defer d.clear()
	dbAccess.Update()
	if d.isCache {
		cacheAccess.Delete(d.getCacheKey())
	}
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
