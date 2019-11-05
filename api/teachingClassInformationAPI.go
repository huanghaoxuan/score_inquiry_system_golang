package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"score_inquiry_system/model"
	"score_inquiry_system/service/teachingClassInformationService"
	"score_inquiry_system/util/middleware"
	"strconv"
)

/**
 * @Author: HuangHaoXuan
 * @Email: huanghaoxuan1998@outlook.com
 * @github https://github.com/huanghaoxuan
 * @Date: 2019/7/6 20:26
 * @Version 1.0
 */

func TeachingClassInformation(basePath *gin.RouterGroup) {
	teacher := basePath.Group("")
	teacher.Use(middleware.ValidateTeacherPermissions)
	teacher.POST("/teachingClassInformation/insert", InsertTeachingClassInformation)
	teacher.POST("/teachingClassInformation/update", UpdateTeachingClassInformation)
	basePath.POST("/teachingClassInformation/selectByPage", SelectTeachingClassInformationByPage)
	teacher.POST("/teachingClassInformation/selectCrossSemester", SelectCrossSemester)
	teacher.GET("/teachingClassInformation/delete/:id", DeleteTeachingClassInformation)
}

// @Summary 删除一条教学班信息
// @Description 删除一条教学班信息
// @Tags 教学班基本信息
// @Accept json
// @Produce json
// @Param Authorization header string true "Token"
// @Param id path string true "主键"
// @Router /teachingClassInformation/delete/{id} [get]
func DeleteTeachingClassInformation(c *gin.Context) {
	id := c.Param("id")
	status := teachingClassInformationService.Delete(id)
	//回调
	c.JSON(http.StatusOK, gin.H{"status": status})
}

// @Summary 分页查询教学班信息
// @Description 分页查询教学班信息，如果查询第一页，返回总条数，条件非必需
// @Tags 教学班基本信息
// @Accept mpfd
// @Produce json
// @Param Authorization header string true "Token"
// @Param pageNum formData string true "查询页码"
// @Param pageSize formData string true "每页条数"
// @Param courseName formData string false "课程名称"
// @Param courseId formData string false "课程id"
// @Param teachingClassId formData string false "教学班号"
// @Param courseTeacherName formData string false "任课老师名字"
// @Param courseTeacherId formData string false "任课老师id"
// @Router /teachingClassInformation/selectByPage [post]
func SelectTeachingClassInformationByPage(c *gin.Context) {
	//模型填充
	var teachingClassInformation model.TeachingClassInformationResult
	_ = c.ShouldBind(&teachingClassInformation)
	pageNum, _ := strconv.Atoi(c.PostForm("pageNum"))
	pageSize, _ := strconv.Atoi(c.PostForm("pageSize"))
	//查询总条数
	count := teachingClassInformation.Count()
	//将course_id 设置为id返回
	teachingClasseInformations := teachingClassInformationService.SelectByPage(pageNum, pageSize, &teachingClassInformation)
	//回调
	c.JSON(http.StatusOK, gin.H{"data": teachingClasseInformations, "count": count})
}

// @Summary 更新教学班信息记录
// @Description 更新教学班信息记录
// @Tags 教学班基本信息
// @Accept mpfd
// @Produce json
// @Param Authorization header string true "Token"
// @Param courseName formData string false "课程名称"
// @Param courseId formData string false "课程id"
// @Param teachingClassId formData string false "教学班号"
// @Param courseTeacherName formData string false "任课老师名字"
// @Param courseTeacherId formData string false "任课老师id"
// @Success 200 {string} json "{"status": 1}"
// @Router /teachingClassInformation/update [post]
func UpdateTeachingClassInformation(c *gin.Context) {
	//模型填充
	var teachingClassInformation model.TeachingClassInformation
	_ = c.ShouldBind(&teachingClassInformation)
	//状态回调
	status := teachingClassInformationService.UpdateAll(&teachingClassInformation)
	//回调
	c.JSON(http.StatusOK, gin.H{"status": status})
}

// @Summary 增加教学班信息记录
// @Description 增加教学班信息记录
// @Tags 教学班基本信息
// @Accept mpfd
// @Produce json
// @Param Authorization header string true "Token"
// @Param courseName formData string false "课程名称"
// @Param courseId formData string false "课程id"
// @Param teachingClassId formData string false "教学班号"
// @Param courseTeacherName formData string false "任课老师名字"
// @Param courseTeacherId formData string false "任课老师id"
// @Success 200 {string} json "{"status": 1}"
// @Router /teachingClassInformation/insert [post]
func InsertTeachingClassInformation(c *gin.Context) {
	//模型填充
	var teachingClassInformation model.TeachingClassInformation
	_ = c.ShouldBind(&teachingClassInformation)
	//状态回调
	status := teachingClassInformationService.Insert(&teachingClassInformation)
	//回调
	c.JSON(http.StatusOK, gin.H{"status": status})
}

// @Summary 分页查询教学班是否存在跨学期内容
// @Description 分页查询教学班是否存在跨学期内容
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
// @Router /teachingClassInformation/selectCrossSemester [post]
func SelectCrossSemester(c *gin.Context) {
	//模型填充
	var teachingClass model.TeachingClassInformationResult
	_ = c.ShouldBind(&teachingClass)
	pageNum, _ := strconv.Atoi(c.PostForm("pageNum"))
	pageSize, _ := strconv.Atoi(c.PostForm("pageSize"))
	//查询总条数
	teachingClasses, count := teachingClassInformationService.SelectCrossSemester(pageNum, pageSize, &teachingClass)
	//回调
	c.JSON(http.StatusOK, gin.H{"data": teachingClasses, "count": count})
}
