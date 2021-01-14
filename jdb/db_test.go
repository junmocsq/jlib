package jdb

import (
	//"database/sql"
	"fmt"
	"testing"
)

func TestDb(t *testing.T) {
	RegisterCacheAccesser("127.0.0.1", "6379", "default")
	RegisterSqlDb("test", false, "root:123456@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	SetDbPoolParams("test")
	db := GetDb("test")
	db.AutoMigrate(&User{})
	//t.Log(sql.Drivers())

	d := NewDao("test")
	d.Begin()
	var u = User{
		Id:        0,
		Name:      "junmo",
		Phone:     "12345678901",
		Password:  "123444",
		Sex:       2,
		Avatar:    "",
		Signature: "",
		UpdatedAt: 0,
		CreatedAt: 0,
	}
	fields := []string{"name", "phone", "password", "sex"}
	d.SetTag(u.TableName()).Insert(&u, fields...)
	fmt.Println(u.Id)
	s := "SELECT * FROM User WHERE id=?"
	d.PrepareSql(s, u.Id).FetchOne(&u)
	fmt.Println(u)
}

type User struct {
	Id        int64  `json:"id";gorm:"column:id"`
	Name      string `json:"name";gorm:"uniqueIndex"`
	Phone     string `json:"phone";gorm:"uniqueIndex"`
	Password  string `json:"password"`
	Sex       int8   `json:"sex"`
	Avatar    string `json:"avatar"`
	Signature string `json:"signature";gorm:"index:idx_phone,length:10"`
	UpdatedAt int64  `json:"updated_at"`
	CreatedAt int64  `json:"created_at"`
}

func (User) TableName() string {
	return "us_user"
}
func ExampleRegisterSqlDb() {

}
