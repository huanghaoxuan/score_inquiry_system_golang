package model

import "time"

/**
 * @Author: HuangHaoXuan
 * @Email: huanghaoxuan1998@outlook.com
 * @github https://github.com/huanghaoxuan
 * @Date: 2019/6/24 23:10
 * @Version 1.0
 */

//教学班信息结构体
type TeachingClass struct {
	ID                string    `form:"id" gorm:"primary_key;column:id" json:"id"`                                                //主键
	Grade             int       `form:"grade" gorm:"column:grade" json:"grade"`                                                   //所在年级
	StudentId         int       `form:"studentId" gorm:"column:student_id;not null;unique;index:idx_student_id" json:"studentId"` //学生学号、老师工号
	Name              string    `form:"name" gorm:"column:name" json:"name"`                                                      //姓名
	Department        string    `form:"department" gorm:"column:department" json:"department"`                                    //所在学院或部门
	Professional      string    `form:"professional" gorm:"column:professional" json:"professional"`                              //所在专业
	Class             string    `form:"class" gorm:"column:class" json:"class"`                                                   //所在班级
	CourseName        string    `form:"courseName" gorm:"column:course_name" json:"courseName"`                                   //课程名称
	Courseid          int       `form:"courseId" gorm:"column:course_id" json:"courseId"`                                         //任课老师名字
	CourseTeacherName string    `form:"courseTeacherName" gorm:"column:course_teacher_name" json:"courseTeacherName"`             //任课老师id
	CourseTeacherId   int       `form:"courseTeacherId" gorm:"column:course_teacher_id" json:"courseTeacherId"`                   //创建时间
	CreatedAt         time.Time `form:"createdAt" gorm:"column:created_at" json:"createdAt"`
}
