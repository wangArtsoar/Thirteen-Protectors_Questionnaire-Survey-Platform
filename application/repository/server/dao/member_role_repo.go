package dao

import (
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/models"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/application/repository/server/facade"
	"Thirteen-Protectors_Questionnaire-Survey-Platform/infrastructure/orm"
	"github.com/go-xorm/xorm"
	"github.com/lib/pq"
)

var _ facade.IMemberRoleRepo = new(MemberRoleRepo)

type MemberRoleRepo struct {
}

func (m *MemberRoleRepo) NewAMemberRole(session *xorm.Session, role *models.MemberRole) (int64, error) {
	var lastInsertId int64
	sql := `INSERT INTO member_role(server_id, name, permissions) VALUES ($1,$2,$3) RETURNING id`
	if _, err := session.SQL(sql, role.ServerId, role.Name, pq.Array(role.Permissions)).Get(&lastInsertId); err != nil {
		return 0, err
	}
	return lastInsertId, nil
}

func (m *MemberRoleRepo) UpdateMemberRole(memberRoleId int64, role *models.MemberRole) (int64, error) {
	return orm.NewXorm().Id(memberRoleId).Update(role)
}
