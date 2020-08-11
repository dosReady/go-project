package dao

import (
	"github.com/dlog/dto"
)

// GetUser export
func GetUser(loginId string) (user dto.UserRsDTO) {
	db := Setup()
	defer db.Close()

	db.Select(`
		login_id,
		role,
		access_token
	`).Table("tb_users").
		Where("login_id = ?", loginId).Scan(&user)
	return user
}

// UpdateUserToken export
func UpdateUserToken(loginId string, token string) {
	db := Setup()
	defer db.Close()

	db.Model(TbUser{}).Where("login_id = ?", loginId).
		Updates(TbUser{AccessToken: token})
}
