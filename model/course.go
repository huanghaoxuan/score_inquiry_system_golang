package model

import (
	"score_inquiry_system/db"
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
	Id        string    `form:"id" gorm:"primary_key;column:id" json:"id"`           //主键
	Name      string    `form:"name" gorm:"column:name;not null;" json:"name"`       //课程名字
	Year      int       `form:"year" gorm:"column:year" json:"year"`                 //学年
	Semester  string    `form:"semester" gorm:"column:semester" json:"semester"`     //学期
	CreatedAt time.Time `form:"createdAt" gorm:"column:created_at" json:"createdAt"` //创建时间

}

//通过id查询
func (course *Course) SelectById() *Course {
	db.DB.Where("id = ?", course.Id).First(&course)
	return course
}

//分页查询
func (course *Course) SelectByPage(pageNum int, pageSize int) []Course {
	courses := make([]Course, 10)
	if pageNum > 0 && pageSize > 0 {
		db.DB.Where(&course).Order("created_at desc").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&courses)
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
	updates := db.DB.Model(&course).Where("id = ?", course.Id).Updates(course)
	return updates.RowsAffected
}

//删除记录
//通过id删除记录
func (course *Course) Delete() int64 {
	//防止记录被全部删除
	if course.Id != "" {
		i := db.DB.Delete(&course)
		return i.RowsAffected
	}
	return 0
}
