package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"score_inquiry_system/model"
	"score_inquiry_system/service/studentInformationService"
	"strconv"
)

/**
 * @Author: HuangHaoXuan
 * @Email: huanghaoxuan1998@outlook.com
 * @github https://github.com/huanghaoxuan
 * @Date: 2019/6/23 20:34
 * @Version 1.0
 */

func StudentInformation(basePath *gin.RouterGroup) {
	basePath.POST("/studentInformation/insert", InsertStudentInformation)
	basePath.POST("/studentInformation/update", UpdateStudentInformation)
	basePath.POST("/studentInformation/upload", UploadStudentInformation)
	basePath.POST("/studentInformation/selectByPage", SelectStudentInformationByPage)
	basePath.GET("/studentInformation/delete/:id", DeleteStudentInformation)
}

// @Summary 增加学生信息记录
// @Description 增加学生信息记录
// @Tags 学籍信息
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
// @Router /studentInformation/insert [post]
func InsertStudentInformation(c *gin.Context) {
	//模型填充
	var studentInformation model.StudentInformation
	_ = c.ShouldBind(&studentInformation)
	//状态回调
	status := studentInformationService.Insert(&studentInformation)
	//回调
	c.JSON(http.StatusOK, gin.H{"status": status})
}

// @Summary 更新学生信息记录
// @Description 更新学生信息记录
// @Tags 学籍信息
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
// @Router /studentInformation/update [post]
func UpdateStudentInformation(c *gin.Context) {
	//模型填充
	var studentInformation model.StudentInformation
	_ = c.ShouldBind(&studentInformation)
	//状态回调
	status := studentInformationService.Update(&studentInformation)
	//回调
	c.JSON(http.StatusOK, gin.H{"status": status})
}

// @Summary 上传学生信息表格文件
// @Description 上传学生信息表格文件，批量添加学生信息
// @Tags 学籍信息
// @Accept mpfd
// @Produce json
// @Param Authorization header string true "Token"
// @Param file formData file true "文件"
// @Success 200 {string} json "{"status": 1}"
// @Router /studentInformation/upload [post]
func UploadStudentInformation(c *gin.Context) {
	// 单文件
	fileHeader, _ := c.FormFile("file")
	fmt.Println("./../public/studentInformation" + fileHeader.Filename)
	s := "public/studentInformation/" + uuid.NewV4().String() + fileHeader.Filename
	_ = c.SaveUploadedFile(fileHeader, s)

	studentInformationService.ProcessingExcelFile(s)

}

// @Summary 分页查询学生信息
// @Description 分页查询学生信息，如果查询第一页，返回总条数，条件非必需
// @Tags 学籍信息
// @Accept json
// @Produce json
// @Param Authorization header string true "Token"
// @Param pageNum formData string true "查询页码"
// @Param pageSize formData string true "每页条数"
// @Param studentId formData string false "学生学号"
// @Param name formData string false "姓名"
// @Param departmentNew formData string false "现所在学院或部门"
// @Param classNew formData string false "现所在班级"
// @Param gradeNew formData string false "现所在年级"
// @Router /studentInformation/selectByPage [post]
func SelectStudentInformationByPage(c *gin.Context) {
	//模型填充
	var studentInformation model.StudentInformation
	_ = c.ShouldBind(&studentInformation)
	pageNum, _ := strconv.Atoi(c.PostForm("pageNum"))
	pageSize, _ := strconv.Atoi(c.PostForm("pageSize"))
	//查询总条数
	count := studentInformation.Count()
	studentInformations := studentInformationService.SelectByPage(pageNum, pageSize, &studentInformation)
	//回调
	c.JSON(http.StatusOK, gin.H{"data": studentInformations, "count": count})
}

// @Summary 删除一条学生信息
// @Description 删除一条学生信息
// @Tags 学籍信息
// @Accept json
// @Produce json
// @Param Authorization header string true "Token"
// @Param id path string true "主键"
// @Router /studentInformation/delete/{id} [get]
func DeleteStudentInformation(c *gin.Context) {
	id := c.Param("id")
	status := studentInformationService.Delete(id)

	//回调
	c.JSON(http.StatusOK, gin.H{"status": status})
}
