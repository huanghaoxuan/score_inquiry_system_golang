package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"score_inquiry_system/service/analysisService"
)

/**
 * @Author: HuangHaoXuan
 * @Email: huanghaoxuan1998@outlook.com
 * @github https://github.com/huanghaoxuan
 * @Date: 2019/7/30 14:08
 * @Version 1.0
 */

func Analysis(basePath *gin.RouterGroup) {
	basePath.GET("/analysis/pass/:studentId", Pass)
	basePath.GET("/analysis/score/:studentId", ScoreCount)
}

// @Summary 分析及格情况
// @Description 分析及格情况，分四等，满分，优秀，及格，不及格
// @Tags 首页数据源
// @Accept json
// @Produce json
// @Param Authorization header string true "Token"
// @Param studentId path string true "学生id"
// @Router /analysis/pass/{studentId} [get]
func Pass(c *gin.Context) {
	studentId := c.Param("studentId")
	data := analysisService.Pass(studentId)
	//回调
	c.JSON(http.StatusOK, gin.H{"data": data})
}

// @Summary 分析每年课程情况
// @Description 分析每年课程情况
// @Tags 首页数据源
// @Accept json
// @Produce json
// @Param Authorization header string true "Token"
// @Param studentId path string true "学生id"
// @Router /analysis/score/{studentId} [get]
func ScoreCount(c *gin.Context) {
	studentId := c.Param("studentId")
	data := analysisService.ScoreCount(studentId)
	//回调
	c.JSON(http.StatusOK, gin.H{"data": data})
}
