package jdb

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

type DbAccessor interface {
	FetchOne(ret interface{}) error
	FetchAll(ret interface{}) error
	RowsAffected() (int64, error)
	Insert(model interface{}, fields ...string) error
	Begin() (err error)
	Rollback() (err error)
	Commit() (err error)
	setSqlAndParams(sql string, params []interface{})
}

var dbMap = make(map[string]*gorm.DB)

func getDb(identifier ...string) *gorm.DB {
	id := defaultIdentifier
	if len(identifier) > 0 {
		id = identifier[0]
	}
	db, ok := dbMap[id]
	if !ok {
		panic("请注册 " + id + " 的数据库")
	}
	return db
}

// 获取db
func GetDb(identifier ...string) *gorm.DB {
	if DEBUG {
		return getDb(identifier...).Debug()
	} else {
		return getDb(identifier...)
	}

}

type db struct {
	isSlave    bool
	sql        string
	identifier string
	params     []interface{}
	gormDb     *gorm.DB
	gormTx     *gorm.DB
	isTx       bool
}

func newDb(identifier string, isDebug bool, isSlave ...bool) *db {
	d := &db{
		isSlave:    false,
		identifier: identifier,
	}
	if len(isSlave) > 0 {
		d.isSlave = isSlave[0]
	}
	d.gormDb = getDb(d.identifier).WithContext(context.Background())
	if isDebug {
		d.gormDb = d.gormDb.Debug()
	}
	return d
}
func (d *db) setSqlAndParams(sql string, params []interface{}) {
	d.sql = sql
	d.params = params
}

func (d *db) FetchOne(ret interface{}) error {
	selectDb := d.gormDb
	if d.isTx {
		selectDb = d.gormTx
	}
	if !d.isSlave {
		selectDb = selectDb.Clauses(dbresolver.Write)
	}
	res := selectDb.Raw(d.sql, d.params...).Scan(ret)
	return res.Error
}

func (d *db) FetchAll(ret interface{}) error {
	selectDb := d.gormDb
	if d.isTx {
		selectDb = d.gormTx
	}
	if !d.isSlave {
		selectDb = selectDb.Clauses(dbresolver.Write)
	}
	res := selectDb.Raw(d.sql, d.params...).Scan(ret)
	return res.Error
}

func (d *db) RowsAffected() (int64, error) {
	selectDb := d.gormDb
	if d.isTx {
		selectDb = d.gormTx
	}
	res := selectDb.Exec(d.sql, d.params...)
	return res.RowsAffected, res.Error
}

func (d *db) Insert(model interface{}, fields ...string) error {
	selectDb := d.gormDb
	if d.isTx {
		selectDb = d.gormTx
	}
	if len(fields) > 0 {
		selectDb = selectDb.Select(fields).Create(model)
	} else {
		selectDb = selectDb.Create(model)
	}
	return selectDb.Error
}

func (d *db) Begin() (err error) {
	d.isTx = true
	d.gormTx = d.gormDb.Begin()
	return d.gormTx.Error
}

func (d *db) Rollback() (err error) {
	if !d.isTx {
		panic("please BeginTx!!!")
	}
	r := d.gormTx.Rollback()
	d.isTx = false
	d.gormTx = nil
	return r.Error
}

func (d *db) Commit() (err error) {
	if !d.isTx {
		panic("please BeginTx!!!")
	}
	r := d.gormTx.Commit()
	d.isTx = false
	d.gormTx = nil
	return r.Error
}
