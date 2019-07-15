package sourceStageService

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
func Count(sourceStage *model.SourceStage) int {
	return sourceStage.Count()
}

//分页查询
func SelectByPage(pageNum int, pageSize int, sourceStage *model.SourceStage) []model.SourceStage {
	return sourceStage.SelectByPage(pageNum, pageSize)
}

//插入
func Insert(sourceStage *model.SourceStage) int64 {
	//设置uuid为主键
	sourceStage.Id = uuid.NewV4().String()
	return sourceStage.Insert()
}

func InsertStudent(sourceStage *model.SourceStage) int64 {
	sourceStageInformation := model.SourceStageInformation{TeachingClassId: sourceStage.TeachingClassId}
	sourceStageInformations := sourceStageInformation.SelectAll()
	var status int64 = 1
	for _, v := range sourceStageInformations {
		sourceStage2 := model.SourceStage{
			Name:            sourceStage.Name,
			StudentId:       sourceStage.StudentId,
			TeachingClassId: v.TeachingClassId,
			SourceStageId:   v.Id,
		}
		status += Insert(&sourceStage2)
	}
	return status
}

//更新相关记录
func Update(sourceStage *model.SourceStage) int64 {
	return sourceStage.Update()
}

//通过id查询
func SelectSourceStageById(id string) *model.SourceStage {
	sourceStage := model.SourceStage{Id: id}
	return sourceStage.SelectById()
}

//删除一条记录
func Delete(id string) int64 {
	//获取相关记录，获取学号
	sourceStage := model.SourceStage{Id: id}
	return sourceStage.Delete()
}
