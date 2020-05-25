package controller

import (
	"net/http"

	"github.com/dlog/core"
	"github.com/dlog/service"
	"github.com/gin-gonic/gin"
)

// MngPost export
func MngPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var info core.INPostInfo
		if err := c.ShouldBind(&info); err != nil {
			panic(err)
		}

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
		if err := c.ShouldBind(&postInfo); err != nil {
			panic(err)
		}
		result := service.GetPost(postInfo)
		c.JSON(http.StatusOK, gin.H{"info": result.TbPost, "category": result.TbCategory})
	}
}

// GetPostList export
func GetPostList() gin.HandlerFunc {
	return func(c *gin.Context) {
		var info core.INPostInfo
		if err := c.ShouldBind(&info); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{"list": service.GetPostList(info)})
	}
}
