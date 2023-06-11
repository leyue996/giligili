package middleware

import (
	"giligili/app/http/pkg/e"
	"giligili/app/http/pkg/util"
	"github.com/gin-gonic/gin"
	"time"
)

// 中间件,登录保护,token正确才能访问下去
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = e.Success
		token := c.GetHeader("access_token")
		if token == "" {
			code = e.Error
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthToken
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorAuthCheckTokenTimeout
			}
		}
		if code != e.Success {
			c.JSON(code, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
