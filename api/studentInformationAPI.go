package api

import (
	"fmt"
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
	basePath.POST("/user/upload", UploadStudentInformation)
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

// @Summary 上传学生信息表格文件
// @Description 上传学生信息表格文件，批量添加学生信息
// @Tags 基本信息
// @Accept mpfd
// @Produce json
// @Param Authorization header string true "Token"
// @Param file formData file true "文件"
// @Success 200 {string} json "{"status": 1}"
// @Router /user/upload [post]
func UploadStudentInformation(c *gin.Context) {
	// 单文件
	fileHeader, _ := c.FormFile("file")
	fmt.Println("./../public/studentInformation" + fileHeader.Filename)
	_ = c.SaveUploadedFile(fileHeader, "public/studentInformation/"+fileHeader.Filename)
	//file, _ := excelize.OpenFile("test.xlsx")
	//rows := file.GetRows("Sheet1")
	//for _, row := range rows {
	//	for _, colCell := range row {
	//		fmt.Print(colCell, "\t")
	//	}
	//	fmt.Println()
	//}
	//service.ProcessingExcelFile(file)

}
