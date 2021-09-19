package db

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"mime/multipart"
	"weekend.side/SocialMedia/config"
)

var Session *session.Session

func ConnectAws() *session.Session {
	AccessKeyID := config.AWS_ACCESS_KEY
	SecretAccessKey := config.AWS_SECRET_ACCESS_KEY
	MyRegion := config.AWS_REGION
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(MyRegion),
		Credentials: credentials.NewStaticCredentials(
			AccessKeyID,
			SecretAccessKey,
			"", // a token will be created when the session it's used.
		),
	})

	if err != nil {
		panic(err)
	}
	return sess
}

func UploadToS3(file multipart.File, filename string) (string, error) {

	uploader := s3manager.NewUploader(Session)
	//upload to the s3 bucket
	up, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(config.AWS_BUCKET),
		ACL:    aws.String("public-read"),
		Key:    aws.String(filename),
		Body:   file,
	})
	if err != nil {
		return "", fmt.Errorf("Failed to upload file %v", err)
	}

	fmt.Print(up)
	return "https://" + config.AWS_BUCKET + "." + "s3.amazonaws.com/" + filename, nil

}
