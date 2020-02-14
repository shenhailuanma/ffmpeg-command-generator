package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenhailuanma/ffmpeg-command-generator/apiserver/controllers"
)

func Run(listenPort string, env string) error {

	if env == "prd" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(CORS())


	r.GET("/healthz", controllers.HealthzController)

	// apis
	signGroup := r.Group("/api")
	{
		// transcode
		signGroup.POST("/transcode", controllers.CreateTranscodeController)

		signGroup.POST("/snapshot", controllers.CreateSnapshotController)

	}

	return r.Run(listenPort)
}