package controller

import (
	"net/http"

	"github.com/dlog/service"

	"github.com/dlog/core"
	"github.com/dlog/dto"

	"github.com/gin-gonic/gin"
)

//GetPostList export
func GetPostList() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"list": service.GetPostList()})
	}
}

func GetPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	}
}

func AddPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var post dto.PostInDTO
		core.GetJSON(c, &post)
		service.AddPost(post)
		c.JSON(http.StatusOK, gin.H{})
	}
}

func RemovePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	}
}
