package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

/**
 * @Author: HuangHaoXuan
 * @Email: huanghaoxuan1998@outlook.com
 * @github https://github.com/huanghaoxuan
 * @Date: 2019/6/23 14:21
 * @Version 1.0
 */

var DB = Init()

//初始化数据库驱动
func Init() *gorm.DB {
	userName := "sa"
	userPassword := "123456"
	fmt.Println(userName, userPassword)
	db, err := gorm.Open("mssql", "sqlserver://"+userName+":"+userPassword+"@localhost:1433?database=score_inquiry_system")
	if err != nil {
		panic("failed to connect database")
	}
	// 关闭复数表名
	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	return db
}

//关闭数据库链接
func Close(db *gorm.DB) {
	db.Close()
}
