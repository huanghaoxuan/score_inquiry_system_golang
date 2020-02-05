package sourceStageInformationService

import (
	uuid "github.com/satori/go.uuid"
	"score_inquiry_system/model"
)

/**
 * @Author: HuangHaoXuan
 * @Email: huanghaoxuan1998@outlook.com
 * @github https://github.com/huanghaoxuan
 * @Date: 2019/7/10 16:06
 * @Version 1.0
 */

//获取页数
func Count(sourceStageInformation *model.SourceStageInformation) int {
	return sourceStageInformation.Count()
}

//分页查询
func SelectByPage(pageNum int, pageSize int, sourceStageInformation *model.SourceStageInformation) []model.SourceStageInformation {
	return sourceStageInformation.SelectByPage(pageNum, pageSize)
}

//查询全部
func SelectAll(sourceStageInformation *model.SourceStageInformation) []model.SourceStageInformation {
	return sourceStageInformation.SelectAll()
}

func InsertSourceStageInformationByCourseAdministrator(sourceStageInformation *model.SourceStageInformation) int64 {
	teachingClassInformation := model.TeachingClassInformation{
		CourseId: sourceStageInformation.CourseId,
	}
	teachingClasses := teachingClassInformation.Select()
	result := int64(0)
	for _, v := range teachingClasses {
		sourceStageInformationTemp := model.SourceStageInformation{
			Id:              uuid.NewV4().String(),
			Name:            sourceStageInformation.Name,
			TeachingClassId: v.TeachingClassId,
			StageId:         sourceStageInformation.StageId,
			AddPeople:       sourceStageInformation.AddPeople,
			Batch:           sourceStageInformation.Batch,
			CourseId:        sourceStageInformation.CourseId,
			StageNote:       sourceStageInformation.StageNote,
			Percentage:      sourceStageInformation.Percentage,
			Type:            sourceStageInformation.Type,
		}
		result += Insert(&sourceStageInformationTemp)
	}
	return result
}

//插入
func Insert(sourceStageInformation *model.SourceStageInformation) int64 {
	//设置uuid为主键
	sourceStageInformation.Id = uuid.NewV4().String()
	teachingClass := model.TeachingClass{
		TeachingClassId: sourceStageInformation.TeachingClassId,
		CourseId:        sourceStageInformation.CourseId,
	}
	teachingClasses := teachingClass.Select()
	for _, v := range teachingClasses {
		sourceStage := model.SourceStage{
			Name:            v.Name,
			StudentId:       v.StudentId,
			TeachingClassId: v.TeachingClassId,
			SourceStageId:   sourceStageInformation.Id,
			CourseId:        v.CourseId,
		}
		sourceStage.Id = uuid.NewV4().String()
		sourceStage.Insert()
	}
	return sourceStageInformation.Insert()
}

//更新相关记录
func Update(sourceStageInformation *model.SourceStageInformation) int64 {
	return sourceStageInformation.Update()
}

//更新全部字段
func UpdateAll(sourceStageInformation *model.SourceStageInformation) int64 {
	return sourceStageInformation.UpdateAll()
}

//更新全部字段通过
func UpdateByBatch(sourceStageInformation *model.SourceStageInformation) int64 {
	return sourceStageInformation.UpdateByBatch()
}

//通过id查询
func SelectSourceStageInformationById(id string) *model.SourceStageInformation {
	sourceStageInformation := model.SourceStageInformation{Id: id}
	return sourceStageInformation.SelectById()
}

//删除一条记录
func Delete(id string) int64 {
	//获取相关记录，获取学号
	sourceStageInformation := model.SourceStageInformation{Id: id}
	return sourceStageInformation.Delete()
}

//删除一条记录
func DeleteByCourseAdministrator(batch string) int64 {
	//获取相关记录，获取学号
	sourceStageInformation := model.SourceStageInformation{Batch: batch}
	sourceStageInformations := sourceStageInformation.SelectAll()
	result := int64(0)
	for _, v := range sourceStageInformations {
		result += v.Delete()
	}
	return result
}
