package db

import (
	"bufio"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
	db, err := gorm.Open("mysql", userName+":"+userPassword+"@/score_inquiry_system?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("数据库连接错误:", err)
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
