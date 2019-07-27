package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

var DB *gorm.DB
func InitDataBase(){
	//连接数据库
	db,err:=gorm.Open("mysql",os.Getenv("MYSQL_DSN"))
	if err!=nil{
		panic(err)
	}
	//连接池设置
	//空闲
	db.DB().SetMaxIdleConns(20)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(30)
	//打开日志
	db.LogMode(true)
	DB=db
	//迁移
	migration()
}
