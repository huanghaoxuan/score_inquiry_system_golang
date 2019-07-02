package model

import (
	"score_inquiry_system/db"
	"time"
)

/**
 * @Author: HuangHaoXuan
 * @Email: huanghaoxuan1998@outlook.com
 * @github https://github.com/huanghaoxuan
 * @Date: 2019/6/24 23:10
 * @Version 1.0
 */

//教学班信息结构体
type TeachingClass struct {
	Id                string    `form:"id" gorm:"primary_key;column:id" json:"id"`                                         //主键
	Grade             string    `form:"grade" gorm:"column:grade" json:"grade"`                                            //所在年级
	StudentId         string    `form:"studentId" gorm:"column:student_id;not null;index:idx_student_id" json:"studentId"` //学生学号、老师工号
	Name              string    `form:"name" gorm:"column:name" json:"name"`                                               //姓名
	Department        string    `form:"department" gorm:"column:department" json:"department"`                             //所在学院或部门
	Professional      string    `form:"professional" gorm:"column:professional" json:"professional"`                       //所在专业
	Class             string    `form:"class" gorm:"column:class" json:"class"`                                            //所在班级
	CourseName        string    `form:"courseName" gorm:"column:course_name;index:idx_course_name" json:"courseName"`      //课程名称
	CourseId          string    `form:"courseId" gorm:"column:course_id" json:"courseId"`                                  //课程id
	CourseTeacherName string    `form:"courseTeacherName" gorm:"column:course_teacher_name" json:"courseTeacherName"`      //任课老师名字
	CourseTeacherId   string    `form:"courseTeacherId" gorm:"column:course_teacher_id" json:"courseTeacherId"`            //任课老师id
	CreatedAt         time.Time `form:"createdAt" gorm:"column:created_at" json:"createdAt"`                               //创建时间
}

//插入记录
func (teachingClass *TeachingClass) Insert() int64 {
	create := db.DB.Create(&teachingClass)
	return create.RowsAffected
}

//获得记录
//通过id查询
func (teachingClass *TeachingClass) SelectByStudentId(StudentId string) *TeachingClass {
	db.DB.Where("student_id = ?", StudentId).First(&teachingClass)
	return teachingClass
}

//分页查询
func (teachingClass *TeachingClass) SelectByPage(pageNum int, pageSize int) []TeachingClass {
	teachingClasses := make([]TeachingClass, 10)
	if pageNum > 0 && pageSize > 0 {
		db.DB.Where(&teachingClass).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&teachingClasses)
	}
	return teachingClasses
}

//查询总记录
func (teachingClass *TeachingClass) Count() int {
	count := 0
	db.DB.Model(&teachingClass).Count(&count)
	return count
}

//更新记录
func (teachingClass *TeachingClass) Update() int64 {
	updates := db.DB.Model(&teachingClass).Where("id = ?", teachingClass.Id).Updates(teachingClass)
	return updates.RowsAffected
}

//删除记录
func (teachingClass *TeachingClass) Delete() int64 {
	//防止记录被全部删除
	if teachingClass.Id != "" {
		i := db.DB.Delete(&teachingClass)
		return i.RowsAffected
	}
	return 0
}
