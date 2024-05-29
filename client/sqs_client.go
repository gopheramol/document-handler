package client

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/gopheramol/document-handler/configuration"
)

type SQSClient interface {
	ReceiveMessages(ctx context.Context, queueUrl string) ([]string, error)
	SendMessage(ctx context.Context, queueUrl, messageBody string) (string, error)
}

type sqsClient struct {
	client *sqs.Client
}

func NewSQSClient(cfg configuration.HandlerServiceConfig) (SQSClient, error) {
	awsCfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(cfg.AWSRegion),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.AccessKey, cfg.SecretKey, "")),
	)
	if err != nil {
		return nil, err
	}

	return &sqsClient{
		client: sqs.NewFromConfig(awsCfg),
	}, nil
}

func (c *sqsClient) ReceiveMessages(ctx context.Context, queueUrl string) ([]string, error) {
	input := &sqs.ReceiveMessageInput{
		QueueUrl:            &queueUrl,
		MaxNumberOfMessages: 10, // adjust as needed
		WaitTimeSeconds:     20, // long polling
	}

	resp, err := c.client.ReceiveMessage(ctx, input)
	if err != nil {
		return nil, err
	}

	var messages []string
	for _, msg := range resp.Messages {
		messages = append(messages, *msg.Body)
	}

	return messages, nil
}

func (c *sqsClient) SendMessage(ctx context.Context, queueUrl, messageBody string) (string, error) {
	input := &sqs.SendMessageInput{
		QueueUrl:    &queueUrl,
		MessageBody: &messageBody,
	}

	resp, err := c.client.SendMessage(ctx, input)
	if err != nil {
		return "", err
	}

	return *resp.MessageId, nil
}
