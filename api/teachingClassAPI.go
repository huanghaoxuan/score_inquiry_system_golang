package api

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"score_inquiry_system/model"
	"score_inquiry_system/service/sourceStageService"
	"score_inquiry_system/service/teachingClassService"
	"strconv"
)

/**
 * @Author: HuangHaoXuan
 * @Email: huanghaoxuan1998@outlook.com
 * @github https://github.com/huanghaoxuan
 * @Date: 2019/6/30 15:36
 * @Version 1.0
 */

func TeachingClass(basePath *gin.RouterGroup) {
	basePath.POST("/teachingClass/insert", InsertTeachingClass)
	basePath.POST("/teachingClass/update", UpdateTeachingClass)
	basePath.POST("/teachingClass/upload", UploadTeachingClass)
	basePath.POST("/teachingClass/selectByPage", SelectTeachingClassByPage)
	basePath.GET("/teachingClass/delete/:id", DeleteTeachingClass)
}

// @Summary 删除一条教学班学生信息
// @Description 删除一条教学班信息
// @Tags 教学班信息
// @Accept json
// @Produce json
// @Param Authorization header string true "Token"
// @Param id path string true "主键"
// @Router /teachingClass/delete/{id} [get]
func DeleteTeachingClass(c *gin.Context) {
	id := c.Param("id")
	status := teachingClassService.Delete(id)
	//回调
	c.JSON(http.StatusOK, gin.H{"status": status})
}

// @Summary 分页查询教学班学生信息
// @Description 分页查询教学班信息，如果查询第一页，返回总条数，条件非必需
// @Tags 教学班信息
// @Accept mpfd
// @Produce json
// @Param Authorization header string true "Token"
// @Param pageNum formData string true "查询页码"
// @Param pageSize formData string true "每页条数"
// @Param studentId formData string true "学生学号"
// @Param name formData string false "姓名"
// @Param grade formData string false "所在年级"
// @Param department formData string false "所在学院或部门"
// @Param professional formData string false "所在专业"
// @Param class formData string false "所在班级"
// @Param courseName formData string false "课程名称"
// @Param courseId formData string false "课程id"
// @Param teachingClassId formData string false "教学班号"
// @Param courseTeacherName formData string false "任课老师名字"
// @Param courseTeacherId formData string false "任课老师id"
// @Router /teachingClass/selectByPage [post]
func SelectTeachingClassByPage(c *gin.Context) {
	//模型填充
	var teachingClass model.TeachingClass
	_ = c.ShouldBind(&teachingClass)
	pageNum, _ := strconv.Atoi(c.PostForm("pageNum"))
	pageSize, _ := strconv.Atoi(c.PostForm("pageSize"))
	//查询总条数
	count := teachingClass.Count()
	teachingClasses := teachingClassService.SelectByPage(pageNum, pageSize, &teachingClass)
	//回调
	c.JSON(http.StatusOK, gin.H{"data": teachingClasses, "count": count})
}

// @Summary 更新教学班信息学生记录
// @Description 更新教学班信息记录
// @Tags 教学班信息
// @Accept mpfd
// @Produce json
// @Param Authorization header string true "Token"
// @Param studentId formData string true "学生学号"
// @Param name formData string false "姓名"
// @Param grade formData string false "所在年级"
// @Param department formData string false "所在学院或部门"
// @Param professional formData string false "所在专业"
// @Param class formData string false "所在班级"
// @Param courseName formData string false "课程名称"
// @Param courseId formData string false "课程id"
// @Param teachingClassId formData string false "教学班号"
// @Param courseTeacherName formData string false "任课老师名字"
// @Param courseTeacherId formData string false "任课老师id"
// @Success 200 {string} json "{"status": 1}"
// @Router /teachingClass/update [post]
func UpdateTeachingClass(c *gin.Context) {
	//模型填充
	var teachingClass model.TeachingClass
	_ = c.ShouldBind(&teachingClass)
	//状态回调
	status := teachingClassService.UpdateAll(&teachingClass)
	//回调
	c.JSON(http.StatusOK, gin.H{"status": status})
}

// @Summary 增加教学班信息学生记录
// @Description 增加教学班信息记录
// @Tags 教学班信息
// @Accept mpfd
// @Produce json
// @Param Authorization header string true "Token"
// @Param studentId formData string true "学生学号"
// @Param name formData string false "姓名"
// @Param grade formData string false "所在年级"
// @Param department formData string false "所在学院或部门"
// @Param professional formData string false "所在专业"
// @Param class formData string false "所在班级"
// @Param courseName formData string false "课程名称"
// @Param courseId formData string false "课程id"
// @Param teachingClassId formData string false "教学班号"
// @Param courseTeacherName formData string false "任课老师名字"
// @Param courseTeacherId formData string false "任课老师id"
// @Success 200 {string} json "{"status": 1}"
// @Router /teachingClass/insert [post]
func InsertTeachingClass(c *gin.Context) {
	//模型填充
	var teachingClass model.TeachingClass
	_ = c.ShouldBind(&teachingClass)
	//状态回调
	status := teachingClassService.Insert(&teachingClass)
	var sourceStage model.SourceStage
	//插入阶段性成绩学生信息
	_ = c.ShouldBind(&sourceStage)
	//状态回调
	sourceStageService.InsertStudent(&sourceStage)
	//回调
	c.JSON(http.StatusOK, gin.H{"status": status})
}

// @Summary 上传教学班学生表格文件
// @Description 上传教学班表格文件，批量添加教学班信息
// @Tags 教学班信息
// @Accept mpfd
// @Produce json
// @Param Authorization header string true "Token"
// @Param file formData file true "文件"
// @Success 200 {string} json "{"status": 1}"
// @Router /teachingClass/upload [post]
func UploadTeachingClass(c *gin.Context) {
	courseId := c.PostForm("courseId")
	courseName := c.PostForm("courseName")
	// 单文件
	fileHeader, _ := c.FormFile("file")
	s := "public/teachingClass/" + uuid.NewV4().String() + fileHeader.Filename
	//序列化表格文件
	_ = c.SaveUploadedFile(fileHeader, s)
	//处理表格文件
	teachingClassService.ProcessingExcelFile(s, courseId, courseName)
}
