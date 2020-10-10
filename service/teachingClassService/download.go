package teachingClassService

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	uuid "github.com/satori/go.uuid"
	"score_inquiry_system/model"
	"score_inquiry_system/util/common"
	"sort"
	"strconv"
)

/**
 * @Author: HuangHaoXuan
 * @Email: huanghaoxuan1998@outlook.com
 * @github https://github.com/huanghaoxuan
 * @Date: 2019/7/29 15:25
 * @Version 1.0
 */

func GeneratedExcel2(teachingClassId string, courseId string) string {
	CalculationResult(teachingClassId, courseId)
	//xlsx := excelize.NewFile()
	//// 创建一个工作表
	//index := xlsx.NewSheet("Sheet1")
	//// 设置工作簿的默认工作表
	//xlsx.SetActiveSheet(index)
	xlsx, _ := excelize.OpenFile("./default.xlsx")
	alignmentStyle := xlsx.GetCellStyle("Sheet1", "A1")
	borderStyle := xlsx.GetCellStyle("Sheet1", "B1")
	//阶段成绩表头信息
	sourceStageInformation := model.SourceStageInformation{TeachingClassId: teachingClassId, CourseId: courseId}
	sourceStageInformations := sourceStageInformation.SelectAll()

	xlsx.MergeCell("Sheet1", "A1", getHeader(len(sourceStageInformations)+5, 1))
	xlsx.SetCellValue("Sheet1", "A1", "东南大学成贤学院过程成绩记录表")
	//查询课程信息
	teachingClassInformat := model.TeachingClassInformation{TeachingClassId: teachingClassId, CourseId: courseId}
	teachingClassInformationResult := teachingClassInformat.SelectByTeachingClassIdAndCourseId()
	teacher := model.TeacherInformation{Name: teachingClassInformationResult.CourseTeacherName}
	//期末成绩、总成绩
	teachingClass := model.TeachingClass{TeachingClassId: teachingClassId, CourseId: courseId}
	teachingClasses := teachingClass.SelectDownload()
	teacher.SelectByName()
	xlsx.MergeCell("Sheet1", "A2", getHeader(len(sourceStageInformations)+5, 2))
	xlsx.SetCellValue("Sheet1", "A2", strconv.Itoa(teachingClassInformationResult.Year)+"-"+strconv.Itoa(teachingClassInformationResult.Year+1)+" "+teachingClassInformationResult.Semester)
	xlsx.MergeCell("Sheet1", "A3", getHeader((len(sourceStageInformations)+5)/3-1, 3))

	xlsx.SetRowHeight("Sheet1", 3, 40)
	xlsx.SetCellValue("Sheet1", "A3", "承担单位："+teacher.Department)

	xlsx.MergeCell("Sheet1", getHeader((len(sourceStageInformations)+5)/3, 3), getHeader((len(sourceStageInformations)+5)/3*2-1, 3))
	xlsx.SetCellValue("Sheet1", getHeader((len(sourceStageInformations)+5)/3, 3), "课程："+teachingClassInformationResult.CourseName)
	xlsx.MergeCell("Sheet1", getHeader((len(sourceStageInformations)+5)/3*2, 3), getHeader(len(sourceStageInformations)+5, 3))
	xlsx.SetCellValue("Sheet1", getHeader((len(sourceStageInformations)+5)/3*2, 3), "人数："+strconv.Itoa(len(teachingClasses)))

	xlsx.MergeCell("Sheet1", "A4", getHeader((len(sourceStageInformations)+5)/3-1, 4))
	xlsx.SetCellValue("Sheet1", "A4", "任课教师："+teacher.Name)
	classes := make([]string, len(teachingClasses))
	for _, v := range teachingClasses {
		classes = append(classes, v.Class)
	}
	sort.Strings(classes)
	classes = common.RemoveRepByLoop(classes)
	class := "上课班级："
	for i, v := range classes {
		if v != "" {
			class += v
			if i+1 != len(classes) {
				class += "，"
			}
		}

	}

	xlsx.MergeCell("Sheet1", getHeader((len(sourceStageInformations)+5)/3, 4), getHeader((len(sourceStageInformations)+5)/3*2-1, 4))
	xlsx.SetRowHeight("Sheet1", 4, float64(len([]rune(class))/12+1)*15)
	xlsx.SetCellValue("Sheet1", getHeader((len(sourceStageInformations)+5)/3, 4), class)
	xlsx.MergeCell("Sheet1", getHeader((len(sourceStageInformations)+5)/3*2, 4), getHeader(len(sourceStageInformations)+5, 4))
	xlsx.SetCellValue("Sheet1", getHeader((len(sourceStageInformations)+5)/3*2, 4), "课序号："+teachingClassId)
	xlsx.MergeCell("Sheet1", "A5", getHeader(len(sourceStageInformations)+5, 5))
	//设置成绩的组成公式
	composition := "总评成绩 = "
	finalPercentage := 100.0
	for _, v := range sourceStageInformations {
		composition += v.StageNote + " * " + v.Percentage + " % + "
		percentage, _ := strconv.ParseFloat(v.Percentage, 64)
		finalPercentage = finalPercentage - percentage
	}
	composition += "期末成绩 * " + fmt.Sprintf("%.1f", finalPercentage) + " % "
	xlsx.SetRowHeight("Sheet1", 5, float64(60))
	xlsx.SetCellValue("Sheet1", "A5", composition)
	/*表格文件总览数据填写完成*/

	//设置成绩的表头
	xlsx.MergeCell("Sheet1", "A6", "A7")
	xlsx.SetCellValue("Sheet1", "A6", "序号")
	xlsx.MergeCell("Sheet1", "B6", "B7")
	xlsx.SetColWidth("Sheet1", "B", "B", 20)
	xlsx.SetCellValue("Sheet1", "B6", "班级")
	xlsx.MergeCell("Sheet1", "C6", "C7")
	xlsx.SetCellValue("Sheet1", "C6", "学号")
	xlsx.MergeCell("Sheet1", "D6", "D7")
	xlsx.SetCellValue("Sheet1", "D6", "姓名")
	xlsx.MergeCell("Sheet1", "E6", getHeader(len(sourceStageInformations)+4, 6))
	xlsx.SetCellValue("Sheet1", "E6", "考核项目")

	//设置分布成绩表头
	//阶段成绩
	sourceStage := model.SourceStage{TeachingClassId: teachingClassId, CourseId: courseId}
	sourceStages := sourceStage.SelectAll()
	for i, v := range sourceStageInformations {
		percentage, _ := strconv.ParseFloat(v.Percentage, 64)
		xlsx.SetCellValue("Sheet1", getHeader(i+4, 7), v.StageNote+fmt.Sprintf("%.1f", percentage)+" % ")
	}
	xlsx.SetCellValue("Sheet1", getHeader(len(sourceStageInformations)+4, 7), "期末成绩")
	xlsx.MergeCell("Sheet1", getHeader(len(sourceStageInformations)+5, 6), getHeader(len(sourceStageInformations)+5, 7))
	xlsx.SetCellValue("Sheet1", getHeader(len(sourceStageInformations)+5, 6), "总评成绩")

	//设置姓名、学号、课程名、学年、学期、期末成绩、总成绩
	for index := 0; index < len(teachingClasses); index++ {
		xlsx.SetCellValue("Sheet1", "A"+strconv.Itoa(index+8), index+1)
		xlsx.SetCellValue("Sheet1", "B"+strconv.Itoa(index+8), teachingClasses[index].Class)
		xlsx.SetCellValue("Sheet1", "C"+strconv.Itoa(index+8), teachingClasses[index].StudentId)
		xlsx.SetCellValue("Sheet1", "D"+strconv.Itoa(index+8), teachingClasses[index].Name)

		for i, v := range sourceStageInformations {
			for _, v2 := range sourceStages {
				if v.Id == v2.SourceStageId && v2.StudentId == teachingClasses[index].StudentId {
					xlsx.SetCellValue("Sheet1", getHeader(i+4, index+8), v2.Scores)
				}
			}
		}

		xlsx.SetCellValue("Sheet1", getHeader(len(sourceStageInformations)+4, index+8), teachingClasses[index].Final)
		xlsx.SetCellValue("Sheet1", getHeader(len(sourceStageInformations)+5, index+8), teachingClasses[index].Result)
	}

	//设置表格样式
	//alignmentStyle, _ := xlsx.NewStyle(`{"alignment":{"horizontal":"center","vertical":"center"}}`)
	xlsx.SetCellStyle("Sheet1", "A1", getHeader(len(sourceStageInformations)+5, 5), alignmentStyle)

	//style, _ := xlsx.NewStyle(`{"border":[{"type":"left","color":"000000","style":1},{"type":"right","color":"000000","style":1},{"type":"top","color":"000000","style":1},{"type":"bottom","color":"000000","style":1}]}`)
	xlsx.SetCellStyle("Sheet1", "A6", getHeader(len(sourceStageInformations)+5, len(teachingClasses)+7), borderStyle)

	xlsx.ProtectSheet("Sheet1", &excelize.FormatSheetProtection{
		Password:      "admin",
		EditScenarios: false,
	})

	// 根据指定路径保存文件
	if err := xlsx.SaveAs("public/finalScore/" + courseId + "-" + teachingClassId + ".xlsx"); err != nil {
		println(err.Error())
	}
	return courseId + "-" + teachingClassId + ".xlsx"
}

func GeneratedExcel(teachingClassId string, courseId string) string {
	xlsx := excelize.NewFile()
	// 创建一个工作表
	index := xlsx.NewSheet("Sheet1")
	//// 设置单元格的值
	xlsx.SetCellValue("Sheet1", "A1", "任课教师")
	xlsx.SetCellValue("Sheet1", "B1", "课程编号")
	xlsx.SetCellValue("Sheet1", "C1", "课程名称")
	xlsx.SetCellValue("Sheet1", "D1", "学号")
	xlsx.SetCellValue("Sheet1", "E1", "姓名")
	xlsx.SetCellValue("Sheet1", "F1", "所在班级")
	xlsx.SetCellValue("Sheet1", "G1", "教学班号")
	xlsx.SetCellValue("Sheet1", "H1", "学年")
	xlsx.SetCellValue("Sheet1", "I1", "学期")
	//期末成绩、总成绩
	teachingClass := model.TeachingClass{TeachingClassId: teachingClassId, CourseId: courseId}
	teachingClasses := teachingClass.SelectDownload()
	//阶段成绩表头
	sourceStageInformation := model.SourceStageInformation{TeachingClassId: teachingClassId, CourseId: courseId}
	sourceStageInformations := sourceStageInformation.SelectAll()
	//阶段成绩
	sourceStage := model.SourceStage{TeachingClassId: teachingClassId, CourseId: courseId}
	sourceStages := sourceStage.SelectAll()
	//设置分布成绩表头
	for i, v := range sourceStageInformations {
		xlsx.SetCellValue("Sheet1", getHeader(i+9, 1), v.StageNote)
	}
	xlsx.SetCellValue("Sheet1", getHeader(len(sourceStageInformations)+9, 1), "期末成绩")
	xlsx.SetCellValue("Sheet1", getHeader(len(sourceStageInformations)+10, 1), "总成绩")
	//--------------------------------------------------------------------------------------------------------------------

	//设置姓名、学号、课程名、学年、学期、期末成绩、总成绩
	for index := 0; index < len(teachingClasses); index++ {
		//course := model.Course{Id: teachingClasses[index].CourseId}
		//course.SelectById()
		xlsx.SetCellValue("Sheet1", "A"+strconv.Itoa(index+2), teachingClasses[index].CourseTeacherName)
		xlsx.SetCellValue("Sheet1", "B"+strconv.Itoa(index+2), teachingClasses[index].CourseId)
		xlsx.SetCellValue("Sheet1", "C"+strconv.Itoa(index+2), teachingClasses[index].CourseName)
		xlsx.SetCellValue("Sheet1", "D"+strconv.Itoa(index+2), teachingClasses[index].StudentId)
		xlsx.SetCellValue("Sheet1", "E"+strconv.Itoa(index+2), teachingClasses[index].Name)
		xlsx.SetCellValue("Sheet1", "F"+strconv.Itoa(index+2), teachingClasses[index].Class)
		xlsx.SetCellValue("Sheet1", "G"+strconv.Itoa(index+2), teachingClasses[index].TeachingClassId)
		xlsx.SetCellValue("Sheet1", "H"+strconv.Itoa(index+2), strconv.Itoa(teachingClasses[index].Year)+"-"+strconv.Itoa(teachingClasses[index].Year+1))
		xlsx.SetCellValue("Sheet1", "I"+strconv.Itoa(index+2), teachingClasses[index].Semester)

		for i, v := range sourceStageInformations {
			for _, v2 := range sourceStages {
				if v.Id == v2.SourceStageId && v2.StudentId == teachingClasses[index].StudentId {
					xlsx.SetCellValue("Sheet1", getHeader(i+9, index+2), v2.Scores)
				}
			}
		}

		xlsx.SetCellValue("Sheet1", getHeader(len(sourceStageInformations)+9, index+2), teachingClasses[index].Final)
		xlsx.SetCellValue("Sheet1", getHeader(len(sourceStageInformations)+10, index+2), teachingClasses[index].Result)
	}
	//设置边框
	style, _ := xlsx.NewStyle("{'type':'1'}")
	xlsx.SetCellStyle("Sheet1", "A3", "D3", style)
	// 设置工作簿的默认工作表
	xlsx.SetActiveSheet(index)
	// 根据指定路径保存文件
	err := xlsx.SaveAs("public/finalScore/" + courseId + "-" + teachingClassId + ".xlsx")
	if err != nil {
		fmt.Println(err)
	}
	return courseId + "-" + teachingClassId + ".xlsx"
}

func GeneratedExcelCrossSemester(teachingClasses []model.TeachingClassResult) string {
	xlsx, _ := excelize.OpenFile("./default.xlsx")
	alignmentStyle := xlsx.GetCellStyle("Sheet1", "A1")
	borderStyle := xlsx.GetCellStyle("Sheet1", "B1")
	//xlsx := excelize.NewFile()
	// 创建一个工作表
	//index := xlsx.NewSheet("Sheet1")
	//// 设置单元格的值
	xlsx.SetCellValue("Sheet1", "A1", "任课教师")
	xlsx.SetCellValue("Sheet1", "B1", "课程号")
	xlsx.SetCellValue("Sheet1", "C1", "课程名")
	xlsx.SetCellValue("Sheet1", "D1", "教学班号")
	xlsx.SetCellValue("Sheet1", "E1", "姓名")
	xlsx.SetCellValue("Sheet1", "F1", "学号")

	//---------------------------------------------------------------------------------------------------------------------
	//组织数据 map的数组
	data := make([]map[string]interface{}, 0)
	head := make([]string, 0)
	for index, v := range teachingClasses {
		yearSem := strconv.Itoa(v.Year) + "-" + strconv.Itoa(v.Year+1) + " 学年" + "-" + v.Semester
		head = append(head, yearSem)
		results := v.SelectCrossSemester()
		for i, v := range results {
			if index == 0 {
				data = append(data, make(map[string]interface{}))
				data[i][yearSem] = v.Result
				data[i]["studentId"] = v.StudentId
				data[i]["Name"] = v.Name
				data[i]["CourseId"] = v.CourseId
				data[i]["CourseName"] = v.CourseName
				data[i]["TeachingClassId"] = v.TeachingClassId
				data[i]["courseTeacherName"] = v.CourseTeacherName
			}
			for i := 0; i < len(data); i++ {
				if data[i]["studentId"] == v.StudentId {
					data[i][yearSem] = v.Result
					break
				}
			}
		}
	}

	//设置分布成绩表头
	for i, v := range head {
		xlsx.SetCellValue("Sheet1", getHeader(i+6, 1), v)
	}
	xlsx.SetCellValue("Sheet1", getHeader(len(head)+6, 1), "总评成绩")
	xlsx.SetCellStyle("Sheet1", "A1", getHeader(len(head)+6, 1), alignmentStyle)
	//--------------------------------------------------------------------------------------------------------------------
	//设置姓名、学号、课程名等值

	for i, v := range data {
		xlsx.SetCellValue("Sheet1", getHeader(0, i+2), v["courseTeacherName"])
		xlsx.SetCellValue("Sheet1", getHeader(1, i+2), v["CourseId"])
		xlsx.SetCellValue("Sheet1", getHeader(2, i+2), v["CourseName"])
		xlsx.SetCellValue("Sheet1", getHeader(3, i+2), v["TeachingClassId"])
		xlsx.SetCellValue("Sheet1", getHeader(4, i+2), v["Name"])
		xlsx.SetCellValue("Sheet1", getHeader(5, i+2), v["studentId"])
		var avg int = 0
		for i1, v2 := range head {
			xlsx.SetCellValue("Sheet1", getHeader(i1+6, i+2), v[v2])
			b, _ := (v[v2]).(string)
			c, _ := strconv.Atoi(b)
			avg = avg + c
		}
		xlsx.SetCellValue("Sheet1", getHeader(len(head)+6, i+2), avg/len(head))
	}
	//设置边框
	xlsx.SetCellStyle("Sheet1", "A2", getHeader(len(head)+6, len(data)+1), borderStyle)
	// 设置工作簿的默认工作表
	//xlsx.SetActiveSheet(index)
	// 根据指定路径保存文件
	fileName := "crossSemester-" + uuid.NewV4().String() + ".xlsx"
	err := xlsx.SaveAs("public/finalScore/" + fileName)
	if err != nil {
		fmt.Println(err)
	}
	return fileName
}

/*
	（0，1）行列转化为“A1”的格式
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

func getRow(row int) string {
	axis := ""
	for {
		axis = string(row%26+65) + axis
		if row/26 == 0 {
			break
		}
		row = row/26 - 1
	}
	return axis
}

func subString(str string, length int) (string, int) {
	count := 1
	result := ""
	for count*length < len([]rune(str)) {
		result += string([]rune(str)[(count-1)*length:count*length]) + "\r\n"
		count++
	}
	result += string([]rune(str)[(count-1)*length:])
	return result, count
}
