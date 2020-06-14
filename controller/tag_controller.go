package controller

import (
	"net/http"

	"github.com/dlog/service"
	"github.com/gin-gonic/gin"
)

// GetTagList :
func GetTagList() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"list": service.GetTagList()})
	}
}
