package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"score_inquiry_system/model"
	"score_inquiry_system/service"
)

/**
 * @Author: HuangHaoXuan
 * @Email: huanghaoxuan1998@outlook.com
 * @github https://github.com/huanghaoxuan
 * @Date: 2019/6/23 20:34
 * @Version 1.0
 */

func StudentInformation(basePath *gin.RouterGroup) {
	basePath.POST("/user/insert", InsertStudentInformation)
	basePath.POST("/user/update", UpdateStudentInformation)
}

// @Summary 增加学生信息记录
// @Description 增加学生信息记录
// @Tags 基本信息
// @Accept mpfd
// @Produce json
// @Param Authorization header string true "Token"
// @Param studentId formData string true "学生学号"
// @Param name formData string false "姓名"
// @Param departmentOld formData string false "原所在学院或部门"
// @Param classOld formData string false "原所在班级"
// @Param gradeOld formData string false "原所在年级"
// @Param departmentNew formData string false "现所在学院或部门"
// @Param classNew formData string false "现所在班级"
// @Param gradeNew formData string false "现所在年级"
// @Success 200 {string} json "{"status": 1}"
// @Router /user/insert [post]
func InsertStudentInformation(c *gin.Context) {
	//模型填充
	var studentInformation model.StudentInformation
	_ = c.ShouldBind(&studentInformation)
	//状态回调
	status := service.Insert(&studentInformation)
	//回调
	c.JSON(http.StatusOK, gin.H{"status": status})
}

// @Summary 更新学生信息记录
// @Description 更新学生信息记录
// @Tags 基本信息
// @Accept mpfd
// @Produce json
// @Param Authorization header string true "Token"
// @Param studentId formData string true "学生学号"
// @Param name formData string false "姓名"
// @Param departmentOld formData string false "原所在学院或部门"
// @Param classOld formData string false "原所在班级"
// @Param gradeOld formData string false "原所在年级"
// @Param departmentNew formData string false "现所在学院或部门"
// @Param classNew formData string false "现所在班级"
// @Param gradeNew formData string false "现所在年级"
// @Param permissions formData string false "权限"
// @Success 200 {string} json "{"status": 1}"
// @Router /user/update [post]
func UpdateStudentInformation(c *gin.Context) {
	//模型填充
	var studentInformation model.StudentInformation
	_ = c.ShouldBind(&studentInformation)
	//状态回调
	status := service.Update(&studentInformation)
	//回调
	c.JSON(http.StatusOK, gin.H{"status": status})
}
