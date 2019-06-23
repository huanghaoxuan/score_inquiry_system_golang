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
	basePath.POST("/user/insert", Insert)
	basePath.POST("/user/update", Update)
}

// @Summary 增加学生信息记录
// @Description 增加学生信息记录
// @Tags 基本信息
// @Accept mpfd
// @Produce json
// @Param Authorization header string true "Token"
// @Param student_id formData string true "学生学号"
// @Param name formData string false "姓名"
// @Param department formData string false "所在学院或部门"
// @Param class formData string false "所在班级"
// @Param entering_time formData string false "入校时间"
// @Success 200 {string} json "{"status": 1}"
// @Router /user/insert [post]
func Insert(c *gin.Context) {
	var studentInformation model.StudentInformation
	_ = c.ShouldBind(&studentInformation)
	service.Insert(&studentInformation)
	c.JSON(http.StatusOK, gin.H{"status": 1})
}

// @Summary 更新学生信息记录
// @Description 更新学生信息记录
// @Tags 基本信息
// @Accept mpfd
// @Produce json
// @Param Authorization header string true "Token"
// @Param student_id formData string true "学生学号"
// @Param name formData string false "姓名"
// @Param department formData string false "所在学院或部门"
// @Param class formData string false "所在班级"
// @Param entering_time formData string false "入校时间"
// @Param permissions formData string false "权限"
// @Success 200 {string} json "{"status": 1}"
// @Router /user/update [post]
func Update(c *gin.Context) {
	var studentInformation model.StudentInformation
	_ = c.ShouldBind(&studentInformation)
	service.Update(&studentInformation)
	c.JSON(http.StatusOK, gin.H{"status": 1})
}
