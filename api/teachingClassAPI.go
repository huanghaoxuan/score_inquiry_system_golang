package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"score_inquiry_system/model"
	"score_inquiry_system/service/sourceStageService"
	"score_inquiry_system/service/teachingClassService"
	"score_inquiry_system/util/middleware"
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
	teacher := basePath.Group("")
	teacher.Use(middleware.ValidateTeacherPermissions)
	teacher.POST("/teachingClass/insert", InsertTeachingClass)
	teacher.POST("/teachingClass/update", UpdateTeachingClass)
	teacher.POST("/teachingClass/upload", UploadTeachingClass)
	basePath.POST("/teachingClass/selectByPage", SelectTeachingClassByPage)
	basePath.POST("/teachingClass/showFinal", ShowFinal)
	basePath.POST("/teachingClass/selectFinal", SelectFinal)
	teacher.GET("/teachingClass/delete/:id", DeleteTeachingClass)
	teacher.POST("/teachingClass/updateFinal", UpdateFinal)
	teacher.GET("/teachingClass/download/:teachingClassId", DownloadFinalScore)
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

// @Summary 查询期末成绩信息
// @Description 查询期末成绩信息
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
func ShowFinal(c *gin.Context) {
	//模型填充
	var teachingClass model.TeachingClass
	_ = c.ShouldBind(&teachingClass)
	pageNum, _ := strconv.Atoi(c.PostForm("pageNum"))
	pageSize, _ := strconv.Atoi(c.PostForm("pageSize"))
	//查询总条数
	count := teachingClass.Count()
	teachingClasses := teachingClassService.ShowFinal(pageNum, pageSize, &teachingClass)
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
	//开启协程，插入期末成绩
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

// @Summary 查询期末成绩信息
// @Description 分页查询期末成绩信息，如果查询第一页，返回总条数，条件非必需
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
// @Router /teachingClass/selectFinal [post]
func SelectFinal(c *gin.Context) {
	//模型填充
	var teachingClass model.TeachingClass
	_ = c.ShouldBind(&teachingClass)
	pageNum, _ := strconv.Atoi(c.PostForm("pageNum"))
	pageSize, _ := strconv.Atoi(c.PostForm("pageSize"))
	//查询总条数
	count := teachingClass.Count()
	teachingClasses := teachingClassService.SelectFinal(pageNum, pageSize, &teachingClass)
	//回调
	c.JSON(http.StatusOK, gin.H{"data": teachingClasses, "count": count})
}

// @Summary 更新成绩信息
// @Description 更新成绩信息
// @Tags 教学班信息
// @Accept json
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
// @Success 200 {string} json "{"status": 1}"
// @Router /teachingClass/updateFinal [post]
func UpdateFinal(c *gin.Context) {

	//获取数组形式的数据
	type teachingClasses struct {
		Data []model.TeachingClass `form:"data[]" json:"data"`
	}

	//模型填充
	var data teachingClasses
	_ = c.ShouldBindJSON(&data)
	//状态回调
	var status int64 = 0
	for _, v := range data.Data {
		status += teachingClassService.Update(&v)
	}

	//回调
	c.JSON(http.StatusOK, gin.H{"status": status})
}

// @Summary 下载成绩信息
// @Description 下载成绩信息
// @Tags 教学班信息
// @Accept json
// @Produce json
// @Param Authorization header string true "Token"
// @Param id path string true "主键"
// @Router /teachingClass/download/{teachingClassId} [get]
func DownloadFinalScore(c *gin.Context) {
	teachingClassId := c.Param("teachingClassId")
	filename := teachingClassService.GeneratedExcel(teachingClassId)
	//fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File("public/finalScore/" + filename)
	//回调
	//c.JSON(http.StatusOK, gin.H{"status": 1})
}
