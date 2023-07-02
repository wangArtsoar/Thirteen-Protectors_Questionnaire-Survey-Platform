package interfaces

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/common"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// LoginDto 登录请求体
type LoginDto struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// LoginResponse 登录响应体
type LoginResponse struct {
	Authentication string `json:"authentication"`
}

// Login 登录接口
func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var loginDto LoginDto
		err := ctx.ShouldBindJSON(&loginDto)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errors.New("参数错误"+err.Error()).Error())
			return
		}
		if loginDto.Name != "xiaoyi" || loginDto.Password != "123456" {
			ctx.JSON(http.StatusInternalServerError, errors.New("登录名称或密码错误").Error())
			return
		}
		token := common.CreateNewToken(loginDto.Name)
		ctx.JSON(http.StatusOK, LoginResponse{Authentication: common.Header + token})
	}
}

// Demo 例子
func Demo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "hello demo")
	}
}
