package interfaces

import (
	ass2 "Thirteen-Protectors_Questionnaire-Survey-Platform/application/ass"
	vo2 "Thirteen-Protectors_Questionnaire-Survey-Platform/application/vo"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/constant"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/persistence/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/util"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/interfaces/ioc"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
		loginResponse, err = ioc.C.UserService.Login(&loginDto)
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
		response, err = ioc.C.UserService.Register(&register)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errors.New("内部错误"+err.Error()).Error())
			return
		}
		ctx.JSON(http.StatusOK, response)
	}
}

// SaveServer 创建服务器
func SaveServer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var serverRequest vo2.ServerRequest
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
		err = ioc.C.ServerService.SaveServer(ass2.ServerRequestToModel(serverRequest), value.(string))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errors.New("内部错误"+err.Error()).Error())
			return
		}
		ctx.JSON(http.StatusOK, constant.Success)
	}
}

// FindAllServerByUser 获取用户的服务器
func FindAllServerByUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		value, exists := ctx.Get(constant.UserName)
		if !exists {
			ctx.JSON(http.StatusUnauthorized, errors.New("用户不存在").Error())
			return
		}
		servers, err := ioc.C.ServerService.FindAllServerByUserEmail(value.(string))
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
		channels, err := ioc.C.ServerService.FindAllChannelByServerId(serverId)
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
			channelRequest vo2.ChannelRequest
			serverID       int64
			err            error
		)
		if serverID, err = strconv.ParseInt(ctx.Param("serverID"), 0, 64); err != nil {
			ctx.JSON(http.StatusBadRequest, errors.New("{serverID}参数错误"+err.Error()).Error())
			return
		}
		if err = ctx.ShouldBindJSON(&channelRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, errors.New("{channelRequest}参数错误"+err.Error()).Error())
			return
		}
		if err = ioc.C.ServerService.SaveChannel(ass2.ChannelRequestToModel(
			channelRequest, serverID)); err != nil {
			ctx.JSON(http.StatusInternalServerError, errors.New("内部错误"+err.Error()).Error())
			return
		}
		ctx.JSON(http.StatusOK, "")
	}
}

// SaveServerMember 保存服务器人员(加入服务器)
func SaveServerMember() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var serverMemberRequest vo2.ServerMemberRequest
		if err := ctx.ShouldBindJSON(&serverMemberRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, errors.New("参数错误"+err.Error()).Error())
			return
		}
		value, exists := ctx.Get(constant.UserName)
		if !exists {
			ctx.JSON(http.StatusBadRequest, errors.New("user not be found").Error())
			return
		}
		if err := ioc.C.ServerService.SaveServerMember(ass2.ServerMemberRequestToModel(
			serverMemberRequest, value.(string))); err != nil {
			ctx.JSON(http.StatusInternalServerError, errors.New("内部错误"+err.Error()).Error())
			return
		}
		ctx.JSON(http.StatusOK, "you have in the waiting list")
	}
}

// EditMemberRoleByMemberId 修改成员身份角色
func EditMemberRoleByMemberId() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// SaveIdentity 保存身份组
func SaveIdentity() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var identityRequest vo2.IdentityRequest
		if err := ctx.ShouldBindJSON(&identityRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, errors.New("{serverID}参数错误"+err.Error()).Error())
			return
		}
		if err := ioc.C.ServerService.SaveIdentity(ass2.IdentityRequestToModel(identityRequest)); err != nil {
			ctx.JSON(http.StatusInternalServerError, errors.New("内部错误"+err.Error()).Error())
			return
		}
		ctx.JSON(http.StatusOK, constant.Success)
	}
}

// SaveMemberRole 保存身份角色
func SaveMemberRole() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var memberRoleRequest vo2.MemberRoleRequest
		err := ctx.ShouldBindJSON(&memberRoleRequest)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errors.New("{serverID}参数错误"+err.Error()).Error())
			return
		}
		err = ioc.C.ServerService.SaveMemberRole(ass2.MemberRoleRequestToModel(memberRoleRequest))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errors.New("内部错误"+err.Error()).Error())
			return
		}
		ctx.JSON(http.StatusOK, constant.Success)
	}
}

// SaveMessage 保存信息
func SaveMessage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			messageRequest vo2.MessageRequest
			message        *models.Message
		)
		err := ctx.ShouldBindJSON(&messageRequest)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errors.New("{serverID}参数错误"+err.Error()).Error())
			return
		}
		name, exists := ctx.Get(constant.UserName)
		if !exists {
			ctx.JSON(http.StatusBadRequest, errors.New("user not be found").Error())
			return
		}
		message, err = ioc.C.ServerService.SaveMessage(ass2.MessageRequestToModel(messageRequest), name.(string))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errors.New("内部错误"+err.Error()).Error())
			return
		}
		ctx.JSON(http.StatusOK, ass2.MessageModelToResponse(message))
	}
}

// FindMessageByKeyword 通过关键词查询 message
func FindMessageByKeyword() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		keyword := ctx.Param("keyword")
		messageList, err := ioc.C.ServerService.FindMessageByKeyword(keyword)
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
		messages, err := ioc.C.ServerService.FindMessageByLimit(limit)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errors.New("内部错误"+err.Error()).Error())
			return
		}
		ctx.JSON(http.StatusOK, ass2.MessageModelToResponseList(messages))
	}
}

// FindJoinServerListByUser 获取用户加入的服务器列表
func FindJoinServerListByUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pageRequest := util.DefaultPage(ctx.Query("page_num"), ctx.Query("page_size"))
		value, exists := ctx.Get(constant.UserName)
		if !exists {
			ctx.JSON(http.StatusBadRequest, errors.New("user not be found").Error())
			return
		}
		serverList, err := ioc.C.ServerService.FindJoinServerListByUser(value.(string), pageRequest)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errors.New("内部错误"+err.Error()).Error())
			return
		}
		ctx.JSON(http.StatusOK, ass2.PageServerModelToServerResponse(serverList))
	}
}
