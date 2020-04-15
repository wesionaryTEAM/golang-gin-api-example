package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/aashishdahal1/go-api/controller"
	"gitlab.com/aashishdahal1/go-api/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()
	// server.GET("/test", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{"message": "Hello fuckers"})
	// })
	// server.GET("/", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{"message": "Welcome to go"})
	// })

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})
	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})


	server.Run(":8080")
}