package main

import (
	"github.com/gin-gonic/gin"
	"gitub.com/golang-gin-poc/controller"
	"gitub.com/golang-gin-poc/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/videos", func(context *gin.Context) {
		context.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(context *gin.Context) {
		context.JSON(200, videoController.Save(context))
	})

	err := server.Run(":3000")
	if err != nil {
		return
	}
}
