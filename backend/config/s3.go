package config

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// S3 bucket names
const (
	UserImagesS3Bucket     = "lilo-user-images"
	ClothingImagesS3Bucket = "lilo-clothing-images"
	OutfitImagesS3Bucket   = "lilo-outfit-images"
)

// CreateS3Buckets creates all required S3 buckets if they don't exist
func CreateS3Buckets(client *s3.Client, region string) error {
	buckets := []string{
		UserImagesS3Bucket,
		ClothingImagesS3Bucket,
		OutfitImagesS3Bucket,
	}

	for _, bucket := range buckets {
		// Check if bucket exists
		_, err := client.HeadBucket(context.TODO(), &s3.HeadBucketInput{
			Bucket: aws.String(bucket),
		})

		if err != nil {
			// Bucket doesn't exist, create it
			_, err = client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
				Bucket: aws.String(bucket),
				CreateBucketConfiguration: &types.CreateBucketConfiguration{
					LocationConstraint: types.BucketLocationConstraint(region),
				},
			})
			if err != nil {
				log.Printf("Error creating bucket %s: %v", bucket, err)
				return err
			}
			log.Printf("Created bucket %s", bucket)

			// Set bucket policy for public read access to images
			policy := `{
				"Version": "2012-10-17",
				"Statement": [
					{
						"Sid": "PublicReadGetObject",
						"Effect": "Allow",
						"Principal": "*",
						"Action": "s3:GetObject",
						"Resource": "arn:aws:s3:::` + bucket + `/*"
					}
				]
			}`

			_, err = client.PutBucketPolicy(context.TODO(), &s3.PutBucketPolicyInput{
				Bucket: aws.String(bucket),
				Policy: aws.String(policy),
			})
			if err != nil {
				log.Printf("Error setting bucket policy for %s: %v", bucket, err)
				return err
			}
			log.Printf("Set public read policy for bucket %s", bucket)

			// Enable CORS for the bucket
			maxAge := int32(3000)
			corsRule := []types.CORSRule{
				{
					AllowedHeaders: []string{"*"},
					AllowedMethods: []string{"GET", "PUT", "POST", "DELETE", "HEAD"},
					AllowedOrigins: []string{"*"},
					ExposeHeaders:  []string{"ETag"},
					MaxAgeSeconds:  &maxAge,
				},
			}

			_, err = client.PutBucketCors(context.TODO(), &s3.PutBucketCorsInput{
				Bucket: aws.String(bucket),
				CORSConfiguration: &types.CORSConfiguration{
					CORSRules: corsRule,
				},
			})
			if err != nil {
				log.Printf("Error setting CORS for bucket %s: %v", bucket, err)
				return err
			}
			log.Printf("Set CORS for bucket %s", bucket)
		} else {
			log.Printf("Bucket %s already exists", bucket)
		}
	}

	return nil
}
