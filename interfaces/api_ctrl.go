package interfaces

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/constant"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/inits"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/interfaces/ass"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/interfaces/ioc"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/interfaces/vo"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Login 登录接口
func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			loginDto      vo.LoginDto
			err           error
			loginResponse *vo.LoginResponse
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
			register vo.RegisterRequest
			err      error
			response *vo.RegisterResponse
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
		server := inits.Server(register.Name)
		err = ioc.Container.ServerService.SaveServer(&server, register.Email)
		if err != nil {
			log.Println("server inits fail")
		}
		log.Println("server inits")
		ctx.JSON(http.StatusOK, response)
	}
}

// SaveServer 创建服务器
func SaveServer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var serverRequest vo.ServerRequest
		err := ctx.ShouldBindJSON(&serverRequest)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errors.New("参数错误"+err.Error()).Error())
			return
		}
		value, exists := ctx.Get(constant.UserName)
		if !exists {
			ctx.JSON(http.StatusBadRequest, errors.New("用户不存在"+err.Error()).Error())
			return
		}
		err = ioc.Container.ServerService.SaveServer(ass.ServerRequestToModel(serverRequest), value.(string))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errors.New("内部错误"+err.Error()).Error())
			return
		}
		ctx.JSON(http.StatusOK, "success")
	}
}

// FindAllServerByUserEmail 获取用户的服务器
func FindAllServerByUserEmail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		value, exists := ctx.Get(constant.UserName)
		if !exists {
			ctx.JSON(http.StatusUnauthorized, errors.New("用户不存在").Error())
			return
		}
		servers, err := ioc.Container.ServerService.FindAllServerByUserEmail(value.(string))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errors.New("内部错误"+err.Error()).Error())
			return
		}
		ctx.JSON(http.StatusOK, servers)
	}
}

// FindAllChannelByServer 获取服务器中的所有频道
func FindAllChannelByServer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		param := ctx.Param("server_id")
		serverId, err := strconv.ParseInt(param, 0, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errors.New("参数错误"+err.Error()).Error())
			return
		}
		channels, err := ioc.Container.ServerService.FindAllChannelByServerId(serverId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errors.New("内部错误"+err.Error()).Error())
			return
		}
		ctx.JSON(http.StatusOK, channels)
	}
}

// SaveChannel 保存频道
func SaveChannel() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			channelRequest vo.ChannelRequest
			serverID       int64
			err            error
		)
		serverID, err = strconv.ParseInt(ctx.Param("serverID"), 0, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errors.New("{serverID}参数错误"+err.Error()).Error())
			return
		}
		err = ctx.ShouldBindJSON(&channelRequest)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errors.New("{channelRequest}参数错误"+err.Error()).Error())
			return
		}
		err = ioc.Container.ServerService.SaveChannel(&models.Channel{
			Name:     channelRequest.Name,
			ServerId: serverID,
			Label:    channelRequest.Label,
			CreateAt: time.Now(),
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errors.New("内部错误"+err.Error()).Error())
			return
		}
		ctx.JSON(http.StatusOK, "")
	}
}

// SaveServerMember 保存服务器人员(加入服务器)
func SaveServerMember() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var serverMemberRequest vo.ServerMemberRequest
		err := ctx.ShouldBindJSON(&serverMemberRequest)
		if err != nil {
			return
		}
		value, exists := ctx.Get(constant.UserName)
		if !exists {
			ctx.JSON(http.StatusBadRequest, errors.New("user not be found"+err.Error()).Error())
			return
		}
		err = ioc.Container.ServerService.SaveServerMember(&models.ServerMember{
			ServerId:   serverMemberRequest.ServerID,
			MemberName: serverMemberRequest.MemberName,
			InviteId:   serverMemberRequest.InviteId,
			UserEmail:  value.(string)})
		if err != nil {
			return
		}
		ctx.JSON(http.StatusOK, "you have in the waiting list")
	}
}

// SaveIdentity 保存身份组
func SaveIdentity() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var identityRequest vo.IdentityRequest
		err := ctx.ShouldBindJSON(&identityRequest)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errors.New("{serverID}参数错误"+err.Error()).Error())
			return
		}
		err = ioc.Container.ServerService.SaveIdentity(ass.IdentityRequestToModel(identityRequest))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errors.New("内部错误"+err.Error()).Error())
			return
		}
		ctx.JSON(http.StatusOK, "success")
	}
}

// SaveMemberRole 保存身份角色
func SaveMemberRole() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var memberRoleRequest vo.MemberRoleRequest
		err := ctx.ShouldBindJSON(&memberRoleRequest)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errors.New("{serverID}参数错误"+err.Error()).Error())
			return
		}
		err = ioc.Container.ServerService.SaveMemberRole(ass.MemberRoleRequestToModel(memberRoleRequest))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errors.New("内部错误"+err.Error()).Error())
			return
		}
		ctx.JSON(http.StatusOK, "success")
	}
}

// SaveMessage 保存信息
func SaveMessage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			messageRequest vo.MessageRequest
			message        *models.Message
		)
		err := ctx.ShouldBindJSON(&messageRequest)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errors.New("{serverID}参数错误"+err.Error()).Error())
			return
		}
		name, exists := ctx.Get(constant.UserName)
		if !exists {
			ctx.JSON(http.StatusBadRequest, errors.New("user not be found"+err.Error()).Error())
			return
		}
		message, err = ioc.Container.ServerService.SaveMessage(ass.MessageRequestToModel(messageRequest), name.(string))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errors.New("内部错误"+err.Error()).Error())
			return
		}
		ctx.JSON(http.StatusOK, ass.MessageModelToResponse(message))
	}
}

// FindMessageByKeyword 通过关键词查询 message
func FindMessageByKeyword() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		keyword := ctx.Param("keyword")
		messageList, err := ioc.Container.ServerService.FindMessageByKeyword(keyword)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errors.New("内部错误"+err.Error()).Error())
			return
		}
		ctx.JSON(http.StatusOK, messageList)
	}
}

// FindMessageLimit 获取对话信息
func FindMessageLimit() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		limit, err := strconv.Atoi(ctx.Param("limit"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errors.New("{serverID}参数错误"+err.Error()).Error())
			return
		}
		messages, err := ioc.Container.ServerService.FindMessageByLimit(limit)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errors.New("内部错误"+err.Error()).Error())
			return
		}
		ctx.JSON(http.StatusOK, ass.MessageModelToResponseList(messages))
	}
}
