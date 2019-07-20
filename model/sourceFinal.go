package model

import (
	"score_inquiry_system/db"
	"time"
)

/**
 * @Author: HuangHaoXuan
 * @Email: huanghaoxuan1998@outlook.com
 * @github https://github.com/huanghaoxuan
 * @Date: 2019/7/16 10:29
 * @Version 1.0
 */

//期末成绩信息结构体
type SourceFinal struct {
	Id              string    `form:"id" gorm:"primary_key;column:id" json:"id"`                                        //主键
	Name            string    `form:"name" gorm:"column:name" json:"name"`                                              //名字
	StudentId       string    `form:"studentId" gorm:"column:student_id;not null;unique;" json:"studentId"`             //学生学号
	ClassId         string    `form:"classId" gorm:"column:class_id;not null;unique" json:"classId"`                    //TeachingClass表中的主键id,唯一
	TeachingClassId string    `form:"teachingClassId" gorm:"column:teaching_class_id;not null;" json:"teachingClassId"` //教学班号
	ScoresNote      string    `form:"scoresNote" gorm:"column:scores_note;not null;" json:"scoresNote"`                 //成绩注释
	Scores          string    `form:"scores" gorm:"column:scores;not null;" json:"scores"`                              //成绩
	CreatedAt       time.Time `form:"createdAt" gorm:"column:created_at" json:"createdAt"`                              //创建时间
}

//通过id查询
func (sourceFinal *SourceFinal) SelectById() *SourceFinal {
	db.DB.Where("id = ?", sourceFinal.Id).First(&sourceFinal)
	return sourceFinal
}

//分页查询
func (sourceFinal *SourceFinal) SelectByPage(pageNum int, pageSize int) []SourceFinal {
	sourceFinals := make([]SourceFinal, 15)
	if pageNum > 0 && pageSize > 0 {
		db.DB.Where(&sourceFinal).Order("created_at desc").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&sourceFinals)
	}
	return sourceFinals
}

//查询总记录
func (sourceFinal *SourceFinal) Count() int {
	count := 0
	db.DB.Model(&sourceFinal).Where(&sourceFinal).Count(&count)
	return count
}

//插入记录
func (sourceFinal *SourceFinal) Insert() int64 {
	create := db.DB.Create(&sourceFinal)
	return create.RowsAffected
}

//更新记录
//更新相关记录权限
func (sourceFinal *SourceFinal) Update() int64 {
	if sourceFinal.Id != "" {
		updates := db.DB.Model(&sourceFinal).Where("id = ?", sourceFinal.Id).Updates(sourceFinal)
		return updates.RowsAffected
	}
	return 0
}

//更新记录
//更新全部字段
func (sourceFinal *SourceFinal) UpdateAll() int64 {
	if sourceFinal.Id != "" {
		updates := db.DB.Model(&sourceFinal).Where("id = ?", sourceFinal.Id).Save(sourceFinal)
		return updates.RowsAffected
	}
	return 0
}

//删除记录
//通过id删除记录
func (sourceFinal *SourceFinal) Delete() int64 {
	//防止记录被全部删除
	if sourceFinal.Id != "" {
		i := db.DB.Delete(&sourceFinal)
		return i.RowsAffected
	}
	return 0
}
