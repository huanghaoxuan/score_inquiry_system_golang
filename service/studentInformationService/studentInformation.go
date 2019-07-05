package studentInformationService

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	uuid "github.com/satori/go.uuid"
	"score_inquiry_system/model"
	"score_inquiry_system/service/loginService"
)

/**
 * @Author: HuangHaoXuan
 * @Email: huanghaoxuan1998@outlook.com
 * @github https://github.com/huanghaoxuan
 * @Date: 2019/6/23 20:32
 * @Version 1.0
 */

//获取页数
func Count(information *model.StudentInformation) int {
	return information.Count()
}

//分页查询
func SelectByPage(pageNum int, pageSize int, information *model.StudentInformation) []model.StudentInformation {
	return information.SelectByPage(pageNum, pageSize)
}

//插入
func Insert(information *model.StudentInformation) int64 {
	//设置uuid为主键
	information.Id = uuid.NewV4().String()
	//默认权限为1
	information.Permissions = 1
	//插入登录信息
	loginService.Registered(information.StudentId, information.StudentId, 1)
	return information.Insert()
}

//更新相关记录
func Update(information *model.StudentInformation) int64 {
	return information.Update()
}

//处理上传的学生信息表格
func ProcessingExcelFile(s string) {
	file, _ := excelize.OpenFile(s)
	rows := file.GetRows("Sheet1")
	for i, row := range rows {
		if i != 0 {
			studentInformation := model.StudentInformation{}
			for j, colCell := range row {
				switch j {
				case 0:
					studentInformation.StudentId = colCell
				case 1:
					studentInformation.Name = colCell
				case 2:
					studentInformation.GradeOld = colCell
				case 3:
					studentInformation.DepartmentOld = colCell
				case 4:
					studentInformation.ClassOld = colCell
				case 5:
					studentInformation.GradeNew = colCell
				case 6:
					studentInformation.DepartmentNew = colCell
				case 7:
					studentInformation.ClassNew = colCell
				}
			}
			//插入基本信息
			Insert(&studentInformation)
		}

	}
}

//通过id查询
func SelectStudentInformationById(id string) *model.StudentInformation {
	information := model.StudentInformation{Id: id}
	return information.SelectById()
}

//删除一条记录
func Delete(id string) int64 {
	//获取相关记录，获取学号
	information := SelectStudentInformationById(id)
	loginService.Delete(information.StudentId)
	return information.Delete()
}
