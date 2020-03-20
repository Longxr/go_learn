package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *gorm.DB
)

func InitMySQL()(err error) {
	//连接MySQL数据库
	DB, err = gorm.Open("mysql", "root:123456@(192.168.31.182:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("open database failed:%v", err)
		panic(err)
	}
	//defer DB.Close()
	return DB.DB().Ping()
}

func Close() {
	DB.Close()
}