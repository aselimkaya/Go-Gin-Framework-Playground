package main

import (
	"github.com/aselimkaya/go-gin/middleware/controller"
	"github.com/aselimkaya/go-gin/middleware/middlewares"
	"github.com/aselimkaya/go-gin/middleware/service"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	//Varsayılanı direk kullanmayıp kendimiz yazacağız.
	server := gin.New()
	//gin.Dump() HTTP istek ve cevap gövdesini dump etmeyi sağlıyor. Debug için faydalı.
	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/video", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run(":8080")
}
