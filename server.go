package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dlog/controller"
	"github.com/dlog/core"
	"github.com/dlog/dao"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
)

func initializeDB() {
	log.Println("========================= DB 초기화 ==========================")
	db := dao.Setup()
	tx := db.Begin()

	tx.AutoMigrate(core.TbPost{}, core.TbUser{}, core.TbTagMst{}, core.TbTagMap{})

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

func vaildateAuth(c *gin.Context) {
	log.Println("=============== 권한 체크  ================")
	reqToken := c.Request.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")
	if len(splitToken) != 2 {
		c.JSON(http.StatusOK, gin.H{"errormsg": "access"})
		c.Abort()
		return
	}

	token := core.VaildAccessToken(strings.TrimSpace(splitToken[1]))
	if token == "" {
		c.JSON(http.StatusOK, gin.H{"errormsg": "access"})
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

	api := r.Group("/api")
	{
		api.Use(vaildateAuth)
		api.POST("/mng/post", controller.MngPost())
		api.POST("/del/post", controller.DelPost())
		api.POST("/get/post", controller.GetPost())
		api.POST("/get/postlist", controller.GetPostList())
		api.POST("/get/taglist", controller.GetTagList())
	}

	r.POST("/proc/login", controller.Login())
	r.POST("/proc/logout", controller.Logout())
	r.POST("/vaild/refresh", controller.VaildRefreshToken())

	mode := os.Getenv("SERVER_MODE")
	if mode != "" {

		m := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist("dveloper.me", "dosready.github.io"),
		}

		s := &http.Server{
			Addr:      ":https",
			TLSConfig: m.TLSConfig(),
			Handler:   r,
		}

		go log.Fatal(http.ListenAndServe(":http", m.HTTPHandler(nil)))

		log.Fatal(s.ListenAndServeTLS("", ""))
	} else {
		if err := r.Run(); err != nil {
			log.Fatal(err)
		}
	}
}
