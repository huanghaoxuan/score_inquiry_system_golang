package model

import "time"

/**
 * @Author: HuangHaoXuan
 * @Email: huanghaoxuan1998@outlook.com
 * @github https://github.com/huanghaoxuan
 * @Date: 2019/7/9 12:33
 * @Version 1.0
 */

//成绩阶段性测验信息结构体
type SourceStageInformation struct {
	Id              string    `form:"id" gorm:"primary_key;column:id" json:"id"`                                        //主键
	Name            string    `form:"name" gorm:"column:name;not null;" json:"name"`                                    //课程名字
	TeachingClassId string    `form:"teachingClassId" gorm:"column:teaching_class_id;not null;" json:"teachingClassId"` //教学班号
	ScoresId        string    `form:"scoresId" gorm:"column:scores_id;not null;" json:"scoresId"`                       //阶段性测验序号
	ScoresNote      string    `form:"scoresNote" gorm:"column:scores_note;not null;" json:"scoresNote"`                 //阶段性测验描述
	CreatedAt       time.Time `form:"createdAt" gorm:"column:created_at" json:"createdAt"`                              //创建时间
}
