package model

/**
 * @Author: HuangHaoXuan
 * @Email: huanghaoxuan1998@outlook.com
 * @github https://github.com/huanghaoxuan
 * @Date: 2019/7/3 15:46
 * @Version 1.0
 */

//课程信息结构体
type Course struct {
	Id       int    `form:"id" gorm:"primary_key;column:id" json:"id"`       //主键
	Name     string `form:"name" gorm:"column:name;not null;" json:"name"`   //课程名字
	Year     int    `form:"year" gorm:"column:year" json:"year"`             //学年
	Semester string `form:"semester" gorm:"column:semester" json:"semester"` //学期
}
