package model

import (
	"score_inquiry_system/db"
	"time"
)

/**
 * @Author: HuangHaoXuan
 * @Email: huanghaoxuan1998@outlook.com
 * @github https://github.com/huanghaoxuan
 * @Date: 2019/7/9 12:33
 * @Version 1.0
 */

//成绩阶段性测验信息结构体
type SourceStage struct {
	Id              string    `form:"id" gorm:"primary_key;column:id" json:"id"`                                        //主键
	Name            string    `form:"name" gorm:"column:name;not null;" json:"name"`                                    //名字
	StudentId       string    `form:"studentId" gorm:"column:student_id;not null;unique;" json:"studentId"`             //学生学号
	TeachingClassId string    `form:"teachingClassId" gorm:"column:teaching_class_id;not null;" json:"teachingClassId"` //教学班号
	SourceStageId   string    `form:"sourceStageId" gorm:"column:source_stage_id;not null;" json:"sourceStageId"`       //阶段性测验id
	ScoresNote      string    `form:"scoresNote" gorm:"column:scores_note;not null;" json:"scoresNote"`                 //成绩注释
	Scores          string    `form:"scores" gorm:"column:scores;not null;" json:"scores"`                              //成绩
	CreatedAt       time.Time `form:"createdAt" gorm:"column:created_at" json:"createdAt"`                              //创建时间
}

//通过id查询
func (sourceStage *SourceStage) SelectById() *SourceStage {
	db.DB.Where("id = ?", sourceStage.Id).First(&sourceStage)
	return sourceStage
}

//分页查询
func (sourceStage *SourceStage) SelectByPage(pageNum int, pageSize int) []SourceStage {
	sourceStages := make([]SourceStage, 10)
	if pageNum > 0 && pageSize > 0 {
		db.DB.Where(&sourceStage).Order("created_at desc").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&sourceStages)
	}
	return sourceStages
}

//查询总记录
func (sourceStage *SourceStage) Count() int {
	count := 0
	db.DB.Model(&sourceStage).Where(&sourceStage).Count(&count)
	return count
}

//插入记录
func (sourceStage *SourceStage) Insert() int64 {
	create := db.DB.Create(&sourceStage)
	return create.RowsAffected
}

//更新记录
//更新相关记录权限
func (sourceStage *SourceStage) Update() int64 {
	updates := db.DB.Model(&sourceStage).Where("id = ?", sourceStage.Id).Updates(sourceStage)
	return updates.RowsAffected
}

//删除记录
//通过id删除记录
func (sourceStage *SourceStage) Delete() int64 {
	//防止记录被全部删除
	if sourceStage.Id != "" {
		i := db.DB.Delete(&sourceStage)
		return i.RowsAffected
	}
	return 0
}
