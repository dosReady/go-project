package controller

import (
	"net/http"
	"strconv"

	"github.com/dlog/core"
	"github.com/dlog/service"
	"github.com/gin-gonic/gin"
)

// MngPost export
func MngPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var param core.PostDTO

		if err := c.ShouldBind(&param); err != nil {
			panic(err)
		}

		postID, _ := strconv.ParseInt(param.PostID, 10, 32)
		if postID > 0 {
			service.UpdPost(param)
		} else {
			service.InstPost(param)
		}
		c.JSON(http.StatusOK, gin.H{})
	}
}

// GetPost : Post 상세정보 가져오기
func GetPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var param core.PostDTO

		if err := c.ShouldBind(&param); err != nil {
			panic(err)
		}

		post, tags := service.GetPost(param)
		c.JSON(http.StatusOK, gin.H{"post": post, "tags": tags})
	}
}

// GetPostList : Post 목록 가져오기
func GetPostList() gin.HandlerFunc {
	return func(c *gin.Context) {
		var param core.PostDTO

		if err := c.ShouldBind(&param); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{"list": service.GetPostList(param)})
	}
}
