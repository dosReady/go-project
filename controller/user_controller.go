package controller

import (
	"net/http"

	"github.com/dlog/core"
	"github.com/dlog/service"
	"github.com/gin-gonic/gin"
)

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

		c.JSON(http.StatusOK, gin.H{"user": service.ProcessLogin(param)})
	}
}

// VaildRefreshToken export
func VaildRefreshToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var param core.UserInDTO
		core.GetJSON(c, &param)

		rtn := service.VaildRefreshToken(param)
		if rtn != "" {
			json := core.UserJSON{
				LoginID:      param.LoginID,
				Role:         param.Role,
				AccessToken:  rtn,
				RefreshToken: param.RefreshToken,
			}
			c.JSON(http.StatusOK, gin.H{"user": json})
		} else {
			c.JSON(http.StatusForbidden, gin.H{})
		}
	}
}
