package main

import (
	"github.com/gin-gonic/gin"
	"github.com/majdsassi/go-url/controllers"
)

func init() {}
func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")
	router.StaticFile("/go-url.png", "./resources/go-url.png")
	router.POST("/create", controllers.CreateNewURL)
	router.GET("/:id", controllers.GetUrl)
	router.GET("/", controllers.IndexController)
	router.Run(":8080")
}
