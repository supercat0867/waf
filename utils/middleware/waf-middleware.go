package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"waf/utils/jwt"
)

// AuthUserMiddleware 用户授权中间件
func AuthUserMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") != "" {
			decodedClaims := jwt.VerifyToken(c.GetHeader("Authorization"), secretKey)
			if decodedClaims != nil {
				c.Set("userId", decodedClaims.UserId)
				c.Next()
				c.Abort()
				return
			}
			c.JSON(http.StatusForbidden, gin.H{"msg": "您没有权限访问！"})
			c.Abort()
			return
		} else {
			c.JSON(http.StatusForbidden, gin.H{"msg": "请先登录！"})
		}
		c.Abort()
		return
	}
}
