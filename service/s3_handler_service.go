package service

import (
	"context"

	"github.com/gopheramol/document-handler/client"
	"github.com/gopheramol/document-handler/model"
)

type S3HandlerService interface {
	GeneratePresignedURL(ctx context.Context, request model.PreSignedURLRequest) (response string, err error)
	SendScanResult(ctx context.Context, request model.ScanResult) (response string, err error)
}
type s3HnadlerService struct {
	s3Client client.S3Client
}

func NewS3HandlerService(s3Client client.S3Client) S3HandlerService {
	return s3HnadlerService{
		s3Client: s3Client,
	}
}

func (service s3HnadlerService) GeneratePresignedURL(ctx context.Context, request model.PreSignedURLRequest) (string, error) {

	presignedURL, err := service.s3Client.GeneratePresignedURL(ctx, request)
	if err != nil {
		return "", err
	}
	return presignedURL, nil
}

func (service s3HnadlerService) SendScanResult(ctx context.Context, request model.ScanResult) (response string, err error) {

	presignedURL, err := service.s3Client.SendScanResult(ctx, request)
	if err != nil {
		return "", err
	}
	return presignedURL, nil
}
