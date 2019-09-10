package awsuploader

type AWSConfig struct {
	AccessKey       string `envconfig:"AWS_ACCESS_KEY_ID"`
	SecretAccessKey string `envconfig:"AWS_SECRET_ACCESS_KEY"`
	Region          string `envconfig:"AWS_DEFAULT_REGION"`
	BucketName      string `envconfig:"AWS_BUCKET_NAME"`
}
