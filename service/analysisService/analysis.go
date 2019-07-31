package analysisService

import (
	"score_inquiry_system/model"
	"strconv"
)

/**
 * @Author: HuangHaoXuan
 * @Email: huanghaoxuan1998@outlook.com
 * @github https://github.com/huanghaoxuan
 * @Date: 2019/7/30 14:11
 * @Version 1.0
 */

//分析成绩通过情况
func Pass(studentId string) interface{} {
	teachingClass := model.TeachingClass{StudentId: studentId}
	teachingClasses := teachingClass.SelectAll()
	a := 0 //满分
	b := 0 //优秀
	c := 0 //及格
	d := 0 //不及格
	for _, v := range teachingClasses {
		if v.Result == "" {
			break
		}
		result, _ := strconv.ParseFloat(v.Result, 64)
		switch {
		case result == 100:
			a++
		case result >= 85:
			b++
		case result >= 60:
			c++
		case result < 60:
			d++
		}
	}
	data := make([]map[string]interface{}, 4)
	data[0] = map[string]interface{}{"item": "满分", "count": a}
	data[1] = map[string]interface{}{"item": "优秀", "count": b}
	data[2] = map[string]interface{}{"item": "及格", "count": c}
	data[3] = map[string]interface{}{"item": "不及格", "count": d}
	return data
}

//分析每年课程情况
func ScoreCount(studentId string) interface{} {
	teachingClass := model.TeachingClass{StudentId: studentId}
	teachingClasses := teachingClass.SelectAll()
	data := make([]map[string]int, 0, len(teachingClasses))
	for index := 0; index < len(teachingClasses); index++ {
		course := model.Course{Id: teachingClasses[index].CourseId}
		course.SelectById()
		if len(data) == 0 {
			data = append(data, make(map[string]int))
			data[len(data)-1]["count"] = 1
			data[len(data)-1]["year"] = course.Year
		} else {
			for index := 0; index < len(data); index++ {
				if data[index]["year"] == course.Year {
					data[len(data)-1]["count"] = data[len(data)-1]["count"] + 1
					break
				} else if index == len(data)-1 {
					data = append(data, make(map[string]int))
					data[len(data)-1]["count"] = 1
					data[len(data)-1]["year"] = course.Year
				}
			}
		}
	}

	//格式化数据源
	fData := make([]map[string]interface{}, 0, len(data))
	for index := 0; index < len(data); index++ {
		fData = append(fData, make(map[string]interface{}))
		fData[index]["count"] = data[index]["count"]
		fData[index]["year"] = strconv.Itoa(data[index]["year"]) + " 学年"
	}
	return fData
}

//教师分析任课班课程情况
func TeachingclassCount(name string) interface{} {
	teachingClassInformation := model.TeachingClassInformation{CourseTeacherName: name}
	teachingClassInformations := teachingClassInformation.Select()
	data := make([]map[string]int, 0, len(teachingClassInformations))
	for index := 0; index < len(teachingClassInformations); index++ {
		course := model.Course{Id: teachingClassInformations[index].CourseId}
		course.SelectById()
		if len(data) == 0 {
			data = append(data, make(map[string]int))
			data[len(data)-1]["count"] = 1
			data[len(data)-1]["year"] = course.Year
		} else {
			for index := 0; index < len(data); index++ {
				if data[index]["year"] == course.Year {
					data[len(data)-1]["count"] = data[len(data)-1]["count"] + 1
					break
				} else if index == len(data)-1 {
					data = append(data, make(map[string]int))
					data[len(data)-1]["count"] = 1
					data[len(data)-1]["year"] = course.Year
				}
			}
		}
	}

	//格式化数据源
	fData := make([]map[string]interface{}, 0, len(data))
	for index := 0; index < len(data); index++ {
		fData = append(fData, make(map[string]interface{}))
		fData[index]["count"] = data[index]["count"]
		fData[index]["year"] = strconv.Itoa(data[index]["year"]) + " 学年"
	}
	return fData
}
