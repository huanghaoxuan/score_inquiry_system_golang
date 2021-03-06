package model

import (
	"score_inquiry_system/db"
)

/**
 * @Author: HuangHaoXuan
 * @Email: huanghaoxuan1998@outlook.com
 * @github https://github.com/huanghaoxuan
 * @Date: 2019/6/23 18:29
 * @Version 1.0
 * 账号密码及权限
 */

type Student struct {
	Id          string `form:"id" gorm:"primary_key;column:id" json:"id"`                            //主键
	StudentId   string `form:"studentId" gorm:"column:student_id;not null;unique;" json:"studentId"` //学生学号、老师工号
	Password    string `form:"password" gorm:"column:password" json:"password"`                      //密码
	Permissions int    `form:"permissions" gorm:"column:permissions" json:"permissions"`             //权限控制
}

func (student *Student) SelectByStudentIdPassword() (int64, int) {
	first := db.DB.Where("student_id = ? and password = ?", student.StudentId, student.Password).First(&student)
	return first.RowsAffected, student.Permissions
}

func (student *Student) Insert() int64 {
	create := db.DB.Create(&student)
	return create.RowsAffected
}

func (student *Student) Update() int64 {
	updates := db.DB.Model(&student).Where("student_id = ?", student.StudentId).Updates(student)
	return updates.RowsAffected
}

//删除记录
//通过StudentId删除记录
func (student *Student) Delete() int64 {
	//防止记录被全部删除
	if student.StudentId != "" {
		i := db.DB.Where("student_id = ?", student.StudentId).Delete(&student)
		return i.RowsAffected
	}
	return 0
}
