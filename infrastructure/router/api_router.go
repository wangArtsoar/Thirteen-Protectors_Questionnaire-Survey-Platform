package router

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/common"
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

	server := r.Group("/server")
	server.Use(common.TokenAuthMiddleware())
	server.POST("/save", common.TokenAuthMiddleware(), interfaces.SaveServer())
	server.GET("/findAllCurrentUser", common.TokenAuthMiddleware(), interfaces.FindAllServerByUserEmail())

	channel := r.Group("/channel")
	channel.GET("/findAllByServerId", common.TokenAuthMiddleware(), interfaces.FindAllChannelByServer())
	channel.POST("/Save", common.TokenAuthMiddleware(), interfaces.SaveChannel())
	return r
}
