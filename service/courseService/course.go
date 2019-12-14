package courseService

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	uuid "github.com/satori/go.uuid"
	"score_inquiry_system/model"
	"strconv"
)

/**
 * @Author: HuangHaoXuan
 * @Email: huanghaoxuan1998@outlook.com
 * @github https://github.com/huanghaoxuan
 * @Date: 2019/7/4 19:14
 * @Version 1.0
 */
//获取页数
func Count(course *model.Course) int {
	return course.Count()
}

//分页查询
func SelectByPage(pageNum int, pageSize int, course *model.Course) []model.Course {
	return course.SelectByPage(pageNum, pageSize)
}

//表格处理
func ProcessingExcelFile(s string) {
	file, _ := excelize.OpenFile(s)
	rows := file.GetRows("Sheet1")
	for i, row := range rows {
		if i != 0 {
			course := model.Course{}
			for j, colCell := range row {
				switch j {
				case 0:
					course.CourseId = colCell
				case 1:
					course.Name = colCell
				case 2:
					course.Year, _ = strconv.Atoi(colCell)
				case 3:
					course.Semester = colCell
				}
			}
			//插入基本信息
			Insert(&course)
		}

	}
}

//插入
func Insert(course *model.Course) int64 {
	//设置uuid为主键
	course.Id = uuid.NewV4().String()
	return course.Insert()
}

//更新相关记录
func Update(course *model.Course) int64 {
	return course.Update()
}

func UpdateAll(course *model.Course) int64 {
	return course.UpdateAll()
}

//通过id查询
func SelectCourseById(id string) *model.Course {
	course := model.Course{Id: id}
	return course.SelectById()
}

//删除一条记录
func Delete(id string) int64 {
	//获取相关记录，获取学号
	course := model.Course{Id: id}
	return course.Delete()
}
