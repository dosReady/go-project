package core

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Run export
func Run() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.Query("name"))
	}
}

// Test export
func Test() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.Query("name"))
	}
}
