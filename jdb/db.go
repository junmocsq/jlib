package jdb

import (
	"context"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"time"
)

type DbAccessor interface {
	FetchOne(ret interface{}) error
	FetchAll(ret interface{}) error
	RowsAffected() (int64, error)
	Insert(model interface{}, fields ...string) error
	Begin() (err error)
	RollBack() (err error)
	Commit() (err error)
	setSqlAndParams(sql string, params []interface{})
}

var dbMap = make(map[string]*gorm.DB)

// 注册数据库
// params ConnMaxLifetime(s) MaxOpenConns MaxIdleConns
// "mysql", "root:123456@tcp(localhost:3306)/test"
func RegisterSqlDb(identifier string, isSlave bool, dsn ...string) {
	if !isSlave {
		db, err := gorm.Open(mysql.Open(dsn[0]), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		dbMap[identifier] = db
	} else {
		db, ok := dbMap[identifier]
		if !ok {
			panic("请设置 " + identifier + " 的主库")
		}
		var arr []gorm.Dialector
		for _, v := range dsn {
			arr = append(arr, mysql.Open(v))
		}
		db.Use(dbresolver.Register(dbresolver.Config{
			Replicas: arr,
			// sources/replicas 负载均衡策略
			Policy: dbresolver.RandomPolicy{},
		}))
	}
}

// ConnMaxLifetime(s) ConnMaxIdleTime(s) MaxOpenConns MaxIdleConns
func SetDbPoolParams(identifier string, params ...int) {
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
	setting := []int{180, 90, 100, 100}
	if len(params) > 4 {
		panic("RegisterSqlDb params is error!")
	}
	for k, v := range params {
		setting[k] = v
	}
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(setting[0]))
	sqlDB.SetConnMaxIdleTime(time.Second * time.Duration(setting[1]))
	sqlDB.SetMaxOpenConns(setting[2])
	sqlDB.SetMaxIdleConns(setting[3])
}

// 获取db
func GetDb(identifier string) *gorm.DB {
	db, ok := dbMap[identifier]
	if !ok {
		panic("请注册 " + identifier + " 的数据库")
	}
	return db
}

type db struct {
	isSlave bool
	sql     string
	dbname  string
	params  []interface{}
	gormDb  *gorm.DB
	gormTx  *gorm.DB
	isTx    bool
}

func newDb(dbname string, isSlave ...bool) *db {
	d := &db{
		isSlave: false,
		dbname:  dbname,
	}
	if len(isSlave) > 0 {
		d.isSlave = isSlave[0]
	}
	d.gormDb = GetDb(d.dbname).WithContext(context.Background())
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

func (d *db) RollBack() (err error) {
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
	r := d.gormTx.Rollback()
	d.isTx = false
	d.gormTx = nil
	return r.Error
}
