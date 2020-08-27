package dao

import (
	"github.com/dlog/dto"
)

// GetUser export
func (session *Session) GetUser(loginId string) (user dto.UserRsDTO) {
	db := session.Db
	db.Select(`
		login_id,
		role,
		access_token
	`).Table("tb_users").
		Where("login_id = ?", loginId).Scan(&user)
	return user
}

// UpdateUserToken export
func (session *Session) UpdateUserToken(loginId string, token string) {
	db := session.Db
	db.Model(TbUser{}).Where("login_id = ?", loginId).
		Updates(TbUser{AccessToken: token})
}
