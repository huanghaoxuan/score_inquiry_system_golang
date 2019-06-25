package model

import (
	"score_inquiry_system/db"
	"time"
)

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
	ID            string    `form:"id" gorm:"primary_key;column:id" json:"id"`                                                  //主键
	StudentID     string    `form:"student_id" gorm:"column:student_id;not null;unique;index:idx_student_id" json:"student_id"` //学生学号、老师工号
	Name          string    `form:"name" gorm:"column:name" json:"name"`                                                        //姓名
	DepartmentOld string    `form:"department_old" gorm:"column:department_old" json:"department_old"`                          //原所在学院或部门
	ClassOld      string    `form:"class_old" gorm:"column:class_old" json:"class_old"`                                         //原所在班级
	GradeOld      int       `form:"grade_old" gorm:"column:grade_old" json:"grade_old"`                                         //原所在年级
	DepartmentNew string    `form:"department_new" gorm:"column:department_new" json:"department_new"`                          //现所在学院或部门
	ClassNew      string    `form:"class_new" gorm:"column:class_new" json:"class_new"`                                         //现所在班级
	GradeNew      int       `form:"grade_new" gorm:"column:grade_new" json:"grade_new"`                                         //现所在年级
	Permissions   int       `form:"permissions" gorm:"column:permissions" json:"permissions"`                                   //权限控制
	CreatedAt     time.Time `form:"created_at" gorm:"column:created_at" json:"created_at"`                                      //创建时间
}

//获得记录
//通过id查询
func (information *StudentInformation) SelectByStudentID(StudentID string) *StudentInformation {
	var studentInformation StudentInformation
	db.DB.Where("student_id = ?", StudentID).First(&studentInformation)
	return &studentInformation
}

//插入记录
func (information *StudentInformation) Insert() int64 {
	create := db.DB.Create(information)
	return create.RowsAffected
}

//更新记录
//更新相关记录权限
func (information *StudentInformation) Update() int64 {
	updates := db.DB.Model(&information).Where("id = ?", information.ID).Updates(information)
	return updates.RowsAffected
}
