package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"score_inquiry_system/model"
	"score_inquiry_system/service/sourceFinalService"
	"score_inquiry_system/util/middleware"
	"strconv"
)

/**
 * @Author: HuangHaoXuan
 * @Email: huanghaoxuan1998@outlook.com
 * @github https://github.com/huanghaoxuan
 * @Date: 2019/7/10 16:08
 * @Version 1.0
 */

func SourceFinal(basePath *gin.RouterGroup) {
	teacher := basePath.Group("")
	teacher.Use(middleware.ValidateTeacherPermissions)
	teacher.POST("/sourceFinal/insertStudent", InsertSourceFinal)
	teacher.POST("/sourceFinal/updates", UpdateSourceFinals)
	//basePath.POST("/sourceFinal/upload", UploadSourceFinal)
	basePath.POST("/sourceFinal/selectByPage", SelectSourceFinalByPage)
	teacher.GET("/sourceFinal/delete/:id", DeleteSourceFinal)
}

// @Summary 增加期末成绩
// @Description 增加期末成绩
// @Tags 期末成绩
// @Accept json
// @Produce json
// @Param Authorization header string true "Token"
// @Param name formData string true "学生名字"
// @Param studentId formData string true "学生学号"
// @Param teachingClassId formData string false "教学班号"
// @Param scoresNote formData string false "成绩注释"
// @Param scores formData string false "成绩"
// @Success 200 {string} json "{"status": 1}"
// @Router /sourceFinal/insert [post]
func InsertSourceFinal(c *gin.Context) {

	//模型填充
	var sourceFinal model.SourceFinal
	_ = c.ShouldBindJSON(&sourceFinal)
	//状态回调
	status := sourceFinalService.Insert(&sourceFinal)
	//回调
	c.JSON(http.StatusOK, gin.H{"status": status})
}

// @Summary 批量更新期末成绩
// @Description 批量更新期末成绩
// @Tags 期末成绩
// @Accept json
// @Produce json
// @Param Authorization header string true "Token"
// @Param name formData string true "学生名字"
// @Param studentId formData string true "学生学号"
// @Param teachingClassId formData string false "教学班号"
// @Param scoresNote formData string false "成绩注释"
// @Param scores formData string false "成绩"
// @Success 200 {string} json "{"status": 1}"
// @Router /sourceFinal/updates [post]
func UpdateSourceFinals(c *gin.Context) {

	//获取数组形式的数据
	type sourceFinals struct {
		Data []model.SourceFinal `form:"data[]" json:"data"`
	}

	//模型填充
	var data sourceFinals
	_ = c.ShouldBindJSON(&data)
	//状态回调
	var status int64 = 0
	for _, v := range data.Data {
		status += sourceFinalService.UpdateAll(&v)
	}

	//回调
	c.JSON(http.StatusOK, gin.H{"status": status})
}

// @Summary 分页查询期末成绩
// @Description 分页查询期末成绩
// @Tags 期末成绩
// @Accept mpfd
// @Produce json
// @Param Authorization header string true "Token"
// @Param pageNum formData string true "查询页码"
// @Param pageSize formData string true "每页条数"
// @Param name formData string true "学生名字"
// @Param studentId formData string true "学生学号"
// @Param teachingClassId formData string false "教学班号"
// @Param scoresNote formData string false "成绩注释"
// @Param scores formData string false "成绩"
// @Router /sourceFinal/selectByPage [post]
func SelectSourceFinalByPage(c *gin.Context) {
	//模型填充
	var sourceFinal model.SourceFinal
	_ = c.ShouldBind(&sourceFinal)
	pageNum, _ := strconv.Atoi(c.PostForm("pageNum"))
	pageSize, _ := strconv.Atoi(c.PostForm("pageSize"))
	//查询总条数
	count := sourceFinal.Count()
	sourceFinals := sourceFinalService.SelectByPage(pageNum, pageSize, &sourceFinal)
	//回调
	c.JSON(http.StatusOK, gin.H{"data": sourceFinals, "count": count})
}

// @Summary 删除期末成绩
// @Description 删除期末成绩
// @Tags 期末成绩
// @Accept json
// @Produce json
// @Param Authorization header string true "Token"
// @Param id path string true "主键"
// @Router /sourceFinal/delete/{id} [get]
func DeleteSourceFinal(c *gin.Context) {
	id := c.Param("id")
	status := sourceFinalService.Delete(id)
	//回调
	c.JSON(http.StatusOK, gin.H{"status": status})
}
