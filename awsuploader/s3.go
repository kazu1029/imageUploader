package awsuploader

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
)

type S3Uploader struct {
	AccessKey       string
	SecretAccessKey string
	Region          string
	BucketName      string
}

func GenerateUUID() string {
	u := uuid.Must(uuid.NewRandom())
	uu := u.String()
	return uu
}

func (s *S3Uploader) PutToS3(path string, filename string) {
	var contentType, extension string
	file, err := os.Open(fmt.Sprintf("%s%s", path, filename))
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()

	extension = strings.Split(filename, ".")[1]

	switch extension {
	case "jpg":
		contentType = "image/jpeg"
	case "jpeg":
		contentType = "image/jpeg"
	case "gif":
		contentType = "image/gif"
	case "png":
		contentType = "image/png"
	default:
		fmt.Println("this extension is invalid")
	}

	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(s.AccessKey, s.SecretAccessKey, ""),
		Region:      aws.String(s.Region),
	}))

	// FIXME: set filename with uuid
	uid := GenerateUUID
	fmt.Printf("uid is %v\n", uid)
	fmt.Printf("uid type is %v\n", reflect.TypeOf(aws.String(uid)))
	// filename_with_uid := uid + "-" + filename

	uploader := s3manager.NewUploader(sess)
	res, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(s.BucketName),
		Key:         aws.String(filename),
		Body:        file,
		ACL:         aws.String("public-read"),
		ContentType: aws.String(contentType),
	})

	if err != nil {
		fmt.Println(res)
		if err, ok := err.(awserr.Error); ok && err.Code() == request.CanceledErrorCode {
			fmt.Fprintf(os.Stderr, "upload canceled due to timeout %v\n", err)
		} else {
			fmt.Fprintf(os.Stderr, "failed to upload object, %v\n", err)
		}
		os.Exit(1)
	}

	fmt.Printf("successfully uploaded file to %s/%s\n", "retweet-users")
	_ = os.Remove(filename)
}
