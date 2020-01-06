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
	Name            string    `form:"name" gorm:"column:name" json:"name"`                                              //名字
	StudentId       string    `form:"studentId" gorm:"column:student_id;not null;" json:"studentId"`                    //学生学号
	TeachingClassId string    `form:"teachingClassId" gorm:"column:teaching_class_id;not null;" json:"teachingClassId"` //教学班号
	SourceStageId   string    `form:"sourceStageId" gorm:"column:source_stage_id;not null;" json:"sourceStageId"`       //阶段性测验id
	CourseId        string    `form:"courseId" gorm:"column:course_id" json:"courseId"`
	ScoresNote      string    `form:"scoresNote" gorm:"column:scores_note;not null;" json:"scoresNote"` //成绩注释
	Scores          string    `form:"scores" gorm:"column:scores;not null;" json:"scores"`              //成绩
	CreatedAt       time.Time `form:"createdAt" gorm:"column:created_at" json:"createdAt"`              //创建时间
}

type SourceStageResult struct {
	SourceStage
	Percentage string `json:"percentage"`
}

//通过id查询
func (sourceStage *SourceStage) SelectById() *SourceStage {
	db.DB.Where("id = ?", sourceStage.Id).First(&sourceStage)
	return sourceStage
}

//查询
func (sourceStage *SourceStage) SelectOne() *SourceStage {
	db.DB.Where(&sourceStage).First(&sourceStage)
	return sourceStage
}

//分页查询
func (sourceStage *SourceStage) SelectByPage(pageNum int, pageSize int) []SourceStage {
	sourceStages := make([]SourceStage, 15)
	if pageNum > 0 && pageSize > 0 {
		db.DB.Where(&sourceStage).Order("created_at desc").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&sourceStages)
	}
	return sourceStages
}

//全部查询
func (sourceStage *SourceStage) SelectAll() []SourceStage {
	sourceStages := make([]SourceStage, 15)
	db.DB.Where(&sourceStage).Order("student_id ASC").Find(&sourceStages)
	return sourceStages
}

//全部查询
func (sourceStage *SourceStageResult) SelectAll() []SourceStageResult {
	sourceStages := make([]SourceStageResult, 15)
	db.DB.Table(" source_stage ss ").
		Select(" ss.*,ssi.percentage ").
		Where(" ss.course_id = ? and ss.teaching_class_id = ?", sourceStage.CourseId, sourceStage.TeachingClassId).
		Joins(" left join source_stage_information ssi on ssi.id = ss.source_stage_id ").
		Order("ss.student_id ASC").
		Find(&sourceStages)
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
	if sourceStage.Id != "" {
		updates := db.DB.Model(&sourceStage).Where("id = ?", sourceStage.Id).Updates(sourceStage)
		return updates.RowsAffected
	}
	return 0
}

//更新记录
//更新相关记录权限
//阶段性成绩更新
func (sourceStage *SourceStage) UpdateSourceStage() int64 {
	if sourceStage.SourceStageId != "" && sourceStage.Id == "" {
		updates := db.DB.Model(&sourceStage).Where("source_stage_id = ? AND course_id = ? AND teaching_class_id = ? AND student_id = ?",
			sourceStage.SourceStageId,
			sourceStage.CourseId,
			sourceStage.TeachingClassId,
			sourceStage.StudentId).Updates(sourceStage)
		return updates.RowsAffected
	} else if sourceStage.SourceStageId != "" {
		updates := db.DB.Model(&sourceStage).Where("id = ? and source_stage_id = ? AND course_id = ? AND teaching_class_id = ? AND student_id = ?",
			sourceStage.Id,
			sourceStage.SourceStageId,
			sourceStage.CourseId,
			sourceStage.TeachingClassId,
			sourceStage.StudentId).Updates(sourceStage)
		return updates.RowsAffected
	}
	return 0
}

//更新记录
//更新全部字段
func (sourceStage *SourceStage) UpdateAll() int64 {
	if sourceStage.Id != "" {
		updates := db.DB.Model(&sourceStage).Where("id = ?", sourceStage.Id).Save(sourceStage)
		return updates.RowsAffected
	}
	return 0
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

//期末成绩查询
func (sourceStage *SourceStage) ShowSourceStage(pageNum int, pageSize int) interface{} {
	type sourceStageRes struct {
		SourceStage
		StageId   string `json:"stageId"`   //阶段性测验序号
		StageNote string `json:"stageNote"` //阶段性测验描述
	}
	stage := []sourceStageRes{}
	db.DB.Table("source_stage").
		Select("source_stage.* , source_stage_information.*").
		Where("source_stage.student_id = ? and source_stage.teaching_class_id = ?", sourceStage.StudentId, sourceStage.TeachingClassId).
		Joins("left join source_stage_information on source_stage.source_stage_id = source_stage_information.id").
		Order("source_stage.created_at desc").Limit(pageSize).Offset((pageNum - 1) * pageSize).
		Find(&stage)
	//SELECT source_final.* , teaching_class_information.course_name , course.* FROM `source_final` left join teaching_class_information on source_final.teaching_class_id = teaching_class_information.teaching_class_id left join course on teaching_class_information.course_name = course.name WHERE (source_final.student_id = '38216137') ORDER BY source_final.created_at desc LIMIT 9 OFFSET 0
	return stage
}
