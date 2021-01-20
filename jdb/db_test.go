package jdb

import (
	"errors"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"gorm.io/gorm"
	"math/rand"
	"testing"
	"time"
)

func TestDb(t *testing.T) {
	SetDbPoolParams("test")
	db := GetDb("test")
	db.AutoMigrate(&User{}, &Topic{}, &Comment{})
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
			//UpdatedAt: time.Now().Unix(),
			//CreatedAt: time.Now().Unix(),
		}, User{
			Name:      "陆雪琪",
			Phone:     "15201400002",
			Signature: "DRY",
			//UpdatedAt: time.Now().Unix(),
			//CreatedAt: time.Now().Unix(),
		}, User{
			Name:      "宋小二",
			Phone:     "15201400003",
			Signature: "KISS",
			//UpdatedAt: time.Now().Unix(),
			//CreatedAt: time.Now().Unix(),
		}, User{
			Name:      "刘阿大",
			Phone:     "15201400004",
			Signature: "陈大宝",
			//UpdatedAt: time.Now().Unix(),
			//CreatedAt: time.Now().Unix(),
		})

		// 添加
		SkipConvey("Insert", func() {
			db.Exec("TRUNCATE test_user")
			d := db.Begin()
			for k, u := range users {
				var res *gorm.DB
				if k%2 == 0 {
					res = d.Select("Name", "Phone", "Signature", "UpdatedAt", "CreatedAt").Create(&u) // 通过数据的指针来创建
				} else {
					res = d.Omit("Signature").Create(&u) // 通过数据的指针来创建
				}
				So(res.Error, ShouldBeNil)
			}
			d.Commit()
			var uu []User
			r := db.Model(&uu).Create([]map[string]interface{}{
				{"Name": "刘小三", "Phone": "15201400005"},
				{"Name": "刘小四", "Phone": "15201400006"},
			})

			uuu := User{Name: "陈小五", Phone: "15201400007"}
			db.Create(&uuu)
			t.Log(r.Error, uu, uuu)
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
			var u User
			db.First(&u)
			u.Name = "张小凡"
			// 更细多列
			ttt := db.Save(u)
			db.First(&u)
			So(ttt.RowsAffected, ShouldEqual, 1)

			rand.Seed(time.Now().Unix())
			// 更新一列
			r1 := db.Model(&User{}).Where("phone >= ?", 15201400005).
				Update("signature", "大葱小时候"+fmt.Sprintf("%d", rand.Int()))
			So(r1.RowsAffected, ShouldEqual, 3)
			r2 := db.Model(&User{}).Where("phone <= ?", 15201400004).
				Update("signature", "小葱长大了"+fmt.Sprintf("%d", rand.Int()))
			So(r2.RowsAffected, ShouldEqual, 4)

			r3 := db.Session(&gorm.Session{DryRun: true}).Model(&u).Update("signature", "张小二").Statement
			// UPDATE `test_user` SET `signature`=?,`updated_at`=? WHERE `id` = ? [张小二 1611140898 1]
			t.Log(r3.SQL.String(), r3.Vars)

			// 更新多列
			r4 := db.Model(&User{}).Where("phone <= ?", 15201400004).
				Updates(User{Signature: "小葱长大了" + fmt.Sprintf("%d", rand.Int()), Sex: 2})
			t.Log(r4.RowsAffected, r4.Statement.Changed("signature"))

			// 更新多列 Select选择字段 Omit忽略字段
			r5 := db.Model(&User{}).Where("phone <= ?", 15201400004).
				Updates(map[string]interface{}{"sex": rand.Intn(10), "signature": "小葱长大了5-" + fmt.Sprintf("%d", rand.Int())})
			t.Log(r5.RowsAffected, r5.Statement.Changed("Sex", "UpdatedAt"))
			// 如果您想在更新时跳过 Hook 方法且不追踪更新时间，可以使用 UpdateColumn、UpdateColumns，其用法类似于 Update、Updates
		})

		SkipConvey("Find", func() {
			var u User
			a := db.First(&u)
			t.Log(u, a.RowsAffected, a.Error, errors.Is(a.Error, gorm.ErrRecordNotFound))
			u.Id = 0
			a = db.Take(&u)
			t.Log(u, a.RowsAffected, a.Error)
			u.Id = 0
			a = db.Last(&u)
			t.Log(u, a.RowsAffected, a.Error)

			b := db.Session(&gorm.Session{DryRun: true}).Take(&u).Statement
			t.Log(b.SQL.String(), b.Vars)
		})
	})
}

type User struct {
	Id        int64  `gorm:"column:id;primaryKey;autoIncrement;<-:create"`
	Name      string `gorm:"uniqueIndex;size:50;not null;default:''"`
	Phone     string `gorm:"uniqueIndex;size:15;not null;default:''"`
	Password  string `gorm:"size:40;not null;default:''"`
	Sex       int8   `gorm:"not null;default:1";json:"sex"`
	Avatar    string `gorm:"size:150;not null;default:''"`
	Signature string `gorm:"size:50;not null;default:''"`
	UpdatedAt int64  `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedAt int64  `gorm:"autoCreateTime;<-:create"`
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
