package common

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/repository"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

// TokenAuthMiddleware 授权中间件
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token   string
			err     error
			claims  jwt.Claims
			subject string
			flag    bool
		)
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
		claims, err = ExtractJwt(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err})
			return
		}
		// 判断email
		subject, err = claims.GetSubject()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "The user not found"})
			return
		}
		// check if subject from database
		flag, err = repository.UserRepo{}.ExistByEmail(subject)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err})
			return
		}
		if !flag {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "The user not found"})
			return
		}
		// 下一步
		c.Next()
	}
}
