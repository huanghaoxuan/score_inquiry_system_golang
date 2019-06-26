package service

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"os"
	"score_inquiry_system/model"
)

/**
 * @Author: HuangHaoXuan
 * @Email: huanghaoxuan1998@outlook.com
 * @github https://github.com/huanghaoxuan
 * @Date: 2019/6/23 20:32
 * @Version 1.0
 */

//插入
func Insert(information *model.StudentInformation) int64 {
	if information.SelectByStudentId(information.StudentId).Id != "" {
		return 0
	}
	//设置uuid为主键
	information.Id = uuid.NewV4().String()
	//默认权限为1
	information.Permissions = 1
	return information.Insert()
}

//更新相关记录权限
func Update(information *model.StudentInformation) int64 {
	informationOld := information.SelectByStudentId(information.StudentId)
	fmt.Println(informationOld)
	information.Id = informationOld.Id
	fmt.Println(information)
	return information.Update()
}

//处理上传的学生信息表格
func ProcessingExcelFile(file os.File) {

}
