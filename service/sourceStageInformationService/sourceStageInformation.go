package sourceStageInformationService

import (
	uuid "github.com/satori/go.uuid"
	"score_inquiry_system/model"
	"score_inquiry_system/service/sourceStageService"
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

//插入
func Insert(sourceStageInformation *model.SourceStageInformation) int64 {
	//设置uuid为主键
	sourceStageInformation.Id = uuid.NewV4().String()
	teachingClass := model.TeachingClass{TeachingClassId: sourceStageInformation.TeachingClassId}
	teachingClasses := teachingClass.SelectAll()
	for _, v := range teachingClasses {
		sourceStage := model.SourceStage{
			Name:            v.Name,
			StudentId:       v.StudentId,
			TeachingClassId: v.TeachingClassId,
			SourceStageId:   sourceStageInformation.Id,
		}
		sourceStageService.Insert(&sourceStage)
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
