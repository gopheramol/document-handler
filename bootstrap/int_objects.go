package bootstrap

import (
	"github.com/gopheramol/document-handler/client"
	"github.com/gopheramol/document-handler/configuration"
	"github.com/gopheramol/document-handler/service"
)

func InitializeObjects(config configuration.HandlerServiceConfig) service.S3HandlerService {

	s3Client := client.NewS3Client(config)

	s3HandlerService := service.NewS3HandlerService(s3Client)

	return s3HandlerService
}
