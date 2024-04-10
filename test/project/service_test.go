package project

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// 假设已经定义了以下模拟存储库
type MockServerRepo struct {
	mock.Mock
}
type MockUserRepo struct {
	mock.Mock
}
type MockLabelRepo struct {
	mock.Mock
}
type MockIdentityRepo struct {
	mock.Mock
}
type MockServerMemberRepo struct {
	mock.Mock
}
type MockMemberRoleRepo struct {
	mock.Mock
}

func TestSaveServer(t *testing.T) {
	// 创建模拟存储库实例
	mockServerRepo := new(MockServerRepo)
	mockUserRepo := new(MockUserRepo)
	mockLabelRepo := new(MockLabelRepo)
	mockIdentityRepo := new(MockIdentityRepo)
	mockServerMemberRepo := new(MockServerMemberRepo)
	mockMemberRoleRepo := new(MockMemberRoleRepo)

	// 定义测试数据
	server := &models.Server{
		Name:   "test_server",
		Labels: []string{"label1", "label2"},
	}
	email := "test@example.com"
	owner := &models.User{
		ID:    "1",
		Email: email,
	}

	// 设置模拟存储库的期望行为
	mockServerRepo.On("ExistServerInNameAndOwner", server.Name, email).Return(false, nil)
	mockUserRepo.On("FindByEmail", email).Return(owner, nil)
	mockLabelRepo.On("ExistByName", mock.Anything).Return(false, nil)
	mockServerRepo.On("SaveServer", mock.Anything, mock.Anything).Return(int64(1), nil)
	mockIdentityRepo.On("SaveIdentity", mock.Anything).Return(int64(1), nil)
	mockServerMemberRepo.On("NewServerMember", mock.Anything, mock.Anything).Return(int64(1), nil)
	mockMemberRoleRepo.On("NewAMemberRole", mock.Anything).Return(int64(1), nil)
	mockUserRepo.On("SaveUser", mock.Anything, mock.Anything, mock.Anything).Return(int64(1), nil)

	s := service.ServerService{}
	// 调用SaveServer方法
	err := s.SaveServer(server, email)

	// 断言
	assert.NoError(t, err, "SaveServer should not return an error")
	mockServerRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
	mockLabelRepo.AssertExpectations(t)
	mockIdentityRepo.AssertExpectations(t)
	mockServerMemberRepo.AssertExpectations(t)
	mockMemberRoleRepo.AssertExpectations(t)
}
