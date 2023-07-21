package router

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/common"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/interfaces"
	"github.com/gin-gonic/gin"
)

// Router 注册中心
func Router() *gin.Engine {
	r := gin.Default()

	// 资源接口
	r.Static("/html", "./web/html")

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
	server.GET("/findAllByCurrUser", common.TokenAuthMiddleware(), interfaces.FindAllServerByUser())
	server.GET("/findJoinByCurrUser", common.TokenAuthMiddleware(), interfaces.FindJoinServerListByUser())

	channel := r.Group("/channel")
	channel.GET("/findAllByServerId", common.TokenAuthMiddleware(), interfaces.FindAllChannelByServer())
	channel.POST("/save/:serverID", common.TokenAuthMiddleware(), interfaces.SaveChannel())

	serverMember := r.Group("/serverMember")
	serverMember.POST("/save", common.TokenAuthMiddleware(), interfaces.SaveServerMember())
	serverMember.PUT("/edit", common.TokenAuthMiddleware(), interfaces.EditMemberRoleByMemberId())

	identity := r.Group("/identity")
	identity.POST("/save", common.TokenAuthMiddleware(), interfaces.SaveIdentity())

	memberRole := r.Group("/memberRole")
	memberRole.POST("/save", common.TokenAuthMiddleware(), interfaces.SaveMemberRole())

	message := r.Group("/message")
	message.POST("/save", common.TokenAuthMiddleware(), interfaces.SaveMessage())
	message.GET("/findByKeyword", common.TokenAuthMiddleware(), interfaces.FindMessageByKeyword())
	message.GET("/findByLimit", common.TokenAuthMiddleware(), interfaces.FindMessageLimit())
	return r
}
