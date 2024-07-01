package store

import (
	"bytes"
	"encoding/json"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3 struct {
	client   *s3.S3
	S3Region string
	S3Bucket *string
	S3Access string
	S3Secret string
}

func (s *S3) init(cfg Config) error {
	s.client = s3.New(session.Must(session.NewSession()))
	s.S3Region = cfg.S3Region
	s.S3Bucket = aws.String(cfg.S3Bucket)
	s.S3Access = cfg.S3Access
	s.S3Secret = cfg.S3Secret
	return nil
}

func (s *S3) IsExist(filePath string) bool {
	_, err := s.client.HeadObject(&s3.HeadObjectInput{
		Bucket: s.S3Bucket,
		Key:    aws.String(filePath),
	})
	return err == nil
}

func (s *S3) CreateFile(path string, file []byte) error {
	_, err := s.client.PutObject(&s3.PutObjectInput{
		Bucket: s.S3Bucket,
		Key:    aws.String(path),
		Body:   bytes.NewReader(file),
	})

	return err
}

func (s *S3) GetFile(path string) ([]byte, error) {
	if !s.IsExist(path) {
		return nil, nil
	}

	out, err := s.client.GetObject(&s3.GetObjectInput{
		Bucket: s.S3Bucket,
		Key:    aws.String(path),
	})

	if err != nil {
		return nil, err
	}

	return io.ReadAll(out.Body)
}

func (s *S3) RemoveFile(path string) error {
	_, err := s.client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: s.S3Bucket,
		Key:    aws.String(path),
	})

	return err
}

// State can return default value???????
func (s *S3) State(path string) (os.FileInfo, error) {
	_, err := s.client.HeadObject(&s3.HeadObjectInput{
		Bucket: s.S3Bucket,
		Key:    aws.String(path),
	})

	return nil, err
}

func (s *S3) ClearDir(path string) error {
	list, err := s.client.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: s.S3Bucket,
		Prefix: aws.String(path),
	})

	if err != nil {
		return err
	}

	for _, obj := range list.Contents {
		_, err := s.client.DeleteObject(&s3.DeleteObjectInput{
			Bucket: s.S3Bucket,
			Key:    obj.Key,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *S3) MkdirAll(path string) error {
	_, err := s.client.PutObject(&s3.PutObjectInput{
		Bucket: s.S3Bucket,
		Key:    aws.String(path),
		Body:   bytes.NewReader([]byte("")),
	})

	return err
}

func (s *S3) CreateJsonFile(path string, data interface{}) error {
	content, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return s.CreateFile(path, content)
}

func (s *S3) GetJsonFile(path string, file interface{}) error {
	content, err := s.GetFile(path)
	if err != nil {
		return err
	}
	if content == nil {
		return nil
	}
	return json.Unmarshal(content, file)
}
