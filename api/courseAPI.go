package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"score_inquiry_system/model"
	"score_inquiry_system/service/courseService"
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
	basePath.POST("/course/insert", InsertCourse)
	basePath.POST("/course/update", UpdateCourse)
	//basePath.POST("/course/upload", UploadCourse)
	basePath.POST("/course/selectByPage", SelectCourseByPage)
	basePath.GET("/course/delete/:id", DeleteCourse)
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
