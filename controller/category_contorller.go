package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pinetree/service"
)

// GetCategoryList export
func GetCategoryList() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"list": service.GetCategoryList("")})
	}
}
