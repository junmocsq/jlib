package jdb

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"time"
)

type DbAccesser interface {
	FetchOne()
	FetchAll()
	Update()
	Insert()
}

var dbPool = make(map[string][]*sql.DB)

// 注册数据库
// params ConnMaxLifetime(s) MaxOpenConns MaxIdleConns
// "mysql", "root:123456@tcp(localhost:3306)/test"
func RegisterSqlDb(dsn, dbname string, isSlave bool, params ...int) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := sqlDB.PingContext(ctx); err != nil {
		logrus.WithField("sql", "ping").Infof("%s  status:down!!!", dbname)
		return
	}

	// ConnMaxLifetime(s) MaxOpenConns MaxIdleConns
	setting := []int{180, 100, 10}
	if len(params) > 3 {
		panic("RegisterSqlDb params is error!")
	}
	for k, v := range params {
		setting[k] = v
	}
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(setting[0]))
	sqlDB.SetMaxOpenConns(setting[1])
	sqlDB.SetMaxIdleConns(setting[2])
	dbPool[getDbPoolKey(dbname, isSlave)] = append(dbPool[getDbPoolKey(dbname, isSlave)], sqlDB)
}

// 获取db库key
func getDbPoolKey(dbname string, isSlave bool) string {
	if isSlave {
		return dbname + "_0"
	} else {
		return dbname + "_1"
	}
}

func getMaster(dbname string) *sql.DB {
	if r, ok := dbPool[getDbPoolKey(dbname, false)]; ok {
		length := len(r)
		if length == 0 {
			panic("dbname:" + dbname + " master is not set!")
		}
		index := rand.Intn(length)
		return r[index]
	}
	panic("dbname:" + dbname + " master is not set!")

}

func getSlave(dbname string) *sql.DB {
	if r, ok := dbPool[getDbPoolKey(dbname, true)]; ok {
		length := len(r)
		if length == 0 {
			return getMaster(dbname)
		}
		index := rand.Intn(length)
		return r[index]
	}
	return getMaster(dbname)
}

// 获取db pool
func getDb(dbname string, isSlave bool) *sql.DB {
	if isSlave {
		return getSlave(dbname)
	} else {
		return getMaster(dbname)
	}
}

type db struct {
	isSlave bool
	sql     string
	dbname  string
	params  []interface{}
	sqlDb   *sql.DB
	sqlTx   *sql.Tx
	isTx    bool
	ctx     context.Context
}

func NewDb(sql, dbname string, params []interface{}, isSlave ...bool) *db {
	d := &db{
		isSlave: false,
		sql:     sql,
		dbname:  dbname,
		params:  params,
	}
	if len(isSlave) > 0 {
		d.isSlave = isSlave[0]
	}
	d.sqlDb = getDb(d.dbname, d.isSlave)
	d.ctx = context.Background()
	return d
}

func (d *db) Update() (int64, error) {
	if d.isTx {
		return d.updateTx()
	} else {
		return d.update()
	}
}

func (d *db) update() (int64, error) {
	conn, err := d.sqlDb.Conn(d.ctx)
	if err != nil {
		logrus.WithField("sql", "conn").Errorf("update conn err:%s", err.Error())
		return 0, err
	}
	defer conn.Close() // Return the connection to the pool.
	result, err := conn.ExecContext(d.ctx, d.sql, d.params...)
	if err != nil {
		logrus.WithField("sql", "ExecContext").Errorf("update ExecContext err:%s", err.Error())
		return 0, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		logrus.WithField("sql", "RowsAffected").Errorf("update RowsAffected err:%s", err.Error())
		return 0, err
	}
	return rows, nil
}

func (d *db) updateTx() (int64, error) {
	result, err := d.sqlTx.ExecContext(d.ctx, d.sql, d.params...)
	if err != nil {
		logrus.WithField("sql", "ExecContext").Errorf("update ExecContext err:%s", err.Error())
		return 0, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		logrus.WithField("sql", "RowsAffected").Errorf("update RowsAffected err:%s", err.Error())
		return 0, err
	}
	return rows, nil
}

func (d *db) Insert() (int64, error) {
	if d.isTx {
		return d.updateTx()
	} else {
		return d.update()
	}
}

func (d *db) insert() (int64, error) {
	conn, err := d.sqlDb.Conn(d.ctx)
	if err != nil {
		logrus.WithField("sql", "conn").Errorf("Insert conn err:%s", err.Error())
		return 0, err
	}
	defer conn.Close() // Return the connection to the pool.
	result, err := conn.ExecContext(d.ctx, d.sql, d.params...)
	if err != nil {
		logrus.WithField("sql", "ExecContext").Errorf("Insert ExecContext err:%s", err.Error())
		return 0, err
	}
	rows, err := result.LastInsertId()
	if err != nil {
		logrus.WithField("sql", "LastInsertId").Errorf("Insert LastInsertId err:%s", err.Error())
		return 0, err
	}
	return rows, nil
}

func (d *db) insertTx() (int64, error) {
	result, err := d.sqlTx.ExecContext(d.ctx, d.sql, d.params...)
	if err != nil {
		logrus.WithField("sql", "ExecContext").Errorf("Insert ExecContext err:%s", err.Error())
		return 0, err
	}
	rows, err := result.LastInsertId()
	if err != nil {
		logrus.WithField("sql", "LastInsertId").Errorf("Insert LastInsertId err:%s", err.Error())
		return 0, err
	}
	return rows, nil
}

func (d *db) BeginTx() (err error) {
	d.isTx = true
	d.sqlTx, err = d.sqlDb.BeginTx(d.ctx, nil)
	return
}

func (d *db) RollBack() (err error) {
	if !d.isTx {
		panic("please BeginTx!!!")
	}
	err = d.sqlTx.Rollback()
	d.isTx = false
	return
}

func (d *db) Commit() (err error) {
	if !d.isTx {
		panic("please BeginTx!!!")
	}
	err = d.sqlTx.Commit()
	d.isTx = false
	return
}

func (d *db) FetchOne() (int64, error) {
	conn, err := d.sqlDb.Conn(d.ctx)
	if err != nil {
		logrus.WithField("sql", "conn").Errorf("Insert conn err:%s", err.Error())
		return 0, err
	}
	defer conn.Close() // Return the connection to the pool.
	rows, err := conn.QueryContext(d.ctx, d.sql, d.params...)
	if err != nil {
		logrus.WithField("sql", "ExecContext").Errorf("Insert ExecContext err:%s", err.Error())
		return 0, err
	}
	defer rows.Close()
	fmt.Println(rows.Columns())
	fmt.Println(rows.ColumnTypes())

	names := make([]int, 0)
	for rows.Next() {
		var id int
		var name int
		if err := rows.Scan(&id, &name); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			log.Fatal(err)
		}
		names = append(names, id)
	}
	fmt.Println(names)
	return 0, nil
}
