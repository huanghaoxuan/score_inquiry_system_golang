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

func GeneratedExcelCrossSemester(teachingClasses []model.TeachingClassResult) string {
	xlsx := excelize.NewFile()
	// 创建一个工作表
	index := xlsx.NewSheet("Sheet1")
	//// 设置单元格的值
	xlsx.SetCellValue("Sheet1", "A1", "姓名")
	xlsx.SetCellValue("Sheet1", "B1", "学号")
	xlsx.SetCellValue("Sheet1", "C1", "课程号")
	xlsx.SetCellValue("Sheet1", "D1", "课程名")
	xlsx.SetCellValue("Sheet1", "E1", "教学班号")
	//---------------------------------------------------------------------------------------------------------------------
	//组织数据 map的数组
	data := make([]map[string]interface{}, 0)
	head := make([]string, 0)
	for _, v := range teachingClasses {
		yearSem := strconv.Itoa(v.Year) + "-" + strconv.Itoa(v.Year+1) + " 学年" + "-" + v.Semester
		head = append(head, yearSem)
		results := v.SelectCrossSemester()
		for _, v := range results {
			if len(data) == 0 {
				data = append(data, make(map[string]interface{}))
				data[0][yearSem] = v.Result
				data[0]["studentId"] = v.StudentId
				data[0]["Name"] = v.Name
				data[0]["CourseId"] = v.CourseId
				data[0]["CourseName"] = v.CourseName
				data[0]["TeachingClassId"] = v.TeachingClassId
			}
			for i := 0; i < len(data); i++ {
				if data[i]["studentId"] == v.StudentId {
					data[i][yearSem] = v.Result
					break
				}
				if i == len(data)-1 {
					data = append(data, make(map[string]interface{}))
					data[i][yearSem] = v.Result
					data[i]["studentId"] = v.StudentId
					data[i]["Name"] = v.Name
					data[i]["CourseId"] = v.CourseId
					data[i]["CourseName"] = v.CourseName
					data[i]["TeachingClassId"] = v.TeachingClassId
					break
				}
			}
		}
	}

	//设置分布成绩表头
	for i, v := range head {
		xlsx.SetCellValue("Sheet1", getHeader(i+5, 1), v)
	}
	//--------------------------------------------------------------------------------------------------------------------
	//设置姓名、学号、课程名等值
	for i, v := range data {
		xlsx.SetCellValue("Sheet1", getHeader(0, i+2), v["Name"])
		xlsx.SetCellValue("Sheet1", getHeader(1, i+2), v["studentId"])
		xlsx.SetCellValue("Sheet1", getHeader(2, i+2), v["CourseId"])
		xlsx.SetCellValue("Sheet1", getHeader(3, i+2), v["CourseName"])
		xlsx.SetCellValue("Sheet1", getHeader(4, i+2), v["TeachingClassId"])
		for i1, v2 := range head {
			xlsx.SetCellValue("Sheet1", getHeader(i1+5, i+2), v[v2])
		}
	}
	//设置边框
	style, _ := xlsx.NewStyle("{'type':'1'}")
	xlsx.SetCellStyle("Sheet1", "A3", "D3", style)
	// 设置工作簿的默认工作表
	xlsx.SetActiveSheet(index)
	// 根据指定路径保存文件
	err := xlsx.SaveAs("public/finalScore/crossSemester.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	return "crossSemester.xlsx"
}

/*
	行列转化为“A1”的格式
*/
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
