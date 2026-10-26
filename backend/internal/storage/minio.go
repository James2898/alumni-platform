package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/James2898/alumni-platform/backend/internal/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type StorageService interface {
	GeneratePresignedUploadURL(ctx context.Context, key, contentType string, lifetime time.Duration) (string, error)
	GetPublicURL(key string) string
}

type minioStorageService struct {
	client       *s3.Client
	presignClient *s3.PresignClient
	cfg          *config.Config
}

func NewMinioStorageService(cfg *config.Config) (StorageService, error) {
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		endpoint := cfg.MinioEndpoint
		scheme := "http"
		if cfg.MinioUseSSL {
			scheme = "https"
		}
		return aws.Endpoint{
			URL:               fmt.Sprintf("%s://%s", scheme, endpoint),
			SigningRegion:     "us-east-1",
			HostnameImmutable: true,
		}, nil
	})

	awsCfg, err := awsconfig.LoadDefaultConfig(context.TODO(),
		awsconfig.WithRegion("us-east-1"),
		awsconfig.WithEndpointResolverWithOptions(customResolver),
		awsconfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			cfg.MinioAccessKey,
			cfg.MinioSecretKey,
			"",
		)),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load s3 config: %w", err)
	}

	client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	presignClient := s3.NewPresignClient(client)

	return &minioStorageService{
		client:        client,
		presignClient: presignClient,
		cfg:           cfg,
	}, nil
}

func (s *minioStorageService) GeneratePresignedUploadURL(ctx context.Context, key, contentType string, lifetime time.Duration) (string, error) {
	req, err := s.presignClient.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.cfg.MinioBucket),
		Key:         aws.String(key),
		ContentType: aws.String(contentType),
	}, s3.WithPresignExpires(lifetime))
	if err != nil {
		return "", fmt.Errorf("failed to presign upload URL: %w", err)
	}

	return req.URL, nil
}

func (s *minioStorageService) GetPublicURL(key string) string {
	if key == "" {
		return ""
	}
	return fmt.Sprintf("%s/%s/%s", s.cfg.MinioPublicURL, s.cfg.MinioBucket, key)
}
