package dao

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/repository/server/facade"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/orm"
	"github.com/lib/pq"
)

var _ facade.IMemberRoleRepo = new(MemberRoleRepo)

type MemberRoleRepo struct {
}

func (m *MemberRoleRepo) NewAMemberRole(role *models.MemberRole) (int64, error) {
	var lastInsertId int64
	sql := `INSERT INTO member_role(channel_id, name, permissions) VALUES ($1,$2,$3) RETURNING id`
	err := orm.NewXorm().DB().QueryRow(sql, role.ChannelId, role.Name, pq.Array(role.Permissions)).Scan(lastInsertId)
	if err != nil {
		return 0, err
	}
	return lastInsertId, nil
}

func (m *MemberRoleRepo) UpdateMemberRole(memberRoleId int64, role *models.MemberRole) (int64, error) {
	return orm.NewXorm().Id(memberRoleId).Update(role)
}
