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
	ScoresId        string    `form:"scoresId" gorm:"column:scores_id;not null;" json:"scoresId"`                       //阶段性测验序号
	ScoresNote      string    `form:"scoresNote" gorm:"column:scores_note;not null;" json:"scoresNote"`                 //阶段性测验描述
	CreatedAt       time.Time `form:"createdAt" gorm:"column:created_at" json:"createdAt"`                              //创建时间
}

//通过id查询
func (sourceStageInformation *SourceStageInformation) SelectById() *SourceStageInformation {
	db.DB.Where("id = ?", sourceStageInformation.Id).First(&sourceStageInformation)
	return sourceStageInformation
}

//通过teachingClassId查询
func (sourceStageInformation *SourceStageInformation) SelectByTeachingClassId() *SourceStageInformation {
	db.DB.Where("teaching_class_id = ?", sourceStageInformation.TeachingClassId).Order("created_at desc").First(&sourceStageInformation)
	return sourceStageInformation
}

//分页查询
func (sourceStageInformation *SourceStageInformation) SelectByPage(pageNum int, pageSize int) []SourceStageInformation {
	sourceStageInformations := make([]SourceStageInformation, 10)
	if pageNum > 0 && pageSize > 0 {
		db.DB.Where(&sourceStageInformation).Order("created_at desc").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&sourceStageInformations)
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
	updates := db.DB.Model(&sourceStageInformation).Where("id = ?", sourceStageInformation.Id).Updates(sourceStageInformation)
	return updates.RowsAffected
}

//删除记录
//通过id删除记录
func (sourceStageInformation *SourceStageInformation) Delete() int64 {
	//防止记录被全部删除
	if sourceStageInformation.Id != "" {
		i := db.DB.Delete(&sourceStageInformation)
		return i.RowsAffected
	}
	return 0
}
