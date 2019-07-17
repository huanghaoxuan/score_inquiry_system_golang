package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"score_inquiry_system/model"
	"score_inquiry_system/service/sourceStageInformationService"
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

func SourceStageInformation(basePath *gin.RouterGroup) {
	teacher := basePath.Group("")
	teacher.Use(middleware.ValidateTeacherPermissions)
	teacher.POST("/sourceStageInformation/insert", InsertSourceStageInformation)
	teacher.POST("/sourceStageInformation/update", UpdateSourceStageInformation)
	//basePath.POST("/sourceStageInformation/upload", UploadSourceStageInformation)
	basePath.POST("/sourceStageInformation/selectByPage", SelectSourceStageInformationByPage)
	teacher.GET("/sourceStageInformation/delete/:id", DeleteSourceStageInformation)
}

// @Summary 增加阶段性测验信息
// @Description 增加阶段性测验信息
// @Tags 阶段性测验信息
// @Accept mpfd
// @Produce json
// @Param Authorization header string true "Token"
// @Param name formData string true "课程名字"
// @Param teachingClassId formData string false "教学班号"
// @Param scoresId formData string false "阶段性测验序号"
// @Param scoresNote formData string false "阶段性测验描述"
// @Success 200 {string} json "{"status": 1}"
// @Router /sourceStageInformation/insert [post]
func InsertSourceStageInformation(c *gin.Context) {
	//模型填充
	var sourceStageInformation model.SourceStageInformation
	_ = c.ShouldBind(&sourceStageInformation)
	//状态回调
	status := sourceStageInformationService.Insert(&sourceStageInformation)
	//回调
	c.JSON(http.StatusOK, gin.H{"status": status})
}

// @Summary 更新阶段性测验信息
// @Description 更新阶段性测验信息
// @Tags 阶段性测验信息
// @Accept mpfd
// @Produce json
// @Param Authorization header string true "Token"
// @Param name formData string true "课程名字"
// @Param teachingClassId formData string false "教学班号"
// @Param scoresId formData string false "阶段性测验序号"
// @Param scoresNote formData string false "阶段性测验描述"
// @Success 200 {string} json "{"status": 1}"
// @Router /sourceStageInformation/update [post]
func UpdateSourceStageInformation(c *gin.Context) {
	//模型填充
	var sourceStageInformation model.SourceStageInformation
	_ = c.ShouldBind(&sourceStageInformation)
	//状态回调
	status := sourceStageInformationService.UpdateAll(&sourceStageInformation)
	//回调
	c.JSON(http.StatusOK, gin.H{"status": status})
}

// @Summary 查询阶段性测验信息
// @Description 分页查询阶段性测验信息
// @Tags 阶段性测验信息
// @Accept mpfd
// @Produce json
// @Param Authorization header string true "Token"
// @Param pageNum formData string true "查询页码"
// @Param pageSize formData string true "每页条数"
// @Param name formData string true "课程名字"
// @Param teachingClassId formData string false "教学班号"
// @Param scoresId formData string false "阶段性测验序号"
// @Router /sourceStageInformation/selectByPage [post]
func SelectSourceStageInformationByPage(c *gin.Context) {
	//模型填充
	var sourceStageInformation model.SourceStageInformation
	_ = c.ShouldBind(&sourceStageInformation)
	pageNum, _ := strconv.Atoi(c.PostForm("pageNum"))
	pageSize, _ := strconv.Atoi(c.PostForm("pageSize"))
	//查询总条数
	count := sourceStageInformation.Count()
	sourceStageInformations := sourceStageInformationService.SelectByPage(pageNum, pageSize, &sourceStageInformation)
	//回调
	c.JSON(http.StatusOK, gin.H{"data": sourceStageInformations, "count": count})
}

// @Summary 删除阶段性测验信息
// @Description 删除阶段性测验信息
// @Tags 阶段性测验信息
// @Accept json
// @Produce json
// @Param Authorization header string true "Token"
// @Param id path string true "主键"
// @Router /sourceStageInformation/delete/{id} [get]
func DeleteSourceStageInformation(c *gin.Context) {
	id := c.Param("id")
	status := sourceStageInformationService.Delete(id)
	//回调
	c.JSON(http.StatusOK, gin.H{"status": status})
}
