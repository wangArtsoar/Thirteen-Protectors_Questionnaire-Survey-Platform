package interfaces

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/interfaces/ioc"
	vo2 "Thirteen-Protectors_Questionnaire-Survey-Platform/interfaces/vo"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login 登录接口
func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			loginDto      vo2.LoginDto
			err           error
			loginResponse *vo2.LoginResponse
		)
		err = ctx.ShouldBindJSON(&loginDto)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errors.New("参数错误"+err.Error()).Error())
			return
		}
		loginResponse, err = ioc.Container.UserService.Login(&loginDto)
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
		var (
			register vo2.RegisterRequest
			err      error
			response *vo2.RegisterResponse
		)
		err = ctx.ShouldBindJSON(&register)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errors.New("参数错误"+err.Error()).Error())
			return
		}
		service := ioc.Container.UserService
		if service == nil {
			ctx.JSON(http.StatusInternalServerError, errors.New("内部错误，service 获取不到"+err.Error()).Error())
			return
		}
		response, err = ioc.Container.UserService.Register(&register)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errors.New("内部错误"+err.Error()).Error())
			return
		}
		ctx.JSON(http.StatusOK, response)
	}
}
