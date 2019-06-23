package service

import (
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
func Insert(information *model.StudentInformation) {
	//设置uuid为主键
	information.ID = uuid.NewV4().String()
	//默认权限为1
	information.Permissions = 1
	model.Insert(information)
}

//更新相关记录权限
func Update(information *model.StudentInformation) {
	informationOld := model.SelectByStudentID(information.StudentID)
	information.ID = informationOld.ID
	model.Update(information)
}
