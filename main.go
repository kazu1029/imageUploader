package main

import (
	"net/http"

	"github.com/kazu1029/imageUploader/awsuploader"
	"github.com/kazu1029/imageUploader/uploader"
	"github.com/kelseyhightower/envconfig"
)

func setupRoutes() {
	http.HandleFunc("/s3/upload", uploadToS3)
	http.HandleFunc("/upload", uploader.UploadFile)
	http.ListenAndServe(":8080", nil)
}

func uploadToS3(w http.ResponseWriter, r *http.Request) {
	var awsconfig awsuploader.AWSConfig
	envconfig.Process("AWS", &awsconfig)

	s3uploader := awsuploader.S3Uploader{
		AccessKey:       awsconfig.AccessKey,
		SecretAccessKey: awsconfig.SecretAccessKey,
		Region:          awsconfig.Region,
		BucketName:      awsconfig.BucketName,
	}
	s3uploader.PutToS3("./temp-images/", "sample1.jpg")
}

func main() {
	setupRoutes()
}
