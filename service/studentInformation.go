package service

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	uuid "github.com/satori/go.uuid"
	"score_inquiry_system/model"
	"strconv"
)

/**
 * @Author: HuangHaoXuan
 * @Email: huanghaoxuan1998@outlook.com
 * @github https://github.com/huanghaoxuan
 * @Date: 2019/6/23 20:32
 * @Version 1.0
 */

//插入
func Insert(information *model.StudentInformation) int64 {
	if information.SelectByStudentId(information.StudentId).Id != "" {
		return 0
	}
	//设置uuid为主键
	information.Id = uuid.NewV4().String()
	//默认权限为1
	information.Permissions = 1
	return information.Insert()
}

//更新相关记录权限
func Update(information *model.StudentInformation) int64 {
	informationOld := information.SelectByStudentId(information.StudentId)
	fmt.Println(informationOld)
	information.Id = informationOld.Id
	fmt.Println(information)
	return information.Update()
}

//处理上传的学生信息表格
func ProcessingExcelFile(s string) {
	file, _ := excelize.OpenFile(s)
	rows := file.GetRows("Sheet1")
	for i, row := range rows {
		if i != 0 {
			studentInformation := model.StudentInformation{}
			studentInformation.Id = uuid.NewV4().String()
			for j, colCell := range row {
				switch j {
				case 0:
					studentInformation.StudentId = colCell
				case 1:
					studentInformation.Name = colCell
				case 2:
					{
						gradeOld, _ := strconv.Atoi(colCell)
						studentInformation.GradeOld = gradeOld
					}
				case 3:
					studentInformation.DepartmentOld = colCell
				case 4:
					studentInformation.ClassOld = colCell
				case 5:
					{
						gradeNew, _ := strconv.Atoi(colCell)
						studentInformation.GradeNew = gradeNew
					}
				case 6:
					studentInformation.DepartmentNew = colCell
				case 7:
					studentInformation.ClassNew = colCell
				}
			}
			fmt.Println(studentInformation)
		}

	}
}
