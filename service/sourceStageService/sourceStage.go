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
func SelectByPage(pageNum int, pageSize int, sourceStage *model.SourceStage) interface{} {
	id := sourceStage.Id
	sourceStage.Id = ""
	sourceStages := sourceStage.SelectByPage(pageNum, pageSize)
	data := make([]map[string]string, 0, len(sourceStages))
	for index := 0; index < len(sourceStages); index++ {
		if sourceStages[index].Name != "" {
			studentId := sourceStages[index].StudentId
			data = append(data, make(map[string]string))
			data[len(data)-1]["name"] = sourceStages[index].Name
			data[len(data)-1]["studentId"] = sourceStages[index].StudentId
			data[len(data)-1]["TeachingClassId"] = sourceStages[index].TeachingClassId
			for i, v := range sourceStages {
				if studentId == v.StudentId {
					if id == v.SourceStageId {
						data[len(data)-1]["scores"] = v.Scores
					}
					data[len(data)-1][v.SourceStageId] = v.Scores
					sourceStages[i].Name = ""
				}
			}
		}
	}
	return data
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

//更新全部字段
func UpdateAll(sourceStage *model.SourceStage) int64 {
	return sourceStage.UpdateAll()
}

//通过id查询
func SelectSourceStageById(id string) *model.SourceStage {
	sourceStage := model.SourceStage{Id: id}
	return sourceStage.SelectById()
}

//删除一条记录
func Delete(id string) int64 {
	//获取相关记录
	sourceStage := model.SourceStage{Id: id}
	return sourceStage.Delete()
}

func ShowSourceStage(pageNum int, pageSize int, sourceStage *model.SourceStage) interface{} {
	return sourceStage.ShowSourceStage(pageNum, pageSize)
}
