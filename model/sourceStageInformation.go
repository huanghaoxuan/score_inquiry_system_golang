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
type SourceStageInformation struct {
	Id              string    `form:"id" gorm:"primary_key;column:id" json:"id"`                                        //主键
	Name            string    `form:"name" gorm:"column:name;not null;" json:"name"`                                    //课程名字
	TeachingClassId string    `form:"teachingClassId" gorm:"column:teaching_class_id;not null;" json:"teachingClassId"` //教学班号
	StageId         string    `form:"stageId" gorm:"column:stage_id;not null;" json:"stageId"`                          //阶段性测验序号
	CourseId        string    `form:"courseId" gorm:"column:course_id" json:"courseId"`                                 //课程id
	StageNote       string    `form:"stageNote" gorm:"column:stage_note;not null;" json:"stageNote"`                    //阶段性测验描述
	Percentage      string    `form:"percentage" gorm:"column:percentage" json:"percentage"`                            //阶段性测验占比
	Type            string    `form:"type" gorm:"column:type" json:"type"`                                              //课程类型
	CreatedAt       time.Time `form:"createdAt" gorm:"column:created_at" json:"createdAt"`                              //创建时间
}

//通过id查询
func (sourceStageInformation *SourceStageInformation) SelectById() *SourceStageInformation {
	db.DB.Where("id = ?", sourceStageInformation.Id).First(&sourceStageInformation)
	return sourceStageInformation
}

//分页查询
func (sourceStageInformation *SourceStageInformation) SelectByPage(pageNum int, pageSize int) []SourceStageInformation {
	sourceStageInformations := make([]SourceStageInformation, 10)
	if pageNum > 0 && pageSize > 0 {
		db.DB.Where(&sourceStageInformation).Order("created_at ASC").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&sourceStageInformations)
	}
	return sourceStageInformations
}

//查询全部
func (sourceStageInformation *SourceStageInformation) SelectAll() []SourceStageInformation {
	sourceStageInformations := make([]SourceStageInformation, 10)
	db.DB.Where(&sourceStageInformation).Order("created_at desc").Find(&sourceStageInformations)
	return sourceStageInformations
}

//查询总记录
func (sourceStageInformation *SourceStageInformation) Count() int {
	count := 0
	db.DB.Model(&sourceStageInformation).Where(&sourceStageInformation).Count(&count)
	return count
}

//插入记录
func (sourceStageInformation *SourceStageInformation) Insert() int64 {
	create := db.DB.Create(&sourceStageInformation)
	return create.RowsAffected
}

//更新记录
//更新相关记录权限
func (sourceStageInformation *SourceStageInformation) Update() int64 {
	if sourceStageInformation.Id != "" {
		updates := db.DB.Model(&sourceStageInformation).Where("id = ?", sourceStageInformation.Id).Updates(sourceStageInformation)
		return updates.RowsAffected
	}
	return 0
}

//更新记录
//全部字段更新
func (sourceStageInformation *SourceStageInformation) UpdateAll() int64 {
	if sourceStageInformation.Id != "" {
		updates := db.DB.Model(&sourceStageInformation).Where("id = ?", sourceStageInformation.Id).Save(sourceStageInformation)
		return updates.RowsAffected
	}
	return 0
}

//删除记录
//通过id删除记录
func (sourceStageInformation *SourceStageInformation) Delete() int64 {
	//防止记录被全部删除
	if sourceStageInformation.Id != "" {
		db.DB.Where("source_stage_id = ?", sourceStageInformation.Id).Delete(SourceStage{})
		i := db.DB.Delete(&sourceStageInformation)
		return i.RowsAffected
	}
	return 0
}
