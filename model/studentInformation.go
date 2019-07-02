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
	Id            string    `form:"id" gorm:"primary_key;column:id" json:"id"`                            //主键
	StudentId     string    `form:"studentId" gorm:"column:student_id;not null;unique;" json:"studentId"` //学生学号、老师工号
	Name          string    `form:"name" gorm:"column:name" json:"name"`                                  //姓名
	DepartmentOld string    `form:"departmentOld" gorm:"column:department_old" json:"departmentOld"`      //原所在学院或部门
	ClassOld      string    `form:"classOld" gorm:"column:class_old" json:"classOld"`                     //原所在班级
	GradeOld      string    `form:"gradeOld" gorm:"column:grade_old" json:"gradeOld"`                     //原所在年级
	DepartmentNew string    `form:"departmentNew" gorm:"column:department_new" json:"departmentNew"`      //现所在学院或部门
	ClassNew      string    `form:"classNew" gorm:"column:class_new" json:"classNew"`                     //现所在班级
	GradeNew      string    `form:"gradeNew" gorm:"column:grade_new" json:"gradeNew"`                     //现所在年级
	Permissions   int       `form:"permissions" gorm:"column:permissions" json:"permissions"`             //权限控制
	CreatedAt     time.Time `form:"createdAt" gorm:"column:created_at" json:"createdAt"`                  //创建时间
}

//获得记录
//通过id查询
func (information *StudentInformation) SelectByStudentId(StudentId string) *StudentInformation {
	db.DB.Where("student_id = ?", StudentId).First(&information)
	return information
}

//分页查询
func (information *StudentInformation) SelectByPage(pageNum int, pageSize int) []StudentInformation {
	studentInformations := make([]StudentInformation, 10)
	if pageNum > 0 && pageSize > 0 {
		db.DB.Where(&information).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&studentInformations)
	}
	return studentInformations
}

//查询总记录
func (information *StudentInformation) Count() int {
	count := 0
	db.DB.Model(&information).Count(&count)
	return count
}

//插入记录
func (information *StudentInformation) Insert() int64 {
	create := db.DB.Create(&information)
	return create.RowsAffected
}

//更新记录
//更新相关记录权限
func (information *StudentInformation) Update() int64 {
	updates := db.DB.Model(&information).Where("id = ?", information.Id).Updates(information)
	return updates.RowsAffected
}

//删除记录
func (information *StudentInformation) Delete() int64 {
	//防止记录被全部删除
	if information.Id != "" {
		i := db.DB.Delete(&information)
		return i.RowsAffected
	}
	return 0
}
