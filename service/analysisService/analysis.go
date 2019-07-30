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
func AnalysisPass(studentId string) interface{} {
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
		result, _ := strconv.Atoi(v.Result)
		switch {
		case result == 100:
			a++
		case result >= 85:
			b++
		case result >= 60:
			c++
		default:
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
