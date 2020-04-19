package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pinetree/core"
	"github.com/pinetree/service"
)

// MngPost export
func MngPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var info core.INPostInfo
		c.BindJSON(&info)
		if info.PostJSON.PostID > 0 {
			service.UpdPost(info)
		} else {
			service.InstPost(info)
		}
		c.JSON(http.StatusOK, gin.H{})
	}
}

// GetPost export
func GetPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var postInfo core.INPostInfo
		c.BindJSON(&postInfo)

		var result core.OUTPostInfo
		result = service.GetPost(postInfo)
		c.JSON(http.StatusOK, gin.H{"info": result.TbPost, "category": result.TbCategory})
	}
}

// GetPostList export
func GetPostList() gin.HandlerFunc {
	return func(c *gin.Context) {
		var info core.INPostInfo
		c.BindJSON(&info)
		c.JSON(http.StatusOK, gin.H{"list": service.GetPostList(info)})
	}
}
