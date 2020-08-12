package service

import (
	"github.com/dlog/core"
	"github.com/dlog/dao"
	"github.com/dlog/dto"
)

// ProcLogin export
func ProcLogin(p dto.UserInDTO) (r dto.UserOutDTO) {
	user := dao.GetUser(p.User.LoginID)

	if &user != nil {
		var objType = struct {
			LoginID string
			ROLE    string
		}{
			LoginID: user.LoginID,
			ROLE:    user.Role,
		}
		jsonbyte := core.EncodingJSON(objType)

		// accessToken 발급
		accessToken := core.GenerateToken(jsonbyte)

		r.LoginID = user.LoginID
		r.Role = user.Role
		r.AccessToken = accessToken

		dao.UpdateUserToken(r.LoginID, r.AccessToken)
	}

	return r
}

// ProcLogOut export
func ProcLogOut(loginId string) {
	dao.UpdateUserToken(loginId, " ")
}
