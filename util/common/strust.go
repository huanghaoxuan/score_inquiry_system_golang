package common

/**
 * @Author: HuangHaoXuan
 * @Email: huanghaoxuan1998@outlook.com
 * @github https://github.com/huanghaoxuan
 * @Date: 2020/1/6 11:43
 * @Version 1.0
 */

type CountRes struct {
	Count    int    `form:"count" json:"count"`
	Year     string `form:"year" json:"year"`
	Semester string `form:"semester" json:"semester"`
}
