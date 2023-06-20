package storage

import (
	"context"
	"fmt"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type S3Options struct {
	Endpoint  string
	AccessKey string
	SecretKey string
	Bucket    string
	UseSSL    bool
}

type S3 struct {
	client  *minio.Client
	bucket  string
	options S3Options
}

func NewS3(options S3Options) (Storage, error) {
	minioClient, err := minio.New(options.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(options.AccessKey, options.SecretKey, ""),
		Secure: options.UseSSL,
	})
	if err != nil {
		return nil, err
	}

	return &S3{
		client:  minioClient,
		bucket:  options.Bucket,
		options: options,
	}, nil
}

func (s *S3) Upload(ctx context.Context, filename, mime string, file io.Reader) error {
	_, err := s.client.PutObject(ctx, s.bucket, filename, file, -1, minio.PutObjectOptions{
		ContentType: mime,
	})
	if err != nil {
		return err
	}

	return nil
}

// GetURL returns the URL to the file.
func (s *S3) GetURL(ctx context.Context, filename string) (string, error) {
	protocol := "http://"
	if s.options.UseSSL {
		protocol = "https://"
	}

	endpoint := s.options.Endpoint
	// if endpoint is a docker container, replace it with localhost
	// Because of docker and s3-compatible stores' quirks, we're not going to sign any URLs.
	if endpoint == "minio:9000" {
		endpoint = "localhost:9000"
	}

	return fmt.Sprintf("%s%s/%s/%s", protocol, endpoint, s.bucket, filename), nil

	// for private buckets, most likely proper s3
	//u, err := s.client.PresignedGetObject(ctx, s.bucket, filename, time.Hour*1, nil)
	//if err != nil {
	//	return "", err
	//}
	//
	//return u.String(), err
}
