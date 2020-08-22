package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dlog/core"

	"github.com/dlog/controller"
	"github.com/dlog/dao"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
)

func initializeDB() {
	log.Println("========================= DB 초기화 ==========================")
	db := dao.Setup()
	tx := db.Begin()

	tx.AutoMigrate(dao.TbUser{}, dao.TbPost{})

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

func checkAuth(c *gin.Context) {
	log.Println("=============== 권한 체크  ================")
	sHeader := c.Request.Header.Get("Authorization")
	sToken := strings.Replace(sHeader, "Bearer", "", 1)
	sDecodeData := core.VaildToken(strings.TrimSpace(sToken))
	log.Println(sDecodeData)
	if len(sDecodeData) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{})
		c.Abort()
		return
	}

	c.Next()
}

func main() {

	initializeDB()

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			log.Println("origin: " + origin)

			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	api := r.Group("")
	{
		api.Use(checkAuth)
		api.POST("/add/post", controller.AddPost())
		api.POST("/remove/post", controller.RemovePost())
		api.POST("/input/post", controller.InputPost())
		api.POST("/echo", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"test": "!!!"})
		})
	}

	r.POST("/get/post", controller.GetPost())
	r.POST("/get/postlist", controller.GetPostList())
	r.POST("/user/login", controller.UserLogIn())
	r.POST("/user/logout", controller.UserLogOut())

	mode := os.Getenv("SERVER_MODE")

	if mode == "oper" {

		m := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist("api.dveloper.me", "dosready.github.io"),
			Cache:      autocert.DirCache("/app"),
		}

		s := &http.Server{
			Addr: ":https",
			TLSConfig: &tls.Config{
				GetCertificate:     m.GetCertificate,
				InsecureSkipVerify: true,
			},
			Handler: r,
		}

		go func() {
			log.Fatal(http.ListenAndServe(":http", m.HTTPHandler(nil)))
		}()

		log.Fatal(s.ListenAndServeTLS("", ""))

	} else {
		if err := r.Run(); err != nil {
			log.Fatal(err)
		}
	}
}
