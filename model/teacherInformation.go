package model

import (
	"score_inquiry_system/db"
	"time"
)

/**
 * @Author: HuangHaoXuan
 * @Email: huanghaoxuan1998@outlook.com
 * @github https://github.com/huanghaoxuan
 * @Date: 2019/7/2 9:21
 * @Version 1.0
 */

//储存学生相关信息
type TeacherInformation struct {
	Id          string    `form:"id" gorm:"primary_key;column:id" json:"id"`                            //主键
	TeacherId   string    `form:"teacherId" gorm:"column:teacher_id;not null;unique;" json:"teacherId"` //老师工号
	Name        string    `form:"name" gorm:"column:name" json:"name"`                                  //姓名
	Department  string    `form:"department" gorm:"column:department" json:"department"`                //所在学院或部门
	Permissions int       `form:"permissions" gorm:"column:permissions" json:"permissions"`             //权限控制
	CreatedAt   time.Time `form:"createdAt" gorm:"column:created_at" json:"createdAt"`                  //创建时间
}

//获得记录
//通过StudentId查询
func (information *TeacherInformation) SelectByStudentId() *TeacherInformation {
	db.DB.Where("teacher_id = ?", information.TeacherId).First(&information)
	return information
}

//通过id查询
func (information *TeacherInformation) SelectById() *TeacherInformation {
	db.DB.Where("id = ?", information.Id).First(&information)
	return information
}

//分页查询
func (information *TeacherInformation) SelectByPage(pageNum int, pageSize int) []TeacherInformation {
	teacherInformation := make([]TeacherInformation, 10)
	if pageNum > 0 && pageSize > 0 {
		db.DB.Where(&information).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&teacherInformation)
	}
	return teacherInformation
}

//查询总记录
func (information *TeacherInformation) Count() int {
	count := 0
	db.DB.Model(&information).Count(&count)
	return count
}

//插入记录
func (information *TeacherInformation) Insert() int64 {
	create := db.DB.Create(&information)
	return create.RowsAffected
}

//更新记录
//更新相关记录权限
func (information *TeacherInformation) Update() int64 {
	updates := db.DB.Model(&information).Where("id = ?", information.Id).Updates(information)
	return updates.RowsAffected
}

//删除记录
//通过id删除记录
func (information *TeacherInformation) Delete() int64 {
	//防止记录被全部删除
	if information.Id != "" {
		i := db.DB.Delete(&information)
		return i.RowsAffected
	}
	return 0
}
