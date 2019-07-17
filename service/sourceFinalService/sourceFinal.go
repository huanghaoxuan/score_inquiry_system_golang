package sourceFinalService

import (
	uuid "github.com/satori/go.uuid"
	"score_inquiry_system/model"
)

/**
 * @Author: HuangHaoXuan
 * @Email: huanghaoxuan1998@outlook.com
 * @github https://github.com/huanghaoxuan
 * @Date: 2019/7/16 11:08
 * @Version 1.0
 */

//获取页数
func Count(sourceFinal *model.SourceFinal) int {
	return sourceFinal.Count()
}

//分页查询
func SelectByPage(pageNum int, pageSize int, sourceFinal *model.SourceFinal) []model.SourceFinal {
	return sourceFinal.SelectByPage(pageNum, pageSize)
}

//插入
func Insert(sourceFinal *model.SourceFinal) int64 {
	//设置uuid为主键
	sourceFinal.Id = uuid.NewV4().String()
	return sourceFinal.Insert()
}

//更新相关记录
func Update(sourceFinal *model.SourceFinal) int64 {
	return sourceFinal.Update()
}

//更新全部字段
func UpdateAll(sourceFinal *model.SourceFinal) int64 {
	return sourceFinal.UpdateAll()
}

//通过id查询
func SelectSourceFinalById(id string) *model.SourceFinal {
	sourceFinal := model.SourceFinal{Id: id}
	return sourceFinal.SelectById()
}

//删除一条记录
func Delete(id string) int64 {
	//获取相关记录
	sourceFinal := model.SourceFinal{Id: id}
	return sourceFinal.Delete()
}
