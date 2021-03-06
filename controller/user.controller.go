package controller

import (
	"net/http"

	"github.com/dlog/service"

	"github.com/dlog/core"
	"github.com/dlog/dto"
	"github.com/gin-gonic/gin"
)

// UserLogin export
func UserLogIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		var in dto.UserInDTO
		core.GetJSON(c, &in)
		c.JSON(http.StatusOK, gin.H{"user": service.ProcLogin(in)})
	}
}

// UserLogOut export
func UserLogOut() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	}
}

/*
	아래는 구버전


// GetUser export
func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var param core.UserJSON
		if err := c.ShouldBindJSON(&param); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{"user": service.GetUser(param)})
	}
}

// Login export
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var param core.UserInDTO
		if err := c.ShouldBindJSON(&param); err != nil {
			panic(err)
		}

		log.Println(c.Cookie("app_cookie"))
		c.JSON(http.StatusOK, gin.H{"user": service.ProcessLogin(param)})
	}
}

// Logout export
func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		var param core.UserInDTO
		if err := c.ShouldBindJSON(&param); err != nil {
			panic(err)
		}
		service.ProcessLogout(param)
		c.JSON(http.StatusOK, gin.H{"ok": "true"})
	}
}

// VaildRefreshToken export
func VaildRefreshToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var param struct {
			LoginID      string `json:"LoginID"`
			RefreshToken string `json:"RefreshToken"`
		}
		core.GetJSON(c, &param)

		rtn := service.VaildRefreshToken(param)
		if rtn != "" {
			var json = struct {
				LoginID      string
				AccessToken  string
				RefreshToken string
			}{
				param.LoginID,
				rtn,
				param.RefreshToken,
			}
			c.JSON(http.StatusOK, gin.H{"token": json})
		} else {
			c.JSON(http.StatusOK, gin.H{"token": ""})
		}
	}
}
*/
