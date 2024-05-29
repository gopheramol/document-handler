package model

type PreSignedURLRequest struct {
	BucketName string `json:"bucket_name"`
	File       string `json:"file"`
}
