package teachingClassService

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"score_inquiry_system/model"
	"strconv"
)

/**
 * @Author: HuangHaoXuan
 * @Email: huanghaoxuan1998@outlook.com
 * @github https://github.com/huanghaoxuan
 * @Date: 2019/7/29 15:25
 * @Version 1.0
 */

func GeneratedExcel(teachingClassId string) string {
	xlsx := excelize.NewFile()
	// 创建一个工作表
	index := xlsx.NewSheet("Sheet1")
	//// 设置单元格的值
	xlsx.SetCellValue("Sheet1", "A1", "姓名")
	xlsx.SetCellValue("Sheet1", "B1", "学号")
	xlsx.SetCellValue("Sheet1", "C1", "课程名")
	xlsx.SetCellValue("Sheet1", "D1", "学年")
	xlsx.SetCellValue("Sheet1", "E1", "学期")
	//期末成绩、总成绩
	teachingClass := model.TeachingClass{TeachingClassId: teachingClassId}
	teachingClasses := teachingClass.SelectAll()
	//阶段成绩表头
	sourceStageInformation := model.SourceStageInformation{TeachingClassId: teachingClassId}
	sourceStageInformations := sourceStageInformation.SelectAll()
	//阶段成绩
	sourceStage := model.SourceStage{TeachingClassId: teachingClassId}
	sourceStages := sourceStage.SelectAll()
	//设置分布成绩表头
	for i, v := range sourceStageInformations {
		xlsx.SetCellValue("Sheet1", getHeader(i+5, 1), v.StageNote)
	}
	xlsx.SetCellValue("Sheet1", getHeader(len(sourceStageInformations)+5, 1), "期末成绩")
	xlsx.SetCellValue("Sheet1", getHeader(len(sourceStageInformations)+6, 1), "总成绩")
	//--------------------------------------------------------------------------------------------------------------------
	course := model.Course{}
	courses := course.SelectAll()
	//设置姓名、学号、课程名、学年、学期、期末成绩、总成绩
	for index := 0; index < len(teachingClasses); index++ {
		for _, v := range courses {
			if v.Id == teachingClasses[index].CourseId {
				course = v
				break
			}
		}
		//course := model.Course{Id: teachingClasses[index].CourseId}
		//course.SelectById()
		xlsx.SetCellValue("Sheet1", "A"+strconv.Itoa(index+2), teachingClasses[index].Name)
		xlsx.SetCellValue("Sheet1", "B"+strconv.Itoa(index+2), teachingClasses[index].StudentId)
		xlsx.SetCellValue("Sheet1", "C"+strconv.Itoa(index+2), course.Name)
		xlsx.SetCellValue("Sheet1", "D"+strconv.Itoa(index+2), strconv.Itoa(course.Year)+"-"+strconv.Itoa(course.Year+1))
		xlsx.SetCellValue("Sheet1", "E"+strconv.Itoa(index+2), course.Semester)
		for i, v := range sourceStageInformations {
			for _, v2 := range sourceStages {
				if v.Id == v2.SourceStageId && v2.StudentId == teachingClasses[index].StudentId {
					xlsx.SetCellValue("Sheet1", getHeader(i+5, index+2), v2.Scores)
				}
			}
		}

		xlsx.SetCellValue("Sheet1", getHeader(len(sourceStageInformations)+5, index+2), teachingClasses[index].Final)
		xlsx.SetCellValue("Sheet1", getHeader(len(sourceStageInformations)+6, index+2), teachingClasses[index].Result)
	}
	//设置边框
	style, _ := xlsx.NewStyle("{'type':'1'}")
	xlsx.SetCellStyle("Sheet1", "A3", "D3", style)
	// 设置工作簿的默认工作表
	xlsx.SetActiveSheet(index)
	// 根据指定路径保存文件
	err := xlsx.SaveAs("public/finalScore/" + teachingClassId + ".xlsx")
	if err != nil {
		fmt.Println(err)
	}
	return teachingClassId + ".xlsx"
}

func getHeader(row int, l int) string {
	axis := strconv.Itoa(l)
	for {
		axis = string(row%26+65) + axis
		if row/26 == 0 {
			break
		}
		row = row/26 - 1
	}
	return axis
}
