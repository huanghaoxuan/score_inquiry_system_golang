package main

/**
 * @Author: HuangHaoXuan
 * @Email: huanghaoxuan1998@outlook.com
 * @github https://github.com/huanghaoxuan
 * @Date: 2019/6/20 21:50
 * @Version 1.0
 */

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"score_inquiry_system/api"
	"score_inquiry_system/db"
	_ "score_inquiry_system/docs"
	"score_inquiry_system/model"
	"score_inquiry_system/util/middleware"
)

// @title 成绩录入及查询系统标准接口文档
// @version 0.0.1
// @description  东南大学成贤学院学生成绩录入及学生进行查询系统的标准API接口文档
// @termsOfService 暂缺

// @contact.name 黄浩轩
// @contact.url https://github.com/huanghaoxuan
// @contact.email huanghaoxuan1998@outlook.com

// @license.name GNU GENERAL PUBLIC LICENSE Version 3, 29 June 2007
// @license.url https://www.gnu.org/licenses/gpl-3.0.en.html

// @host localhost:5201
// @BasePath /api

func main() {
	//gin.SetMode(gin.ReleaseMode)
	//service.InsertStudent("123456", "123456")

	r := gin.New()
	r.Use(gin.Logger())   //使用Logger中间件
	r.Use(gin.Recovery()) //使用Recovery中间件
	//开启数据库连接池
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//用户登录
	r.POST("/api/user/login", api.Login)
	//用户密码重置
	r.GET("/api/user/reset/:student_id", api.Reset)
	//除登录外全部分组全部加入"/api"前缀
	basePath := r.Group("/api")
	//除登录与文档链接外使用JWT中间件
	basePath.Use(middleware.ValidateToken)
	{
		api.StudentInformation(basePath)
		api.TeachingClass(basePath)
		api.TeacherInformation(basePath)
		api.Course(basePath)
		api.TeachingClassInformation(basePath)
	}
	//数据库结构自动更新
	db.DB.AutoMigrate(&model.Student{}, &model.StudentInformation{}, &model.TeachingClass{}, &model.TeacherInformation{}, &model.TeacherInformation{})
	_ = r.Run(":5201")

}
