package common

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/repository/user"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/interfaces/vo"
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
		// 判断 token 是否已经退出登录
		if claims["isLoggedOut"].(bool) {
			c.JSON(http.StatusOK, gin.H{"error": "您已经退出了登录，请重新登录"})
		}
		// 获取email
		name := claims["name"].(string)
		// check if subject from database
		repo := user.UserRepo{}
		if flag, err := repo.ExistByEmail(name); flag && err == nil {
			// 下一步
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User is not be Found" + err.Error()})
			return
		}
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
		_, claims, err := GetMapClaims(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		// 设置为退出登录
		claims["isLoggedOut"] = true
		newToken := CreateNewToken(claims["name"].(string), claims["role"].(string),
			claims["isLoggedOut"].(bool))
		ctx.Header("Authorization", "Bearer "+newToken)
		ctx.JSON(http.StatusOK, &vo.RegisterResponse{
			Message: "退出成功", Authentication: "Bearer " + newToken})
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
