package main

//gorm demo

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type UserInfo struct {
	ID uint
	Name string
	Gender string
	Hobby string
}

type User struct {
	gorm.Model	//内嵌结构体
	Name 		string
	Age 		sql.NullInt64	//零值类型
	Birthday 	*time.Time
	Email 		string 		`gorm:"type:varchar(100);unique_index"`
	Role 		string 		`gorm:"size:255"`	//设置字段大小
	Membernumber *string 	`gorm:"unique;not null"` //gorm中类型零值会直接滤掉，插入零值需要使用指针类型或者sql.NullString
	Num 		int 		`gorm:"AUTO_INCREMENT"`
	Address 	string 		`gorm:"index:addr"` //给address字段创建addr的索引
	Ignore 		int 		`gorm:"-"`	//忽略本字段
}

func main() {
	//连接MySQL数据库
	db, err := gorm.Open("mysql", "root:123456@(192.168.31.182:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("open database failed:%v", err)
		panic(err)
	}
	defer db.Close()

	////创建表 自动迁移（把结构体和数据表进行对应）
	//db.AutoMigrate(&UserInfo{})
	//
	//// 创建数据行
	//u1 := UserInfo{ID:1, Name:"zhangsan", Gender:"男", Hobby: "蛙泳"}
	//db.Create(&u1)
	//// 查询
	//var u UserInfo
	//db.First(&u)
	//fmt.Printf("u:%#v\n", u)
	//// 更新
	//db.Model(&u).Update("hobby", "篮球")
	//// 删除
	//db.Delete(&u)

	db.AutoMigrate(&User{})
}
