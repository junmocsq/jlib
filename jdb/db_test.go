package jdb

import (
	"database/sql"
	"testing"
)

func TestDb(t *testing.T) {
	RegisterSqlDb("test", false, "root:123456@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	db := GetDb("test")
	db.AutoMigrate(&User{})
	t.Log(sql.Drivers())

	//sql := "UPDATE user set age=age+1 WHERE id=?"
	//sql = "select * from user where id>?"
	//params := []interface{}{1}
	//d := GetDb("test")
	//t.Log(d.BeginTx())
	//d.Update()
	//d.RollBack()
	//d.Update()
	//d.RollBack()

	//d.FetchOne()
}

type User struct {
	UserId             int64 `gorm:"primarykey"`
	UserName           string
	UserPsw            string
	UserStatus         int8
	UserDisplay        int8
	UserNickname       string
	UserEmail          string
	UserQq             string
	UserTel            string
	UserSex            int
	UserBirthday       int
	UserHead           string
	UserSign           string
	UserTicketAct      int
	UserGold2          float64
	UserGold2Frozen    float64
	UserTicketRec      int
	UserIscharge       int8
	UserExp            int
	UserBrowse         int
	UserLogintime      int
	UserRegtime        int
	UserRegip          string
	UserDisplaytime    int
	UserNicknameBefore string
	IsEditor           int8
	UserAuth           int8
	DeleteTime         int
	RefreshToken       string
	NationCode         int
	Credit             int
}

func (User) TableName() string {
	return "wm_user_user"
}
func ExampleRegisterSqlDb() {

}
