package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// TokenAuthMiddleware 授权中间件
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		// 获取jwtToken
		token = c.Request.Header.Get("Authorization")
		// 是否前缀为"Bearer "
		if !strings.HasPrefix(token, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid API token"})
			return
		}
		// 截取7位之后
		token = token[7:]
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "API token required"})
			return
		}
		// 验证 token 是否有效
		claims, err := ExtractJwt(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err})
			return
		}
		// 判断username
		if username := claims.Username; username != "xiaoyi" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "The user not found"})
			return
		}
		// 下一步
		c.Next()
	}
}
