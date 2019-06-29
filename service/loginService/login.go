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

func Login(student *model.Student) int64 {
	return student.SelectByStudentIdPassword()
}

func InsertStudent(studentId string, password string) int64 {
	id := uuid.NewV4().String()
	password = md5.GeneratedMD5(password)
	student := model.Student{ID: id, StudentId: studentId, Password: password, Permissions: 1}
	return student.Insert()
}
