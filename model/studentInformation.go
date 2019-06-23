package model

import "score_inquiry_system/db"

/**
 * @Author: HuangHaoXuan
 * @Email: huanghaoxuan1998@outlook.com
 * @github https://github.com/huanghaoxuan
 * @Date: 2019/6/23 20:08
 * @Version 1.0
 * 学生信息
 */

//储存学生相关信息
type StudentInformation struct {
	ID           string `form:"id" gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`       //主键
	StudentID    string `form:"student_id" gorm:"column:student_id" json:"student_id"`          //学生学号、老师工号
	Name         string `form:"name" gorm:"column:name" json:"name"`                            //姓名
	Department   string `form:"department" gorm:"column:department" json:"department"`          //所在学院或部门
	Class        string `form:"class" gorm:"column:class" json:"class"`                         //所在班级
	EnteringTime string `form:"entering_time" gorm:"column:entering_time" json:"entering_time"` //入校时间
	Permissions  int    `form:"permissions" gorm:"column:permissions" json:"permissions"`       //权限控制
}

//获得记录
//通过id查询
func SelectByStudentID(StudentID string) *StudentInformation {
	var studentInformation StudentInformation
	db.DB.Where("student_id = ?", StudentID).First(&studentInformation)
	return &studentInformation
}

//插入记录
func Insert(information *StudentInformation) {
	db.DB.Create(information)
}

//更新记录
//更新相关记录权限
func Update(information *StudentInformation) {
	db.DB.Model(&information).Updates(information)
}
