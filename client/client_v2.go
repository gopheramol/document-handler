package client

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/gopheramol/document-handler/configuration"
	"github.com/gopheramol/document-handler/model"
)

type S3Client interface {
	GeneratePresignedURL(ctx context.Context, req model.PreSignedURLRequest) (URL string, err error)
	SendScanResult(ctx context.Context, request model.ScanResult) (response string, err error)
}

type s3Client struct {
	config configuration.HandlerServiceConfig
}

func NewS3Client(
	config configuration.HandlerServiceConfig,
) s3Client {
	return s3Client{config: config}
}

// GeneratePresignedURL generates a presigned URL for accessing an S3 object.
func (client s3Client) GeneratePresignedURL(ctx context.Context, req model.PreSignedURLRequest) (URL string, err error) {

	expiry := client.config.Expiry
	regionName := client.config.AWSRegion
	secretKey := client.config.SecretKey
	accessKey := client.config.AccessKey

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
		config.WithRegion(regionName),
	)
	if err != nil {
		log.Fatalf("not able to load aws config %+v", err)
		return
	}

	URL, err = putPresignURL(cfg, req, expiry)
	if err != nil {
		log.Fatalf("not able get pre signed url: %+v", err)
	}
	return
}

func putPresignURL(cfg aws.Config, req model.PreSignedURLRequest, expiry time.Duration) (url string, err error) {
	s3client := s3.NewFromConfig(cfg)
	presignClient := s3.NewPresignClient(s3client)

	presignedUrl, err := presignClient.PresignPutObject(context.Background(),
		&s3.PutObjectInput{
			Bucket: aws.String(req.BucketName),
			Key:    aws.String(req.File),
		},
		s3.WithPresignExpires(expiry))
	if err != nil {
		log.Fatal(err)
		return
	}
	url = presignedUrl.URL
	return

}

// SendScanResult sends the scan result to an SQS queue.
func (client s3Client) SendScanResult(ctx context.Context, request model.ScanResult) (response string, err error) {
	regionName := client.config.AWSRegion
	secretKey := client.config.SecretKey
	accessKey := client.config.AccessKey
	queueURL := client.config.SQSQueueURL

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
		config.WithRegion(regionName),
	)
	if err != nil {
		return "", fmt.Errorf("unable to load SDK config, %v", err)
	}

	sqsClient := sqs.NewFromConfig(cfg)

	messageBody, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("failed to marshal scan result to JSON, %v", err)
	}

	_, err = sqsClient.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:    aws.String(queueURL),
		MessageBody: aws.String(string(messageBody)),
	})
	if err != nil {
		return "", fmt.Errorf("failed to send message to SQS, %v", err)
	}

	return "Message sent to SQS successfully!", nil
}
