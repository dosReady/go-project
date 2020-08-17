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
		var param struct {
			PostCategory string `json:"PostCategory"`
		}
		core.GetJSON(c, &param)

		c.JSON(http.StatusOK, gin.H{"list": service.GetPostList(param.PostCategory)})
	}
}

//GetPost export
func GetPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var param struct {
			PostKey string `json:"postkey"`
		}
		core.GetJSON(c, &param)
		c.JSON(http.StatusOK, gin.H{"post": service.GetPost(param.PostKey)})
	}
}

//AddPost export
func AddPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var post dto.PostInDTO
		core.GetJSON(c, &post)
		service.AddPost(post)
		c.JSON(http.StatusOK, gin.H{})
	}
}

//RemovePost export
func RemovePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	}
}
