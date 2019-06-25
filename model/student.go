package model

/**
 * @Author: HuangHaoXuan
 * @Email: huanghaoxuan1998@outlook.com
 * @github https://github.com/huanghaoxuan
 * @Date: 2019/6/23 18:29
 * @Version 1.0
 * 账号密码及权限
 */

type Student struct {
	ID          string `form:"id" gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`                                   //主键
	StudentID   string `form:"student_id" gorm:"column:student_id;not null;unique;index:idx_student_id" json:"student_id"` //学生学号、老师工号
	Password    string `form:"password" gorm:"column:password" json:"password"`                                            //密码
	Permissions int    `form:"permissions" gorm:"column:permissions" json:"permissions"`                                   //权限控制
}
