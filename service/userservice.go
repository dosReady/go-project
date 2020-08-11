package service

import (
	"github.com/dlog/core"
	"github.com/dlog/dao"
)

// GetUser exprot: 사용자를 가져온다.
func GetUser(p core.UserJSON) *core.UserJSON {
	db := dao.Setup()
	defer db.Close()
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
	defer db.Close()

	var user core.UserOutDTO
	db.Select(`
		login_id,
		role
	`).Table("tb_users").
		Where("login_id = ? and password = ?", p.LoginID, p.Password).Scan(&user)

	if &user != nil {
		type pType struct {
			LoginID string
			ROLE    string
		}

		// accessToken 발급
		accessToken := core.GenerateToken(pType{}, "access")

		// refreshToken 발급
		param := pType{
			LoginID: user.LoginID,
			ROLE:    user.Role,
		}
		refreshToken := core.GenerateToken(param, "refresh")

		user.AccessToken = accessToken
		user.RefreshToken = refreshToken

		db.Table("tb_users").Where("login_id = ?", p.LoginID).
			Updates(core.TbUser{RefreshToken: refreshToken})
	}

	return &user
}

// ProcessLogout export: 로그아웃 처리를 한다.
func ProcessLogout(p core.UserInDTO) {
	db := dao.Setup()
	defer db.Close()
	db.Table("tb_users").Where("login_id = ?", p.LoginID).
		Updates(core.TbUser{RefreshToken: " "})
}

// VaildRefreshToken exprot: 리프레시 토큰 유효성 검사
func VaildRefreshToken(p struct {
	LoginID      string `json:"LoginID"`
	RefreshToken string `json:"RefreshToken"`
}) string {
	refreshToken := core.VaildRefreshToken(p.RefreshToken)

	// 토큰 유효성 검사
	if refreshToken == "" {
		return ""
	}

	var userParam = core.UserJSON{
		LoginID: p.LoginID,
	}

	// DB 내용과 검사
	if user := GetUser(userParam); user != nil {
		if user.RefreshToken != p.RefreshToken {
			return ""
		}
	}

	// Refresh Token이 유효하면 AccessToken 재발급
	var param struct {
		LoginID string
		ROLE    string
	}
	core.DecodingJSON([]byte(refreshToken), &param)
	return core.GenerateToken(param, "access")
}
