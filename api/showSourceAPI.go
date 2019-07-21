package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"score_inquiry_system/model"
	"score_inquiry_system/service/sourceFinalService"
	"score_inquiry_system/service/sourceStageService"
	"strconv"
)

/**
 * @Author: HuangHaoXuan
 * @Email: huanghaoxuan1998@outlook.com
 * @github https://github.com/huanghaoxuan
 * @Date: 2019/7/21 14:16
 * @Version 1.0
 */

func ShowSource(basePath *gin.RouterGroup) {
	basePath.POST("/ShowSource/Stage", ShowSourceStageByPage)
	basePath.POST("/ShowSource/final", ShowSourceFinalByPage)
}

// @Summary 阶段性成绩查询
// @Description 学生进行阶段性成绩查询
// @Tags 成绩查询
// @Accept mpfd
// @Produce json
// @Param Authorization header string true "Token"
// @Param pageNum formData string true "查询页码"
// @Param pageSize formData string true "每页条数"
// @Param name formData string false "学生名字"
// @Param studentId formData string true "学生学号"
// @Param teachingClassId formData string true "教学班号"
// @Param sourceStageId formData string false "阶段性测验id"
// @Param scoresNote formData string false "成绩注释"
// @Param scores formData string false "成绩"
// @Router /ShowSource/Stage [post]
func ShowSourceStageByPage(c *gin.Context) {
	//模型填充
	var sourceStage model.SourceStage
	_ = c.ShouldBind(&sourceStage)
	pageNum, _ := strconv.Atoi(c.PostForm("pageNum"))
	pageSize, _ := strconv.Atoi(c.PostForm("pageSize"))
	//查询总条数
	count := sourceStage.Count()
	sourceStages := sourceStageService.ShowSourceStage(pageNum, pageSize, &sourceStage)
	//回调
	c.JSON(http.StatusOK, gin.H{"data": sourceStages, "count": count})
}

// @Summary 期末成绩查询
// @Description 学生进行期末成绩查询
// @Tags 成绩查询
// @Accept mpfd
// @Produce json
// @Param Authorization header string true "Token"
// @Param pageNum formData string true "查询页码"
// @Param pageSize formData string true "每页条数"
// @Param name formData string false "学生名字"
// @Param studentId formData string true "学生学号"
// @Param teachingClassId formData string false "教学班号"
// @Param scoresNote formData string false "成绩注释"
// @Param scores formData string false "成绩"
// @Router /ShowSource/final [post]
func ShowSourceFinalByPage(c *gin.Context) {
	//模型填充
	var sourceFinal model.SourceFinal
	_ = c.ShouldBind(&sourceFinal)
	pageNum, _ := strconv.Atoi(c.PostForm("pageNum"))
	pageSize, _ := strconv.Atoi(c.PostForm("pageSize"))
	//查询总条数
	count := sourceFinal.Count()
	sourceFinals := sourceFinalService.ShowSourceFinal(pageNum, pageSize, &sourceFinal)
	//回调
	c.JSON(http.StatusOK, gin.H{"data": sourceFinals, "count": count})
}
