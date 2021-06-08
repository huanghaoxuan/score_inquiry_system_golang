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
 * @Date: 2019/7/3 15:46
 * @Version 1.0
 */

//课程信息结构体
type Course struct {
	Id                  string    `form:"id" gorm:"primary_key;column:id" json:"id"`                                          //主键
	Status              int       `form:"status" gorm:"column:status" json:"status"`                                          //成绩状态（1、教师录入中，2、教师已确认，3、成绩已发布）
	Name                string    `form:"name" gorm:"column:name;not null;" json:"name"`                                      //课程名字
	CourseId            string    `form:"courseId" gorm:"column:course_id;not null;" json:"courseId"`                         //课程号
	CourseAdministrator string    `form:"courseAdministrator" gorm:"column:course_administrator;" json:"courseAdministrator"` //课程管理员
	Year                int       `form:"year" gorm:"column:year" json:"year"`                                                //学年
	Semester            string    `form:"semester" gorm:"column:semester" json:"semester"`                                    //学期
	CreatedAt           time.Time `form:"createdAt" gorm:"column:created_at" json:"createdAt"`                                //创建时间

}

//通过id查询
func (course *Course) SelectById() *Course {
	db.DB.Where("id = ?", course.Id).First(&course)
	return course
}

//查询
func (course *Course) SelectAll() []Course {
	courses := make([]Course, 10)
	db.DB.Where(&course).Order("created_at desc").Find(&courses)
	return courses
}

//分页查询
func (course *Course) SelectByPage(pageNum int, pageSize int) []Course {
	courses := make([]Course, 10)
	extSql := ""
	if course.Year != 0 {
		extSql += " AND year = " + strconv.Itoa(course.Year) + " "
	}
	if course.Semester != "" {
		extSql += " AND semester = '" + course.Semester + "' "
	}
	if pageNum > 0 && pageSize > 0 {
		if course.CourseAdministrator != "" {
			db.DB.Where("name LIKE ? AND course_id LIKE ? AND course_administrator = ?"+extSql,
				"%"+course.Name+"%",
				"%"+course.CourseId+"%",
				course.CourseAdministrator).
				Order("created_at desc").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&courses)
		} else {
			db.DB.Where("name LIKE ? AND course_id LIKE ? "+extSql, "%"+course.Name+"%", "%"+course.CourseId+"%").
				Order("created_at desc").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&courses)
		}
	}
	return courses
}

//查询总记录
func (course *Course) Count() int {
	count := 0
	db.DB.Model(&course).Where(&course).Count(&count)
	return count
}

//插入记录
func (course *Course) Insert() int64 {
	create := db.DB.Create(&course)
	return create.RowsAffected
}

//更新记录
//更新相关记录权限
func (course *Course) Update() int64 {
	if course.Id != "" {
		updates := db.DB.Model(&course).Where("id = ?", course.Id).Updates(course)
		return updates.RowsAffected
	}
	return 0
}

func (course *Course) UpdateAll() int64 {
	if course.Id != "" {
		updates := db.DB.Model(&course).Where("id = ?", course.Id).Save(course)
		return updates.RowsAffected
	}
	return 0
}

//删除记录
//通过id删除记录
func (course *Course) Delete() int64 {

	//防止记录被全部删除
	if course.Id != "" {
		course.SelectById()
		teachingClassInformation := TeachingClassInformation{CourseId: course.Id}
		teachingClassInformations := teachingClassInformation.Select()
		for _, v := range teachingClassInformations {
			db.DB.Where("teaching_class_id = ? and course_id = ?", v.TeachingClassId, course.Id).Delete(SourceStageInformation{})
			db.DB.Where("teaching_class_id = ? and course_id = ?", v.TeachingClassId, course.Id).Delete(SourceStage{})
		}
		db.DB.Where("course_id = ?", course.Id).Delete(TeachingClassInformation{})
		db.DB.Where("course_id = ?", course.Id).Delete(TeachingClass{})
		i := db.DB.Delete(&course)
		return i.RowsAffected
	}
	return 0
}
