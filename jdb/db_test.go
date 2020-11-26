package jdb

import (
	"database/sql"
	"testing"
)

func TestDb(t *testing.T) {
	RegisterSqlDb("root:123456@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local", "test", false)
	getMaster("test")
	t.Log(sql.Drivers())

	sql := "UPDATE user set age=age+1 WHERE id=?"
	sql = "select * from user where id>?"
	params := []interface{}{1}
	d := NewDb(sql, "test", params)
	//t.Log(d.BeginTx())
	//d.Update()
	//d.RollBack()
	//d.Update()
	//d.RollBack()

	d.FetchOne()
}

func ExampleRegisterSqlDb() {

}
