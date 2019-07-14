package teachingClassService

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	uuid "github.com/satori/go.uuid"
	"score_inquiry_system/model"
	"score_inquiry_system/service/teachingClassInformationService"
)

/**
 * @Author: HuangHaoXuan
 * @Email: huanghaoxuan1998@outlook.com
 * @github https://github.com/huanghaoxuan
 * @Date: 2019/6/30 15:12
 * @Version 1.0
 */

//插入
func Insert(teachingClass *model.TeachingClass) int64 {
	//设置uuid为主键
	teachingClass.Id = uuid.NewV4().String()
	return teachingClass.Insert()
}

//处理上传的学生信息表格
func ProcessingExcelFile(s string, courseId string, courseName string) {
	file, _ := excelize.OpenFile(s)
	rows := file.GetRows("Sheet1")
	for i, row := range rows {
		if i != 0 {
			teachingClass := model.TeachingClass{}
			teachingClassInformation := model.TeachingClassInformation{}
			for j, colCell := range row {
				switch j {
				case 0:
					teachingClass.Grade = colCell
				case 1:
					teachingClass.StudentId = colCell
				case 2:
					teachingClass.Name = colCell
				case 3:
					teachingClass.Department = colCell
				case 4:
					teachingClass.Professional = colCell
				case 5:
					teachingClass.Class = colCell
				case 6:
					teachingClass.CourseName = colCell
					teachingClassInformation.CourseName = colCell
				case 7:
					teachingClass.TeachingClassId = colCell
					teachingClassInformation.TeachingClassId = colCell
				case 8:
					teachingClass.CourseTeacherName = colCell
					teachingClassInformation.CourseTeacherName = colCell
				}
			}
			if courseName == teachingClass.CourseName {
				//插入基本信息
				teachingClass.CourseId = courseId
				teachingClassInformation.CourseId = courseId
				teachingClassInformationService.Insert(&teachingClassInformation)
				Insert(&teachingClass)
			}
		}

	}
}

//获取页数
func Count(teachingClass *model.TeachingClass) int {
	return teachingClass.Count()
}

//分页查询
func SelectByPage(pageNum int, pageSize int, teachingClass *model.TeachingClass) []model.TeachingClass {
	return teachingClass.SelectByPage(pageNum, pageSize)
}

//分页所有
func SelectAll(teachingClass *model.TeachingClass) []model.TeachingClass {
	return teachingClass.SelectAll()
}

//更新相关记录权限
func Update(teachingClass *model.TeachingClass) int64 {
	return teachingClass.Update()
}

//删除一条记录
func Delete(id string) int64 {
	teachingClass := model.TeachingClass{Id: id}
	return teachingClass.Delete()
}
