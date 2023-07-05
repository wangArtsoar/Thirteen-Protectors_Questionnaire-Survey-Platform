package common

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/repository"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// TokenAuthMiddleware 授权中间件
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取token
		token, err := getAccessToken(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		// 验证 token 是否有效
		_, claims, err := GetMapClaims(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		// 获取email
		name := claims["name"].(string)
		// check if subject from database
		repo := repository.UserRepo{}
		flag, err := repo.ExistByEmail(name)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		if !flag {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "the user not found"})
			return
		}
		// 下一步
		c.Next()
	}
}

// LogoutHandle 退出登录中间件
func LogoutHandle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := getAccessToken(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		// 验证 token 是否有效
		jwtToken, claims, err := GetMapClaims(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		// 设置为退出登录
		claims["isLoggedOut"] = true
		jwtToken.Valid = false
		newToken := CreateNewToken(claims["name"].(string), claims["isLoggedOut"].(bool))
		ctx.Header("Authorization", "Bearer "+newToken)
	}
}

// getAccessToken 获取access token
func getAccessToken(c *gin.Context) (string, error) {
	// 获取jwtToken
	token := c.Request.Header.Get("Authorization")
	// 是否前缀为"Bearer "
	if !strings.HasPrefix(token, "Bearer ") {
		return "", errors.New("invalid API token")
	}
	// 截取7位之后
	token = token[7:]
	if token == "" {
		return "", errors.New("API token required")
	}
	return token, nil
}
