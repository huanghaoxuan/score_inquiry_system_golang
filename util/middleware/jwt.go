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
func GeneratedToken(permissions int) string {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	//加入Token过期时间
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(24)).Unix()
	//加入签发时间
	claims["iat"] = time.Now().Unix()
	//加入签发者
	claims["iss"] = SecretKey
	claims["permissions"] = permissions
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
			//finToken := token.Claims.(jwt.MapClaims)
			//校验下token是否过期
			//succ := finToken.VerifyExpiresAt(time.Now().Unix(), true)
			//获取token中保存的用户信息
			//fmt.Println(succ)
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

func ValidateTeacherPermissions(c *gin.Context) {
	token, _ := request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})
	finToken := token.Claims.(jwt.MapClaims)
	//校验下token是否过期
	//succ := finToken.VerifyExpiresAt(time.Now().Unix(), true)
	//fmt.Println(succ)
	//获取token中保存的用户信息
	if (int)(finToken["permissions"].(float64)) == 1 {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"msg": "该账号权限不允许使用该方法或该请求"})
	} else {
		c.Next()
	}
}
