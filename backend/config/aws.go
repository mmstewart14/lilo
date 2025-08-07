package config

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// AWSConfig holds AWS configuration
type AWSConfig struct {
	Region          string
	DynamoDBClient  *dynamodb.Client
	S3Client        *s3.Client
}

// InitAWS initializes AWS services
func InitAWS() (*AWSConfig, error) {
	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"), // Default region, can be overridden by environment variables
	)
	if err != nil {
		log.Printf("Error loading AWS config: %v", err)
		return nil, err
	}

	// Create DynamoDB client
	dynamoClient := dynamodb.NewFromConfig(cfg)

	// Create S3 client
	s3Client := s3.NewFromConfig(cfg)

	// Create and return AWS config
	awsConfig := &AWSConfig{
		Region:          cfg.Region,
		DynamoDBClient:  dynamoClient,
		S3Client:        s3Client,
	}

	return awsConfig, nil
}