package jdb

import (
	. "github.com/smartystreets/goconvey/convey"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestDb(t *testing.T) {
	SetDbPoolParams("test")
	db := GetDb("test")
	Convey("DB", t, func() {
		SkipConvey("AutoMigrate", func() {
			db.Exec(" drop table if exists test_user")
			err := db.AutoMigrate(&User{}, &Topic{}, &Comment{})
			So(err, ShouldBeNil)
		})
		var users []User
		users = append(users, User{
			Name:      "张小凡",
			Phone:     "15201400001",
			Signature: "最坚固的城堡也是最坚固的死牢",
			UpdatedAt: time.Now().Unix(),
			CreatedAt: time.Now().Unix(),
		}, User{
			Name:      "陆雪琪",
			Phone:     "15201400002",
			Signature: "DRY",
			UpdatedAt: time.Now().Unix(),
			CreatedAt: time.Now().Unix(),
		}, User{
			Name:      "宋小二",
			Phone:     "15201400003",
			Signature: "KISS",
			UpdatedAt: time.Now().Unix(),
			CreatedAt: time.Now().Unix(),
		}, User{
			Name:      "刘阿大",
			Phone:     "15201400004",
			Signature: "陈大宝",
			UpdatedAt: time.Now().Unix(),
			CreatedAt: time.Now().Unix(),
		})

		// 添加
		Convey("Insert", func() {
			db.Exec("TRUNCATE test_user")
			d := db.Begin()
			for k, u := range users {
				var res *gorm.DB
				if k%2 == 0 {
					res = d.Select("Name", "Phone", "Signature", "CreatedAt").Create(&u) // 通过数据的指针来创建
				} else {
					res = d.Omit("UpdatedAt", "Signature", "CreatedAt").Create(&u) // 通过数据的指针来创建
				}
				So(res.Error, ShouldBeNil)
			}
			d.Commit()
			var uu []User
			r := db.Model(uu).Create([]map[string]interface{}{
				{"Name": "刘小三", "Phone": "15201400005"},
				{"Name": "刘小四", "Phone": "15201400006"},
			})
			t.Log(r.Error, uu)
		})

		SkipConvey("BatchInsert", func() {
			db.Exec("TRUNCATE test_user")
			db.Select("Name", "Phone", "Signature", "CreatedAt").Create(users)

			db.Exec("TRUNCATE test_user")
			db.Select("Name", "Phone", "Signature", "CreatedAt").CreateInBatches(users, 2)
			r := db.Session(&gorm.Session{DryRun: true}).
				Select("Name", "Phone", "Signature", "CreatedAt").Create(users).Statement
			t.Log("sql:", r.SQL.String(), r.Vars)
			t.Log(users)
		})

		Convey("Update", func() {
			//err := db.AutoMigrate(&User{}, &Topic{}, &Comment{})
			//So(err, ShouldBeNil)
		})

		Convey("Find", func() {
			//err := db.AutoMigrate(&User{}, &Topic{}, &Comment{})
			//So(err, ShouldBeNil)
		})
	})
}

type User struct {
	Id        int64  `gorm:"column:id"`
	Name      string `gorm:"uniqueIndex;size:50;not null;default:''"`
	Phone     string `gorm:"uniqueIndex;size:15;not null;default:''"`
	Password  string `gorm:"size:40;not null;default:''"`
	Sex       int8   `gorm:"not null;default:1";json:"sex"`
	Avatar    string `gorm:"size:150;not null;default:''"`
	Signature string `gorm:"size:50;not null;default:''"`
	UpdatedAt int64  `gorm:"not null;default:0";json:"updated_at"`
	CreatedAt int64  `gorm:"not null;default:0"`
}

func (User) TableName() string {
	return "test_user"
}

func (User) BeforeSave(tx *gorm.DB) (err error) {

	return nil
}
func (User) BeforeCreate(tx *gorm.DB) (err error) {

	return nil
}
func (User) AfterCreate(tx *gorm.DB) (err error) {

	return nil
}
func (User) AfterSave(tx *gorm.DB) (err error) {

	return nil
}

type Topic struct {
	Id        int64
	Title     string `gorm:"size:50"`
	Uid       int64  `gorm:"index"`
	Content   string `gorm:"size:500"`
	Images    string `gorm:"size:1500"`
	UpdatedAt int64  `json:"updated_at"`
	CreatedAt int64  `json:"created_at"`
}

func (Topic) TableName() string {
	return "test_topic"
}

type Comment struct {
	Id        int64
	Tid       int64  `gorm:"index"`
	Uid       int64  `gorm:"index"`
	Content   string `gorm:"size:500"`
	CreatedAt int64
}

func (Comment) TableName() string {
	return "test_comment"
}

func ExampleRegisterSqlDb() {

}
