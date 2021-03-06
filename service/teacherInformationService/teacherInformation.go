package teacherInformationService

import (
	"github.com/360EntSecGroup-Skylar/excelize"
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

func SelectTeacherInformationByName(name string) []model.TeacherInformation {
	teacherInformation := model.TeacherInformation{Name: name}
	return teacherInformation.SelectByNameMore()
}

//分页查询
func SelectByPage(pageNum int, pageSize int, information *model.TeacherInformation) []model.TeacherInformation {
	return information.SelectByPage(pageNum, pageSize)
}

//表格处理
func ProcessingExcelFile(s string) {
	file, _ := excelize.OpenFile(s)
	rows := file.GetRows("Sheet1")
	for i, row := range rows {
		if i != 0 {
			teacherInformation := model.TeacherInformation{}
			for j, colCell := range row {
				switch j {
				case 0:
					teacherInformation.StudentId = colCell
				case 1:
					teacherInformation.Name = colCell
				case 2:
					teacherInformation.Department = colCell
				}
			}
			//插入基本信息
			Insert(&teacherInformation)
		}

	}
}

//插入
func Insert(information *model.TeacherInformation) int64 {
	//设置uuid为主键
	information.Id = uuid.NewV4().String()
	//权限为2
	information.Permissions = 2
	//插入登录信息
	loginService.Registered(information.StudentId, information.StudentId, 2)
	return information.Insert()
}

//更新相关记录权限
func Update(information *model.TeacherInformation) int64 {
	student := model.Student{
		StudentId:   information.StudentId,
		Permissions: information.Permissions}
	student.Update()
	return information.Update()
}

//通过StudentId查询
func SelectTeacherInformationByStudentId(studentId string) *model.TeacherInformation {
	information := model.TeacherInformation{StudentId: studentId}
	return information.SelectByStudentId()
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
	loginService.Delete(information.StudentId)
	return information.Delete()
}
