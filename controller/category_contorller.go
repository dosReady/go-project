package controller

import (
	"net/http"

	"github.com/dlog/service"
	"github.com/gin-gonic/gin"
)

// GetCategoryList export
func GetCategoryList() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"list": service.GetCategoryList("")})
	}
}
