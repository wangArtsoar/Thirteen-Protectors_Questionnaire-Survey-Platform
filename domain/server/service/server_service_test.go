package service

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/persistence/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// MockServerService 是 IServerService 接口的Mock实现
type MockServerService struct {
	mock.Mock
}

func (m *MockServerService) SaveServer(server *models.Server, email string) error {
	args := m.Called(server, email)
	return args.Error(0)
}

func (m *MockServerService) FindAllServerByUserEmail(userEmail string) ([]*models.Server, error) {
	args := m.Called(userEmail)
	return args.Get(0).([]*models.Server), args.Error(1)
}

func (m *MockServerService) FindAllChannelByServerId(serverId int64) ([]*models.Channel, error) {
	args := m.Called(serverId)
	return args.Get(0).([]*models.Channel), args.Error(1)
}

func (m *MockServerService) SaveChannel(channel *models.Channel) error {
	args := m.Called(channel)
	return args.Error(0)
}

func (m *MockServerService) SaveServerMember(serverMember *models.ServerMember) error {
	args := m.Called(serverMember)
	return args.Error(0)
}

func (m *MockServerService) SaveIdentity(identity *models.Identity) error {
	args := m.Called(identity)
	return args.Error(0)
}

func (m *MockServerService) SaveMemberRole(role *models.MemberRole) error {
	args := m.Called(role)
	return args.Error(0)
}

func (m *MockServerService) SaveMessage(message *models.Message, userEmail string) (*models.Message, error) {
	args := m.Called(message, userEmail)
	return args.Get(0).(*models.Message), args.Error(1)
}

func (m *MockServerService) FindMessageByKeyword(keyword string) ([]*models.Message, error) {
	args := m.Called(keyword)
	return args.Get(0).([]*models.Message), args.Error(1)
}

func (m *MockServerService) FindMessageByLimit(limit int) ([]*models.Message, error) {
	args := m.Called(limit)
	return args.Get(0).([]*models.Message), args.Error(1)
}

func (m *MockServerService) FindJoinServerListByUser(userEmail string, request util.PageRequest) (*util.PageList[models.Server], error) {
	args := m.Called(userEmail, request)
	return args.Get(0).(*util.PageList[models.Server]), args.Error(1)
}

func TestSaveServer(t *testing.T) {
	// 创建Mock对象
	mockService := new(MockServerService)

	// 设置预期调用和返回值
	//expectedServer := &models.Server{}
	expectedEmail := "xiaoyi@icloud.com"
	fakeServers := []*models.Server{
		{Id: 35, Name: "xiaoyi_second"},
		{Id: 44, Name: "xiaoyi_first"},
	}
	mockService.On("FindAllServerByUserEmail", expectedEmail).Return(fakeServers, nil)

	// 在测试中使用Mock对象
	// 例如，在你的业务代码中调用了SaveServer方法：
	servers, err := MyBusinessLogicUsingFindAllServerByUserEmail(mockService)
	if err != nil {
		t.Error(err)
	}
	// 断言期望的调用是否发生
	mockService.AssertExpectations(t)

	//// 断言业务逻辑的结果
	assert.Equal(t, fakeServers, servers, "Returned servers should match the expected servers")
}

// 在你的业务逻辑中调用SaveServer方法
func MyBusinessLogicUsingFindAllServerByUserEmail(service IServerService) ([]*models.Server, error) {
	return service.FindAllServerByUserEmail("xiaoyi@icloud.com")
}
