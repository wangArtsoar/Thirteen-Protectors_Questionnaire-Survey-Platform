package router

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/common"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/interfaces"
	"github.com/gin-gonic/gin"
)

// Router 注册中心
func Router() *gin.Engine {
	r := gin.Default()
	// 公共接口
	public := r.Group("/auth")
	public.POST("/login", interfaces.Login())
	public.POST("/register", interfaces.Register())
	// 受保护的接口
	protected := r.Group("/exam")
	protected.Use(common.TokenAuthMiddleware())
	protected.GET("/demo", common.TokenAuthMiddleware(), interfaces.Demo())
	protected.GET("/logout", common.LogoutHandle())
	return r
}
