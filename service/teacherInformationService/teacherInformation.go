package teacherInformationService

import (
	uuid "github.com/satori/go.uuid"
	"score_inquiry_system/model"
	"score_inquiry_system/service/loginService"
)

/**
 * @Author: HuangHaoXuan
 * @Email: huanghaoxuan1998@outlook.com
 * @github https://github.com/huanghaoxuan
 * @Date: 2019/6/23 20:32
 * @Version 1.0
 */

//获取页数
func Count(information *model.TeacherInformation) int {
	return information.Count()
}

//分页查询
func SelectByPage(pageNum int, pageSize int, information *model.TeacherInformation) []model.TeacherInformation {
	return information.SelectByPage(pageNum, pageSize)
}

//插入
func Insert(information *model.TeacherInformation) int64 {
	//设置uuid为主键
	information.Id = uuid.NewV4().String()
	//权限为2
	information.Permissions = 2
	//插入登录信息

	loginService.Registered(information.TeacherId, information.TeacherId, 2)
	return information.Insert()
}

//更新相关记录权限
func Update(information *model.TeacherInformation) int64 {
	return information.Update()
}

//通过id查询
func SelectTeacherInformationById(id string) *model.TeacherInformation {
	information := model.TeacherInformation{Id: id}
	return information.SelectById()
}

//删除一条记录
func Delete(id string) int64 {
	//获取相关记录，获取学号
	information := SelectTeacherInformationById(id)
	loginService.Delete(information.TeacherId)
	return information.Delete()
}
