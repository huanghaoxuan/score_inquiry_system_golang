package service

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
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
	if information.SelectByStudentID(information.StudentID).ID != "" {
		return 0
	}
	//设置uuid为主键
	information.ID = uuid.NewV4().String()
	//默认权限为1
	information.Permissions = 1
	return information.Insert()
}

//更新相关记录权限
func Update(information *model.StudentInformation) int64 {
	informationOld := information.SelectByStudentID(information.StudentID)
	fmt.Println(informationOld)
	information.ID = informationOld.ID
	fmt.Println(information)
	return information.Update()
}
