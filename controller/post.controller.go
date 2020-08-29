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

//GetPost export
func GetPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var param struct {
			PostKey string `json:"postkey"`
		}
		core.GetJSON(c, &param)

		post, tag := service.GetPost(param.PostKey)
		c.JSON(http.StatusOK, gin.H{"post": post, "tag": tag})
	}
}

//RemovePost export
func RemovePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var param struct {
			PostKey string `json:"postkey"`
		}
		core.GetJSON(c, &param)
		service.RemovePost(param.PostKey)
		c.JSON(http.StatusOK, gin.H{})
	}
}

//InputPost export
func InputPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var post dto.PostInDTO
		core.GetJSON(c, &post)
		service.InputPost(post)
		c.JSON(http.StatusOK, gin.H{})
	}
}
