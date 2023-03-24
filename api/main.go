package main

import (
	"net/http"
	"shortener/api/controllers/link"
	"shortener/api/controllers/auth"
	"shortener/api/config"
	"github.com/gin-gonic/gin"
)

func cors(c *gin.Context) {

	c.Header("Access-Control-Allow-Origin", config.EnvVariable("APP_URL"))
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}

func main() {

	//Configure database
	r := gin.Default()

	r.NoRoute(func(c *gin.Context) {
		app := config.EnvVariable("APP_URL") + "/app/"

		c.Redirect(http.StatusMovedPermanently, app)
		c.Next()
	})

	r.Use(cors)
	

	r.GET("/:id", link.GetById)
	r.POST("/link", link.CreateNew)

	r.GET("/auth", auth.Auth)
	r.POST("/auth/login", auth.Login)
	r.POST("/auth/register", auth.Register)
	r.GET("/auth/logout", auth.Logout)

	r.Run()
}
