package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/gopheramol/document-handler/configuration"
	"github.com/gopheramol/document-handler/controller"
)

func InitializeRoutes(config configuration.HandlerServiceConfig) *gin.Engine {

	s3HandlerService := InitializeObjects(config)

	router := gin.Default()

	handlerAPI := router.Group("/api/v1")
	{
		handlerAPI.POST("/get-presigned-url", controller.NewController(s3HandlerService.GeneratePresignedURL).HandlerFunc)
		handlerAPI.POST("/scan", controller.NewController(s3HandlerService.SendScanResult).HandlerFunc)
	}
	return router
}
