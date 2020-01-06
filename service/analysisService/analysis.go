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
	teachingClass := model.TeachingClass{StudentId: studentId, Status: 3}
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
	scoreCountRes := teachingClass.ScoreCount()
	for i, v := range scoreCountRes {
		year, _ := strconv.Atoi(v.Year)
		scoreCountRes[i].Year = v.Year + " - " + strconv.Itoa(year+1) + " 学年" + v.Semester
	}
	return scoreCountRes
}

//教师分析任课班课程情况
func TeachingclassCount(name string) interface{} {
	teachingClassInformation := model.TeachingClassInformation{CourseTeacherName: name}
	teachingclassCountRes := teachingClassInformation.ScoreCount()
	for i, v := range teachingclassCountRes {
		year, _ := strconv.Atoi(v.Year)
		teachingclassCountRes[i].Year = v.Year + " - " + strconv.Itoa(year+1) + " 学年 - " + v.Semester
	}
	return teachingclassCountRes
}
