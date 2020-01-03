package model

import (
	"score_inquiry_system/db"
	"strconv"
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
	Id                string    `form:"id" gorm:"primary_key;column:id" json:"id"`                                        //主键
	Status            int       `form:"status" gorm:"column:status" json:"status"`                                        //成绩状态（1、教师录入中，2、教师已确认，3、成绩已发布）
	UniqueSign        string    `form:"uniqueSign" gorm:"column:unique_sign;unique;" json:"uniqueSign"`                   //唯一标志
	CourseName        string    `form:"courseName" gorm:"column:course_name;index:idx_course_name" json:"courseName"`     //课程名称
	CourseId          string    `form:"courseId" gorm:"column:course_id" json:"courseId"`                                 //课程id
	TeachingClassId   string    `form:"teachingClassId" gorm:"column:teaching_class_id;not null;" json:"teachingClassId"` //教学班号
	CourseTeacherName string    `form:"courseTeacherName" gorm:"column:course_teacher_name" json:"courseTeacherName"`     //任课老师名字
	CourseTeacherId   string    `form:"courseTeacherId" gorm:"column:course_teacher_id" json:"courseTeacherId"`           //任课老师id
	CreatedAt         time.Time `form:"createdAt" gorm:"column:created_at" json:"createdAt"`                              //创建时间
}

type TeachingClassInformationResult struct {
	TeachingClassInformation
	Year     int    `form:"year" json:"year"`
	Semester string `form:"semester" json:"semester"`
}

//插入记录
func (teachingClassInformation *TeachingClassInformation) Insert() int64 {
	create := db.DB.Create(&teachingClassInformation)
	return create.RowsAffected
}

//分页查询
func (teachingClass *TeachingClassInformationResult) SelectCrossSemester(pageNum int, pageSize int) ([]TeachingClassInformationResult, int) {
	result := make([]TeachingClassInformationResult, 10)
	count := 0
	sql := "t.course_id IN ( SELECT c.id FROM `course` c WHERE c.course_id = '" + teachingClass.CourseId + "' )AND t.teaching_class_id = '" +
		teachingClass.TeachingClassId + "'"
	if pageNum > 0 && pageSize > 0 {
		db.DB.
			Table(" teaching_class_information t ").
			Select("	t.`status`,t.`id`,t.`course_name`,t.`teaching_class_id`,t.`course_teacher_name`,t.`created_at`,t.`course_id`,c.`year`,c.`semester` ").
			Joins(" LEFT JOIN `course` c ON t.course_id = c.id ").
			Where(sql).
			Order(" created_at desc ").
			Limit(pageSize).Offset((pageNum - 1) * pageSize).
			Scan(&result)

		db.DB.
			Table("teaching_class_information t").
			Joins("LEFT JOIN `course` c ON t.course_id = c.id").
			Where(sql).
			Order("created_at desc").
			Limit(pageSize).Offset((pageNum - 1) * pageSize).
			Count(&count)
	}
	return result, count
}

//分页查询
func (teachingClassInformation *TeachingClassInformationResult) SelectByPage(pageNum int, pageSize int) []TeachingClassInformationResult {
	sql := "course_name LIKE ? AND teaching_class_id LIKE ? "
	if teachingClassInformation.CourseId != "" {
		sql = sql + " AND t.course_id = '" + teachingClassInformation.CourseId + "' "
	}
	if teachingClassInformation.CourseTeacherName != "" {
		sql = sql + " AND t.course_teacher_name = '" + teachingClassInformation.CourseTeacherName + "' "
	}
	if teachingClassInformation.Year != 0 {
		sql = sql + " AND c.year = '" + strconv.Itoa(teachingClassInformation.Year) + "' "
	}
	if teachingClassInformation.Semester != "" {
		sql = sql + " AND c.semester = '" + teachingClassInformation.Semester + "' "
	}
	result := make([]TeachingClassInformationResult, 10)
	//将course_id 设置为id返回
	if pageNum > 0 && pageSize > 0 {
		db.DB.
			Table("teaching_class_information t").
			Select(" t.`status`,t.`id` unique_sign,t.`course_id` id,t.`course_name`,t.`teaching_class_id`,t.`course_teacher_name`,t.`created_at`,c.`course_id`,c.`year`,c.`semester`").
			Joins("LEFT JOIN `course` c ON t.course_id = c.id").
			Where(sql,
				"%"+teachingClassInformation.CourseName+"%",
				"%"+teachingClassInformation.TeachingClassId+"%").
			Order("teaching_class_id ASC").
			Limit(pageSize).Offset((pageNum - 1) * pageSize).
			Scan(&result)

		//db.DB.
		//	Where("course_name LIKE ? AND teaching_class_id LIKE ? AND course_teacher_name = ?",
		//		"%"+teachingClassInformation.CourseName+"%",
		//		"%"+teachingClassInformation.TeachingClassId+"%",
		//		teachingClassInformation.CourseTeacherName).
		//	Order("created_at desc").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&teachingClassInformationes)
	}
	return result
}

//通过教学班号和课程号进行查询
func (teachingClassInformation *TeachingClassInformation) SelectByTeachingClassIdAndCourseId() *TeachingClassInformationResult {
	teachingClassInformationResult := TeachingClassInformationResult{}
	db.DB.Table("teaching_class_information t").
		Select(" t.*,c.`year`,c.`semester`").
		Joins("LEFT JOIN `course` c ON t.course_id = c.id").
		Where("t.teaching_class_id = ? AND t.course_id = ?", teachingClassInformation.TeachingClassId, teachingClassInformation.CourseId).Scan(&teachingClassInformationResult)
	return &teachingClassInformationResult
}

//通过id查询
func (teachingClassInformation *TeachingClassInformation) SelectById() *TeachingClassInformation {
	db.DB.Where("id = ?", teachingClassInformation.Id).First(&teachingClassInformation)
	return teachingClassInformation
}

func (teachingClassInformation *TeachingClassInformation) Select() []TeachingClassInformation {
	teachingClassInformationes := make([]TeachingClassInformation, 10)
	db.DB.Where(&teachingClassInformation).Find(&teachingClassInformationes)
	return teachingClassInformationes
}

//查询总记录
func (teachingClassInformation *TeachingClassInformation) Count() int {
	count := 0
	db.DB.Model(&teachingClassInformation).Where(&teachingClassInformation).Count(&count)
	return count
}

//查询已完成录入的个数
func (teachingClassInformation *TeachingClassInformation) CompleteInputCount() int {
	count := 0
	db.DB.Model(&teachingClassInformation).Where(&teachingClassInformation).Count(&count)
	return count
}

//更新记录
func (teachingClassInformation *TeachingClassInformation) Update() int64 {
	if teachingClassInformation.Id != "" {
		updates := db.DB.Model(&teachingClassInformation).Where("id = ?", teachingClassInformation.Id).Updates(teachingClassInformation)
		return updates.RowsAffected
	}
	return 0
}

func (teachingClassInformation *TeachingClassInformation) UpdateAll() int64 {
	if teachingClassInformation.Id != "" {
		updates := db.DB.Model(&teachingClassInformation).Where("id = ?", teachingClassInformation.Id).Save(teachingClassInformation)
		return updates.RowsAffected
	}
	return 0
}

//删除记录
func (teachingClassInformation *TeachingClassInformation) Delete() int64 {
	//防止记录被全部删除
	if teachingClassInformation.Id != "" {
		db.DB.Where("teaching_class_id = ? and course_id = ?", teachingClassInformation.TeachingClassId, teachingClassInformation.CourseId).Delete(TeachingClass{})
		db.DB.Where("teaching_class_id = ? and course_id = ?", teachingClassInformation.TeachingClassId, teachingClassInformation.CourseId).Delete(SourceStageInformation{})
		db.DB.Where("teaching_class_id = ? and course_id = ?", teachingClassInformation.TeachingClassId, teachingClassInformation.CourseId).Delete(SourceStage{})
		i := db.DB.Delete(&teachingClassInformation)
		return i.RowsAffected
	}
	return 0
}
