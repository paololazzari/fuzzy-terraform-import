package fuzzys3_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/paololazzari/fuzzy-terraform-import/internal/provider/aws/fuzzys3"
	"github.com/stretchr/testify/assert"
)

func TestS3BucketProperties(t *testing.T) {
	region := "us-east-1"
    profile := "local"

	sess, err := session.NewSessionWithOptions(session.Options{
        Config: aws.Config{Region: aws.String(region),
            CredentialsChainVerboseErrors: aws.Bool(true)},
        Profile: profile,
    })
	if err != nil {
        fmt.Println(err)
		os.Exit(1)
    }
	svc := s3.New(sess, aws.NewConfig().WithEndpoint(("http://localhost:5000")))
	bucket := fuzzys3.S3BucketProperties(svc)
	assert := assert.New(t)
	assert.NotEqual(len(bucket),0)
	assert.NotEmpty(bucket[0]["Name"],"")
	assert.NotEmpty(bucket[0]["CreationDate"],"")
}