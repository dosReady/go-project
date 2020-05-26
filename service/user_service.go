package service

import (
	"github.com/dlog/core"
	"github.com/dlog/dao"
)

// GetUser exprot: 사용자를 가져온다.
func GetUser(p core.UserJSON) *core.UserJSON {
	db := dao.Setup()

	var user core.UserJSON

	db.Select(`
		login_id,
		role,
		refresh_token
	`).Table("tb_users").
		Where("login_id = ?", p.LoginID).Scan(&user)

	return &user
}

// ProcessLogin exprot: 로그인 처리를한다.
func ProcessLogin(p core.UserInDTO) *core.UserOutDTO {
	db := dao.Setup()

	var user core.UserOutDTO
	db.Select(`
		login_id,
		role
	`).Table("tb_users").
		Where("login_id = ? and password = ?", p.LoginID, p.Password).Scan(&user)

	if &user != nil {
		// accessToken 발급
		accessToken := core.GenerateToken(p, "access")

		// refreshToken 발급
		refreshToken := core.GenerateToken(p, "refresh")

		user.AccessToken = accessToken
		user.RefreshToken = refreshToken

		db.Table("tb_users").Where("login_id = ?", p.LoginID).
			Updates(core.TbUser{RefreshToken: refreshToken})
	}

	return &user
}

// VaildRefreshToken exprot: 리프레시 토큰 유효성 검사
func VaildRefreshToken(p core.UserInDTO) string {
	refreshToken := core.VaildRefreshToken(p.RefreshToken)

	// 토큰 유효성 검사
	if refreshToken == "" {
		return ""
	}

	// DB 내용과 검사
	if user := GetUser(p.UserJSON); user != nil {
		if user.RefreshToken != p.RefreshToken {
			return ""
		}

		// if user.AccessToken != p.AccessToken {
		// 	return ""
		// }
	}

	// Refresh Token이 유효하면 AccessToken 재발급
	var strgeInfo core.UserJSON
	core.DecodingJSON([]byte(refreshToken), &strgeInfo)
	return core.GenerateToken(strgeInfo, "access")
}
