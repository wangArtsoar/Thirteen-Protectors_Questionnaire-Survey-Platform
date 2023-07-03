package interfaces

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/interfaces/ioc"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/vo"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login 登录接口
func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var loginDto vo.LoginDto
		err := ctx.ShouldBindJSON(&loginDto)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errors.New("参数错误"+err.Error()).Error())
			return
		}
		loginResponse, err := ioc.DIContainer.UserService.Login(&loginDto)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errors.New("内部错误"+err.Error()).Error())
			return
		}
		ctx.JSON(http.StatusOK, loginResponse)
	}
}

// Demo 例子
func Demo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "hello demo")
	}
}

// Register 注册
func Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var register vo.RegisterRequest
		err := ctx.ShouldBindJSON(&register)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errors.New("参数错误"+err.Error()).Error())
			return
		}
		response, err := ioc.DIContainer.UserService.Register(&register)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errors.New("内部错误"+err.Error()).Error())
			return
		}
		ctx.JSON(http.StatusOK, response)
	}
}
