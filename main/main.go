package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	Name   string
	Gender bool
	Age    int
	//给结构体添加 id(主键)、创建时间、更新时间、删除时间
	gorm.Model
}

func main() {
	//建立数据库连接
	dsn := "root:ayxqmhy@0411@tcp(127.0.0.1:3306)/ginlearning?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	//自动迁移，可以创建含有 Student 结构体字段表属性的数据库表
	db.AutoMigrate(&Student{})
	/*	与已有数据库连接
			sqlDB, err := sql.Open("mysql", "ginlearning")
			gormDB, err := gorm.Open(mysql.New(mysql.Config{
		  		Conn: sqlDB,
			}), &gorm.Config{})
	*/

	//创建数据
	s1 := Student{
		Name:   "zzc",
		Gender: true,
		Age:    18,
	}

	//只添加选中的属性名
	//db.Select("Name", "Gender", "Age").Create(&s1)

	//忽略添加选中的属性名
	//db.Omit("Age").Create(&s1)
	db.Create(&s1)
}
