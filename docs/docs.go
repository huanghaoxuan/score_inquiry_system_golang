// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-06-26 20:57:19.5445933 +0800 CST m=+0.065045501

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "东南大学成贤学院学生成绩录入及学生进行查询系统的标准API接口文档",
        "title": "成绩录入及查询系统标准接口文档",
        "termsOfService": "暂缺",
        "contact": {
            "name": "黄浩轩",
            "url": "https://github.com/huanghaoxuan",
            "email": "huanghaoxuan1998@outlook.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "0.0.1"
    },
    "host": "localhost:5201",
    "basePath": "/api",
    "paths": {
        "/user/insert": {
            "post": {
                "description": "增加学生信息记录",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "基本信息"
                ],
                "summary": "增加学生信息记录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "学生学号",
                        "name": "studentId",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "姓名",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "原所在学院或部门",
                        "name": "departmentOld",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "原所在班级",
                        "name": "classOld",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "原所在年级",
                        "name": "gradeOld",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "现所在学院或部门",
                        "name": "departmentNew",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "现所在班级",
                        "name": "classNew",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "现所在年级",
                        "name": "gradeNew",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\": 1}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "用于用户登录",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "学生学号",
                        "name": "studentId",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "  \"token\": \"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjEyNjczNDIsImlhdCI6MTU2MTI2Mzc0MiwiaXNzIjoi5Lic5Y2X5aSn5a2m5oiQ6LSk5a2m6Zmi5oiQ57up5p-l6K-i57O757ufLS3pu4TmtanovakifQ.juqOf-lEq8bmWRBg1KHbmaqQK7vJMXJ-R5_tYrJAJs4",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/reset/{student_id}": {
            "get": {
                "description": "将密码重置为身份证后六位",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "重置密码",
                "parameters": [
                    {
                        "type": "string",
                        "description": "学生学号",
                        "name": "studentId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\": 1}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/update": {
            "post": {
                "description": "更新学生信息记录",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "基本信息"
                ],
                "summary": "更新学生信息记录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "学生学号",
                        "name": "studentId",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "姓名",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "原所在学院或部门",
                        "name": "departmentOld",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "原所在班级",
                        "name": "classOld",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "原所在年级",
                        "name": "gradeOld",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "现所在学院或部门",
                        "name": "departmentNew",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "现所在班级",
                        "name": "classNew",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "现所在年级",
                        "name": "gradeNew",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "权限",
                        "name": "permissions",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\": 1}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/upload": {
            "post": {
                "description": "上传学生信息表格文件，批量添加学生信息",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "基本信息"
                ],
                "summary": "上传学生信息表格文件",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "文件",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\": 1}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
