package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"score_inquiry_system/model"
	"score_inquiry_system/service/sourceStageService"
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

func SourceStage(basePath *gin.RouterGroup) {
	teacher := basePath.Group("")
	teacher.Use(middleware.ValidateTeacherPermissions)
	teacher.POST("/sourceStage/insertStudent", InsertSourceStage)
	teacher.POST("/sourceStage/updates", UpdateSourceStages)
	//basePath.POST("/sourceStage/upload", UploadSourceStage)
	basePath.POST("/sourceStage/selectByPage", SelectSourceStageByPage)
	teacher.GET("/sourceStage/delete/:id", DeleteSourceStage)
}

// @Summary 增加阶段性测验成绩
// @Description 增加阶段性测验成绩
// @Tags 阶段性测验成绩
// @Accept json
// @Produce json
// @Param Authorization header string true "Token"
// @Param name formData string false "学生名字"
// @Param studentId formData string true "学生学号"
// @Param teachingClassId formData string false "教学班号"
// @Param sourceStageId formData string false "阶段性测验id"
// @Param scoresNote formData string false "成绩注释"
// @Param scores formData string false "成绩"
// @Success 200 {string} json "{"status": 1}"
// @Router /sourceStage/insert [post]
func InsertSourceStage(c *gin.Context) {

	//模型填充
	var sourceStage model.SourceStage
	_ = c.ShouldBindJSON(&sourceStage)
	//状态回调
	status := sourceStageService.Insert(&sourceStage)
	//回调
	c.JSON(http.StatusOK, gin.H{"status": status})
}

// @Summary 批量更新阶段性测验成绩
// @Description 批量更新阶段性测验成绩
// @Tags 阶段性测验成绩
// @Accept json
// @Produce json
// @Param Authorization header string true "Token"
// @Param name formData string false "学生名字"
// @Param studentId formData string true "学生学号"
// @Param teachingClassId formData string false "教学班号"
// @Param sourceStageId formData string false "阶段性测验id"
// @Param scoresNote formData string false "成绩注释"
// @Param scores formData string false "成绩"
// @Success 200 {string} json "{"status": 1}"
// @Router /sourceStage/updates [post]
func UpdateSourceStages(c *gin.Context) {

	//获取数组形式的数据
	type sourceStages struct {
		Data []model.SourceStage `form:"data[]" json:"data"`
	}

	//模型填充
	var data sourceStages
	_ = c.ShouldBindJSON(&data)
	//状态回调
	var status int64 = 0
	for _, v := range data.Data {
		status += sourceStageService.UpdateAll(&v)
	}

	//回调
	c.JSON(http.StatusOK, gin.H{"status": status})
}

// @Summary 分页查询阶段性测验成绩
// @Description 分页查询阶段性测验成绩
// @Tags 阶段性测验成绩
// @Accept mpfd
// @Produce json
// @Param Authorization header string true "Token"
// @Param pageNum formData string true "查询页码"
// @Param pageSize formData string true "每页条数"
// @Param name formData string false "学生名字"
// @Param studentId formData string true "学生学号"
// @Param teachingClassId formData string false "教学班号"
// @Param sourceStageId formData string false "阶段性测验id"
// @Param scoresNote formData string false "成绩注释"
// @Param scores formData string false "成绩"
// @Router /sourceStage/selectByPage [post]
func SelectSourceStageByPage(c *gin.Context) {
	//模型填充
	var sourceStage model.SourceStage
	_ = c.ShouldBind(&sourceStage)
	pageNum, _ := strconv.Atoi(c.PostForm("pageNum"))
	pageSize, _ := strconv.Atoi(c.PostForm("pageSize"))
	//查询总条数
	count := sourceStage.Count()
	sourceStages := sourceStageService.SelectByPage(pageNum, pageSize, &sourceStage)
	//回调
	c.JSON(http.StatusOK, gin.H{"data": sourceStages, "count": count})
}

// @Summary 删除阶段性测验成绩
// @Description 删除阶段性测验成绩
// @Tags 阶段性测验成绩
// @Accept json
// @Produce json
// @Param Authorization header string true "Token"
// @Param id path string true "主键"
// @Router /sourceStage/delete/{id} [get]
func DeleteSourceStage(c *gin.Context) {
	id := c.Param("id")
	status := sourceStageService.Delete(id)
	//回调
	c.JSON(http.StatusOK, gin.H{"status": status})
}
