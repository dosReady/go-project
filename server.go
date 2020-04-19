package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pinetree/controller"
	"github.com/pinetree/core"
	"github.com/pinetree/dao"
)

func initializeDB() {
	log.Println("========================= DB 초기화 ==========================")
	db := dao.Setup()
	tx := db.Begin()

	tx.AutoMigrate(&core.TbPost{}, &core.TbCategory{})

	defer tx.Close()
	defer db.Close()
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

}

func main() {

	initializeDB()

	r := gin.New()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	api := r.Group("/api")
	{
		api.POST("/inst/post", controller.MngPost())
		api.POST("/get/post", controller.GetPost())
		api.POST("/get/postlist", controller.GetPostList())

		api.POST("/get/categorylist", controller.GetCategoryList())
	}
	r.Run()
}
