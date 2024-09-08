package minio_connect

import (
	"context"
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func MinIOConnect() (*minio.Client, error) {
	endpoint := "minio:9000"
	accessKeyID := "user"
	secretAccessKey := "password"
	useSSL := false

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	bucketName := "another-bucket"

	err = minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(context.Background(), bucketName)
		if errBucketExists != nil && exists {
			log.Println(err)
			return nil, err
		}
	}

	policy := fmt.Sprintf(`{
      "Version": "2012-10-17",
      "Statement": [
          {
              "Effect": "Allow",
              "Principal": "*",
              "Action": ["s3:GetObject"],
              "Resource": ["arn:aws:s3:::%s/*"]
          }
      ]
  }`, bucketName)

	err = minioClient.SetBucketPolicy(context.Background(), bucketName, policy)
	if err != nil {
		log.Println("error while setting bucket policy : ", err)
		return nil, err
	}

	return minioClient, err
}
