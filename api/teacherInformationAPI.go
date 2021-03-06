package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"score_inquiry_system/model"
	"score_inquiry_system/service/teacherInformationService"
	"score_inquiry_system/util/middleware"
	"strconv"
)

/**
 * @Author: HuangHaoXuan
 * @Email: huanghaoxuan1998@outlook.com
 * @github https://github.com/huanghaoxuan
 * @Date: 2019/6/23 20:34
 * @Version 1.0
 */

func TeacherInformation(basePath *gin.RouterGroup) {
	teacher := basePath.Group("")
	teacher.Use(middleware.ValidateTeacherPermissions)
	teacher.POST("/teacherInformation/insert", InsertTeacherInformation)
	teacher.POST("/teacherInformation/update", UpdateTeacherInformation)
	teacher.POST("/teacherInformation/upload", UploadTeacherInformation)
	teacher.POST("/teacherInformation/selectByPage", SelectTeacherInformationByPage)
	teacher.GET("/teacherInformation/delete/:id", DeleteTeacherInformation)
	teacher.GET("/teacherInformation/selectByName/:name", SelectTeacherInformationByName)
}

// @Summary 上传教师信息表格文件
// @Description 上传教师信息表格文件，批量添加学生信息
// @Tags 老师信息
// @Accept mpfd
// @Produce json
// @Param Authorization header string true "Token"
// @Param file formData file true "文件"
// @Success 200 {string} json "{"status": 1}"
// @Router /studentInformation/upload [post]
func UploadTeacherInformation(c *gin.Context) {
	// 单文件
	fileHeader, _ := c.FormFile("file")
	fmt.Println("./../public/teacherInformation" + fileHeader.Filename)
	s := "public/studentInformation/" + uuid.NewV4().String() + fileHeader.Filename
	_ = c.SaveUploadedFile(fileHeader, s)

	teacherInformationService.ProcessingExcelFile(s)

}

// @Summary 增加老师信息记录
// @Description 增加老师信息记录
// @Tags 老师信息
// @Accept mpfd
// @Produce json
// @Param Authorization header string true "Token"
// @Param teacherId formData string true "教师工号"
// @Param name formData string false "姓名"
// @Param department formData string false "所在学院或部门"
// @Success 200 {string} json "{"status": 1}"
// @Router /teacherInformation/insert [post]
func InsertTeacherInformation(c *gin.Context) {
	//模型填充
	var teacherInformation model.TeacherInformation
	_ = c.ShouldBind(&teacherInformation)
	//状态回调
	status := teacherInformationService.Insert(&teacherInformation)
	//回调
	c.JSON(http.StatusOK, gin.H{"status": status})
}

// @Summary 更新老师信息记录
// @Description 更新老师信息记录
// @Tags 老师信息
// @Accept mpfd
// @Produce json
// @Param Authorization header string true "Token"
// @Param teacherId formData string true "教师工号"
// @Param name formData string false "姓名"
// @Param department formData string false "所在学院或部门"
// @Success 200 {string} json "{"status": 1}"
// @Router /teacherInformation/update [post]
func UpdateTeacherInformation(c *gin.Context) {
	//模型填充
	var teacherInformation model.TeacherInformation
	_ = c.ShouldBind(&teacherInformation)
	//状态回调
	status := teacherInformationService.Update(&teacherInformation)
	//回调
	c.JSON(http.StatusOK, gin.H{"status": status})
}

// @Summary 分页查询老师信息
// @Description 分页查询老师信息，如果查询第一页，返回总条数，条件非必需
// @Tags 老师信息
// @Accept mpfd
// @Produce json
// @Param Authorization header string true "Token"
// @Param pageNum formData string true "查询页码"
// @Param pageSize formData string true "每页条数"
// @Param teacherId formData string false "教师工号"
// @Param name formData string false "姓名"
// @Param department formData string false "所在学院或部门"
// @Router /teacherInformation/selectByPage [post]
func SelectTeacherInformationByPage(c *gin.Context) {

	//模型填充
	var teacherInformation model.TeacherInformation
	_ = c.ShouldBind(&teacherInformation)
	pageNum, _ := strconv.Atoi(c.PostForm("pageNum"))
	pageSize, _ := strconv.Atoi(c.PostForm("pageSize"))
	//查询总条数
	count := teacherInformation.Count()
	teacherInformations := teacherInformationService.SelectByPage(pageNum, pageSize, &teacherInformation)
	//回调
	c.JSON(http.StatusOK, gin.H{"data": teacherInformations, "count": count})
}

// @Summary 通过名字查询老师信息
// @Description
// @Tags 老师信息
// @Accept json
// @Produce json
// @Param name formData string false "姓名"
// @Router /teacherInformation/SelectTeacherInformationByName/{name} [get]
func SelectTeacherInformationByName(c *gin.Context) {
	//模型填充
	name := c.Param("name")
	teacherInformations := teacherInformationService.SelectTeacherInformationByName(name)
	//回调
	c.JSON(http.StatusOK, gin.H{"data": teacherInformations})
}

// @Summary 删除一条老师信息
// @Description 删除一条老师信息
// @Tags 老师信息
// @Accept json
// @Produce json
// @Param Authorization header string true "Token"
// @Param id path string true "主键"
// @Router /teacherInformation/delete/{id} [get]
func DeleteTeacherInformation(c *gin.Context) {
	id := c.Param("id")
	status := teacherInformationService.Delete(id)

	//回调
	c.JSON(http.StatusOK, gin.H{"status": status})
}
