package main

//gorm demo

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model	//内嵌结构体
	Name 		string
	Age 		int64
}

func main() {
	//连接MySQL数据库
	db, err := gorm.Open("mysql", "root:123456@(192.168.31.182:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("open database failed:%v", err)
		panic(err)
	}
	defer db.Close()

	//创建表 自动迁移（把结构体和数据表进行对应）
	db.AutoMigrate(&User{})

	// 创建
	//u1 := User{Name: "zhangsan", Age: 18}
	//db.Create(&u1)
	//u2 := User{Name: "lisi", Age: 22}
	//db.Create(&u2)

	//查询
	var user User
	db.First(&user)
	fmt.Printf("user:%#v\n", user)

	var users []User
	db.Debug().Find(&users)
	fmt.Printf("users:%#v\n", users)

	var user2 User
	db.Where("age > ?", 20).First(&user2)
	fmt.Printf("user:%#v\n", user2)

	//更新
	user.Name = "biubiu"
	user.Age = 99
	//db.Debug().Save(&user)	//默认修改所有字段
	db.Debug().Model(&user).Update("name", "lalala") //修改指定字段，默认的updated_at也会更新
	db.Debug().Model(&user2).UpdateColumn("age", 66) //只修改指定字段，默认的updated_at不更新

	//让所有users表中用户age+2
	db.Model(&User{}).Update("age", gorm.Expr("age+?", 2))

	//删除，修改delete_at字段的软删除
	db.Debug().Where("name=?", "lalala").Delete(&User{})

}
