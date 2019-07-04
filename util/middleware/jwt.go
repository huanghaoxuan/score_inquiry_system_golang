package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/**
 * @Author: HuangHaoXuan
 * @Email: huanghaoxuan1998@outlook.com
 * @github https://github.com/huanghaoxuan
 * @Date: 2019/6/23 10:16
 * @Version 1.0
 * 中间件相关
 */

//常量设置
const (
	SecretKey = "东南大学成贤学院成绩查询系统--黄浩轩"
)

//生成Token验证码
func GeneratedToken() string {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	//加入Token过期时间
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(24)).Unix()
	//加入签发时间
	claims["iat"] = time.Now().Unix()
	//加入签发者
	claims["iss"] = SecretKey
	token.Claims = claims

	//生成token字符串
	tokenString, _ := token.SignedString([]byte(SecretKey))
	return tokenString

}

/**
引入上下文 *gin.Context
进行解码验证后返回上下文
*/
func ValidateToken(c *gin.Context) {

	token, err := request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})

	if err == nil {
		if token.Valid {
			//验证通过继续往下走
			c.Next()
		} else {
			//验证不通过
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"msg": "当前登录信息无效或已过期，请重新登录"})
		}
	} else {
		//有错误
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"msg": "当前登录信息无效或已过期，请重新登录"})
	}

}
