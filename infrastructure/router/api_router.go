package router

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/interfaces"
	"github.com/gin-gonic/gin"
	"net/http/pprof"
)

// Router 注册中心
func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/debug/pprof/", gin.WrapF(pprof.Index))
	r.GET("/debug/pprof/heap", gin.WrapF(pprof.Index))
	r.GET("/debug/pprof/goroutine", gin.WrapF(pprof.Index))
	r.GET("/debug/pprof/block", gin.WrapF(pprof.Index))
	r.GET("/debug/pprof/threadcreate", gin.WrapF(pprof.Index))
	// 资源接口
	r.Static("/html", "./web/html")

	// 公共接口
	public := r.Group("/auth")
	public.POST("/login", interfaces.Login())
	public.POST("/register", interfaces.Register())

	// 受保护的接口
	protected := r.Group("/exam")
	protected.Use(TokenAuthMiddleware())
	protected.GET("/demo", TokenAuthMiddleware(), interfaces.Demo())
	protected.GET("/logout", LogoutHandle())

	server := r.Group("/server")
	server.Use(TokenAuthMiddleware())
	server.POST("/save", TokenAuthMiddleware(), interfaces.SaveServer())
	server.GET("/findAllByCurrUser", TokenAuthMiddleware(), interfaces.FindAllServerByUser())
	server.GET("/findJoinByCurrUser", TokenAuthMiddleware(), interfaces.FindJoinServerListByUser())

	channel := r.Group("/channel")
	channel.GET("/findAllByServerId", TokenAuthMiddleware(), interfaces.FindAllChannelByServer())
	channel.POST("/save/:serverID", TokenAuthMiddleware(), interfaces.SaveChannel())

	serverMember := r.Group("/serverMember")
	serverMember.POST("/save", TokenAuthMiddleware(), interfaces.SaveServerMember())
	serverMember.PUT("/edit", TokenAuthMiddleware(), interfaces.EditMemberRoleByMemberId())

	identity := r.Group("/identity")
	identity.POST("/save", TokenAuthMiddleware(), interfaces.SaveIdentity())

	memberRole := r.Group("/memberRole")
	memberRole.POST("/save", TokenAuthMiddleware(), interfaces.SaveMemberRole())

	message := r.Group("/message")
	message.POST("/save", TokenAuthMiddleware(), interfaces.SaveMessage())
	message.GET("/findByKeyword", TokenAuthMiddleware(), interfaces.FindMessageByKeyword())
	message.GET("/findByLimit", TokenAuthMiddleware(), interfaces.FindMessageLimit())
	return r
}
