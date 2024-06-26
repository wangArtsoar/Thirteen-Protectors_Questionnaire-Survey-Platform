package dao

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/constant"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/persistence"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/persistence/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/persistence/repository/server/facade"
	"github.com/go-xorm/xorm"
	"github.com/lib/pq"
)

var _ facade.IServerMemberRepo = new(ServerMemberRepo)

type ServerMemberRepo struct {
}

func NewServerMemberRepo() *ServerMemberRepo {
	return &ServerMemberRepo{}
}

func (s *ServerMemberRepo) FindByUser(userEmail string) ([]models.ServerMember, error) {
	var serverMember []models.ServerMember
	return serverMember, persistence.NewXorm().Cols("id", "member_name", "user_name", "server_id").
		Where("user_email = ?", userEmail).Find(&serverMember)
}

func (s *ServerMemberRepo) NewServerMember(session *xorm.Session, member *models.ServerMember) (int64, error) {
	var lastInsertId int64
	sql := `INSERT INTO server_member(
                        user_id, user_email, user_name, server_id, member_name, identity_id, channels, create_at) 
			VALUES (?,?,?,?,?,?,?,?) RETURNING id`
	if _, err := session.SQL(sql, member.UserId, member.UserEmail, member.UserName, member.ServerId, member.MemberName,
		member.IdentityId, pq.Array([]string{}), member.CreateAt).Get(&lastInsertId); err != nil {
		return 0, err
	}
	return lastInsertId, nil
}

func (s *ServerMemberRepo) FindAllServerMember() ([]*models.ServerMember, error) {
	var serverMembers []*models.ServerMember
	return serverMembers, persistence.NewXorm().Find(serverMembers)
}

func (s *ServerMemberRepo) FindAllServerMemberByServerId(serverId int64) ([]*models.ServerMember, error) {
	var serverMembers []*models.ServerMember
	return serverMembers, persistence.NewXorm().Where("server_id = ?", serverId).Find(serverMembers)
}

func (s *ServerMemberRepo) FindAllServerMemberByChannel(channel *models.Channel) ([]*models.ServerMember, error) {
	var serverMembers []*models.ServerMember
	return serverMembers, persistence.NewXorm().Where("channel = ?", channel.Name).Find(serverMembers)
}

func (s *ServerMemberRepo) FindServerMemberByIdentityId(identityId int64) ([]*models.ServerMember, error) {
	var serverMembers []*models.ServerMember
	return serverMembers, persistence.NewXorm().Where("identity_id = ?", identityId).Find(serverMembers)
}

func (s *ServerMemberRepo) UpdateServerMemberByUserEmail(userEmail string, member *models.ServerMember) (int64, error) {
	return persistence.NewXorm().Where("user_email = ?", userEmail).Update(member)
}

func (s *ServerMemberRepo) DeleteServerMemberById(id int64) (int64, error) {
	return persistence.NewXorm().Id(id).Update("is_delete", constant.Yes)
}
