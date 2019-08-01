package loginService

import (
	uuid "github.com/satori/go.uuid"
	"score_inquiry_system/model"
	"score_inquiry_system/util/md5"
)

/**
 * @Author: HuangHaoXuan
 * @Email: huanghaoxuan1998@outlook.com
 * @github https://github.com/huanghaoxuan
 * @Date: 2019/6/25 21:17
 * @Version 1.0
 */

func Login(student *model.Student) (int64, int) {
	return student.SelectByStudentIdPassword()
}

func Registered(studentId string, password string, permissions int) int64 {
	id := uuid.NewV4().String()
	password = md5.GeneratedMD5(password)
	student := model.Student{Id: id, StudentId: studentId, Password: password, Permissions: permissions}
	return student.Insert()
}

func FirstAdmin(name string, studentId string, password string, permissions int) int64 {
	student := model.Student{
		Id:          uuid.NewV4().String(),
		StudentId:   studentId,
		Password:    md5.GeneratedMD5(password),
		Permissions: permissions}

	teacherInformation := model.TeacherInformation{
		Id:          uuid.NewV4().String(),
		Name:        name,
		StudentId:   studentId,
		Permissions: permissions}

	teacherInformation.Insert()
	return student.Insert()
}

//删除一条记录
func Delete(studentId string) int64 {
	student := model.Student{StudentId: studentId}
	return student.Delete()
}
