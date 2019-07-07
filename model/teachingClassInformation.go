package model

import (
	"score_inquiry_system/db"
	"time"
)

/**
 * @Author: HuangHaoXuan
 * @Email: huanghaoxuan1998@outlook.com
 * @github https://github.com/huanghaoxuan
 * @Date: 2019/7/6 20:08
 * @Version 1.0
 */

//教学班信息结构体
type TeachingClassInformation struct {
	Id                string    `form:"id" gorm:"primary_key;column:id" json:"id"`                                               //主键
	CourseName        string    `form:"courseName" gorm:"column:course_name;index:idx_course_name" json:"courseName"`            //课程名称
	CourseId          string    `form:"courseId" gorm:"column:course_id" json:"courseId"`                                        //课程id
	TeachingClassId   string    `form:"teachingClassId" gorm:"column:teaching_class_id;not null;unique;" json:"teachingClassId"` //教学班号
	CourseTeacherName string    `form:"courseTeacherName" gorm:"column:course_teacher_name" json:"courseTeacherName"`            //任课老师名字
	CourseTeacherId   string    `form:"courseTeacherId" gorm:"column:course_teacher_id" json:"courseTeacherId"`                  //任课老师id
	CreatedAt         time.Time `form:"createdAt" gorm:"column:created_at" json:"createdAt"`                                     //创建时间
}

//插入记录
func (teachingClassInformation *TeachingClassInformation) Insert() int64 {
	create := db.DB.Create(&teachingClassInformation)
	return create.RowsAffected
}

//分页查询
func (teachingClassInformation *TeachingClassInformation) SelectByPage(pageNum int, pageSize int) []TeachingClassInformation {
	teachingClassInformationes := make([]TeachingClassInformation, 10)
	if pageNum > 0 && pageSize > 0 {
		db.DB.Where(&teachingClassInformation).Order("created_at desc").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&teachingClassInformationes)
	}
	return teachingClassInformationes
}

//通过id查询
func (teachingClassInformation *TeachingClassInformation) SelectById() *TeachingClassInformation {
	db.DB.Where("id = ?", teachingClassInformation.Id).First(&teachingClassInformation)
	return teachingClassInformation
}

//查询总记录
func (teachingClassInformation *TeachingClassInformation) Count() int {
	count := 0
	db.DB.Model(&teachingClassInformation).Where(&teachingClassInformation).Count(&count)
	return count
}

//更新记录
func (teachingClassInformation *TeachingClassInformation) Update() int64 {
	updates := db.DB.Model(&teachingClassInformation).Where("id = ?", teachingClassInformation.Id).Updates(teachingClassInformation)
	return updates.RowsAffected
}

//删除记录
func (teachingClassInformation *TeachingClassInformation) Delete() int64 {
	//防止记录被全部删除
	if teachingClassInformation.Id != "" {
		i := db.DB.Delete(&teachingClassInformation)
		return i.RowsAffected
	}
	return 0
}
