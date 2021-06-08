package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"score_inquiry_system/model"
	"score_inquiry_system/service/courseService"
	"score_inquiry_system/service/teachingClassInformationService"
	"score_inquiry_system/service/teachingClassService"
	"score_inquiry_system/util/middleware"
	"strconv"
)

/**
 * @Author: HuangHaoXuan
 * @Email: huanghaoxuan1998@outlook.com
 * @github https://github.com/huanghaoxuan
 * @Date: 2019/7/5 13:31
 * @Version 1.0
 */

func Course(basePath *gin.RouterGroup) {
	teacher := basePath.Group("")
	teacher.Use(middleware.ValidateTeacherPermissions)
	teacher.POST("/course/insert", InsertCourse)
	teacher.POST("/course/update", UpdateCourse)
	teacher.POST("/course/releaseCourse", ReleaseCourse)
	basePath.POST("/course/upload", UploadCourse)
	basePath.POST("/course/selectByPage", SelectCourseByPage)
	teacher.GET("/course/delete/:id", DeleteCourse)
}

// @Summary 上传课程信息表格文件
// @Description 上传课程信息表格文件，批量添加学生信息
// @Tags 课程信息
// @Accept mpfd
// @Produce json
// @Param Authorization header string true "Token"
// @Param file formData file true "文件"
// @Success 200 {string} json "{"status": 1}"
// @Router /course/upload [post]
func UploadCourse(c *gin.Context) {
	// 单文件
	fileHeader, _ := c.FormFile("file")
	fmt.Println("./../public/course/" + fileHeader.Filename)
	s := "public/course/" + uuid.NewV4().String() + fileHeader.Filename
	_ = c.SaveUploadedFile(fileHeader, s)

	courseService.ProcessingExcelFile(s)

}

// @Summary 更改课程成绩的发布状态
// @Description 更改教学班成绩的发布状态
// @Tags 教学班基本信息
// @Accept json
// @Produce json
// @Router /course/releaseCourse [post]
func ReleaseCourse(c *gin.Context) {
	status, _ := strconv.Atoi(c.PostForm("status"))
	id := c.PostForm("id")
	teachingClassInformation := model.TeachingClassInformation{Status: status, CourseId: id}
	res := teachingClassInformationService.ReleaseCourse(&teachingClassInformation)

	teachingClass := model.TeachingClass{Status: status, CourseId: id}
	teachingClassService.ReleaseCourse(&teachingClass)

	course := model.Course{Id: id, Status: status}
	course.Update()
	//回调
	c.JSON(http.StatusOK, gin.H{"status": res})
}

// @Summary 增加课程信息记录
// @Description 增加课程信息记录
// @Tags 课程信息
// @Accept mpfd
// @Produce json
// @Param Authorization header string true "Token"
// @Param name formData string true "课程名"
// @Param year formData string false "学年"
// @Param semester formData string false "学期"
// @Success 200 {string} json "{"status": 1}"
// @Router /course/insert [post]
func InsertCourse(c *gin.Context) {
	//模型填充
	var course model.Course
	_ = c.ShouldBind(&course)
	//状态回调
	status := courseService.Insert(&course)
	//回调
	c.JSON(http.StatusOK, gin.H{"status": status})
}

// @Summary 更新课程信息记录
// @Description 更新课程信息记录
// @Tags 课程信息
// @Accept mpfd
// @Produce json
// @Param Authorization header string true "Token"
// @Param name formData string true "课程名"
// @Param year formData string false "学年"
// @Param semester formData string false "学期"
// @Success 200 {string} json "{"status": 1}"
// @Router /course/update [post]
func UpdateCourse(c *gin.Context) {
	//模型填充
	var course model.Course
	_ = c.ShouldBind(&course)
	//状态回调
	status := courseService.UpdateAll(&course)
	//回调
	c.JSON(http.StatusOK, gin.H{"status": status})
}

// @Summary 分页查询课程信息记录
// @Description 分页查询课程信息记录
// @Tags 课程信息
// @Accept mpfd
// @Produce json
// @Param Authorization header string true "Token"
// @Param pageNum formData string true "查询页码"
// @Param pageSize formData string true "每页条数"
// @Param name formData string true "课程名"
// @Param year formData string false "学年"
// @Param semester formData string false "学期"
// @Router /course/selectByPage [post]
func SelectCourseByPage(c *gin.Context) {
	//模型填充
	var course model.Course
	_ = c.ShouldBind(&course)
	pageNum, _ := strconv.Atoi(c.PostForm("pageNum"))
	pageSize, _ := strconv.Atoi(c.PostForm("pageSize"))
	//查询总条数
	count := course.Count()
	courses := courseService.SelectByPage(pageNum, pageSize, &course)
	//回调
	c.JSON(http.StatusOK, gin.H{"data": courses, "count": count})
}

// @Summary 删除一条课程信息
// @Description 删除一条课程信息
// @Tags 学籍信息
// @Accept json
// @Produce json
// @Param Authorization header string true "Token"
// @Param id path string true "主键"
// @Router /course/delete/{id} [get]
func DeleteCourse(c *gin.Context) {
	id := c.Param("id")
	status := courseService.Delete(id)
	//回调
	c.JSON(http.StatusOK, gin.H{"status": status})
}
