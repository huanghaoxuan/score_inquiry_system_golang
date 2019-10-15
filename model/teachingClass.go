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

//教学班学生信息结构体
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
	TeachingClassId   string    `form:"teachingClassId" gorm:"column:teaching_class_id" json:"teachingClassId"`            //教学班号
	CourseTeacherName string    `form:"courseTeacherName" gorm:"column:course_teacher_name" json:"courseTeacherName"`      //任课老师名字
	Final             string    `form:"final" gorm:"column:final" json:"final"`                                            //期末成绩
	Result            string    `form:"result" gorm:"column:result" json:"result"`                                         //最终成绩
	CourseTeacherId   string    `form:"courseTeacherId" gorm:"column:course_teacher_id" json:"courseTeacherId"`            //任课老师id
	CreatedAt         time.Time `form:"createdAt" gorm:"column:created_at" json:"createdAt"`                               //创建时间
}

//插入记录
func (teachingClass *TeachingClass) Insert() int64 {
	create := db.DB.Create(&teachingClass)
	return create.RowsAffected
}

//获得记录
//通过Id查询
func (teachingClass *TeachingClass) SelectById() *TeachingClass {
	db.DB.Where("id = ?", teachingClass.Id).First(&teachingClass)
	return teachingClass
}

func (teachingClass *TeachingClass) Select() []TeachingClass {
	teachingClasses := make([]TeachingClass, 10)
	db.DB.Where(&teachingClass).Order("created_at desc").Find(&teachingClasses)
	return teachingClasses
}

//分页查询
func (teachingClass *TeachingClass) SelectByPage(pageNum int, pageSize int) []TeachingClass {
	teachingClasses := make([]TeachingClass, 10)
	if pageNum > 0 && pageSize > 0 {
		db.DB.Where(&teachingClass).
			Order("created_at desc").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&teachingClasses)
	}
	return teachingClasses
}

//分页模糊查询
func (teachingClass *TeachingClass) SelectLikeByPage(pageNum int, pageSize int) []TeachingClass {
	teachingClasses := make([]TeachingClass, 10)
	if pageNum > 0 && pageSize > 0 {
		db.DB.
			Where("course_name LIKE ? AND student_id = ?", "%"+teachingClass.CourseName+"%", teachingClass.StudentId).
			Order("created_at desc").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&teachingClasses)
	}
	return teachingClasses
}

//查询所有
func (teachingClass *TeachingClass) SelectAll() []TeachingClass {
	teachingClasses := make([]TeachingClass, 10)
	db.DB.Where(&teachingClass).Order("created_at desc").Find(&teachingClasses)
	return teachingClasses
}

//查询总记录
func (teachingClass *TeachingClass) Count() int {
	count := 0
	db.DB.Model(&teachingClass).Where(&teachingClass).Count(&count)
	return count
}

//更新记录
func (teachingClass *TeachingClass) Update() int64 {
	if teachingClass.Id != "" {
		updates := db.DB.Model(&teachingClass).Where("id = ?", teachingClass.Id).Updates(teachingClass)
		return updates.RowsAffected
	}
	return 0
}

func (teachingClass *TeachingClass) UpdateAll() int64 {
	if teachingClass.Id != "" {
		updates := db.DB.Model(&teachingClass).Where("id = ?", teachingClass.Id).Save(teachingClass)
		return updates.RowsAffected
	}
	return 0
}

//删除记录
func (teachingClass *TeachingClass) Delete() int64 {
	//防止记录被全部删除
	if teachingClass.Id != "" {
		teachingClass.SelectById()
		db.DB.Where("teaching_class_id = ? and student_id = ?", teachingClass.TeachingClassId, teachingClass.StudentId).Delete(SourceStage{})
		i := db.DB.Delete(&teachingClass)
		return i.RowsAffected
	}
	return 0
}
