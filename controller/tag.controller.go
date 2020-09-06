package controller

import (
	"net/http"

	"github.com/dlog/dto"

	"github.com/dlog/service"
	"github.com/gin-gonic/gin"
)

//Contorller - GetTagList  export
func GetTagList() gin.HandlerFunc {
	return func(c *gin.Context) {
		var param dto.TagDTO
		//core.GetJSON(c, &param)

		c.JSON(http.StatusOK, gin.H{"list": service.GetTagList(param)})
	}
}
