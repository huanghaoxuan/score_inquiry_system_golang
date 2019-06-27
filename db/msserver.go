package db

import (
	"bufio"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"io"
	"os"
	"strings"
)

/**
 * @Author: HuangHaoXuan
 * @Email: huanghaoxuan1998@outlook.com
 * @github https://github.com/huanghaoxuan
 * @Date: 2019/6/23 14:21
 * @Version 1.0
 */

var DB = Init()

func read() map[string]string {
	var properties = make(map[string]string)
	file, _ := os.Open("db/db.properties")
	defer file.Close()
	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		n := strings.Split(line, "=")
		properties[n[0]] = n[1]
		if err != nil {
			if err == io.EOF {
				//fmt.Println("File read ok!")
				break
			} else {
				//fmt.Println("Read file error!", err)
			}
		}
	}
	return properties
}

//初始化数据库驱动
func Init() *gorm.DB {

	//读取配置文件
	properties := read()
	userName := properties["userName"]
	userPassword := properties["userPassword"]
	//fmt.Println(userName, userPassword)
	db, err := gorm.Open("mssql", "sqlserver://"+userName+":"+userPassword+"@localhost:1433?database=score_inquiry_system")
	if err != nil {
		panic("failed to connect database")
	}
	// 关闭复数表名
	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.AutoMigrate()
	return db
}

//关闭数据库链接
func Close(db *gorm.DB) {
	db.Close()
}
