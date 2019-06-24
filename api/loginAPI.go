package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"score_inquiry_system/util/middleware"
)

/**
 * @Author: HuangHaoXuan
 * @Email: huanghaoxuan1998@outlook.com
 * @github https://github.com/huanghaoxuan
 * @Date: 2019/6/21 19:52
 * @Version 1.0
 */

// @Summary 登录
// @Description 用于用户登录
// @Tags 用户
// @Accept mpfd
// @Produce json
// @Param student_id formData string true "学生学号"
// @Param password formData string true "密码"
// @Success 200 {string} json "  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjEyNjczNDIsImlhdCI6MTU2MTI2Mzc0MiwiaXNzIjoi5Lic5Y2X5aSn5a2m5oiQ6LSk5a2m6Zmi5oiQ57up5p-l6K-i57O757ufLS3pu4TmtanovakifQ.juqOf-lEq8bmWRBg1KHbmaqQK7vJMXJ-R5_tYrJAJs4""
// @Router /user/login [post]
func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": 1, "token": middleware.GeneratedToken()})
}

// @Summary 重置密码
// @Description 将密码重置为身份证后六位
// @Tags 用户
// @Accept json
// @Produce json
// Param Authorization header string true "Token"
// @Param student_id path string true "学生学号"
// @Success 200 {string} json "{"status": 1}"
// @Router /user/reset/{student_id} [get]
func Reset(c *gin.Context) {
	studentID := c.Param("student_id")
	fmt.Println(studentID)
	c.JSON(http.StatusOK, gin.H{"status": 1})
}