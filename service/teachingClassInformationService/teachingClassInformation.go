package teachingClassInformationService

import (
	uuid "github.com/satori/go.uuid"
	"score_inquiry_system/model"
)

/**
 * @Author: HuangHaoXuan
 * @Email: huanghaoxuan1998@outlook.com
 * @github https://github.com/huanghaoxuan
 * @Date: 2019/7/6 20:15
 * @Version 1.0
 */

//获取页数
func Count(information *model.TeachingClassInformation) int {
	return information.Count()
}

//分页查询
func SelectByPage(pageNum int, pageSize int, information *model.TeachingClassInformation) []model.TeachingClassInformation {
	return information.SelectByPage(pageNum, pageSize)
}

//插入
func Insert(information *model.TeachingClassInformation) int64 {
	//设置uuid为主键
	information.Id = uuid.NewV4().String()
	return information.Insert()
}

//更新相关记录权限
func Update(information *model.TeachingClassInformation) int64 {
	return information.Update()
}

func UpdateAll(information *model.TeachingClassInformation) int64 {
	return information.UpdateAll()
}

//通过id查询
func SelectTeachingClassInformationById(id string) *model.TeachingClassInformation {
	information := model.TeachingClassInformation{Id: id}
	return information.SelectById()
}

//删除一条记录
func Delete(id string) int64 {
	//获取相关记录，获取学号
	information := SelectTeachingClassInformationById(id)
	return information.Delete()
}
