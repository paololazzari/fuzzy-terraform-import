package fuzzyec2_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/paololazzari/fuzzy-terraform-import/internal/provider/aws/fuzzyec2"
	"github.com/stretchr/testify/assert"
)

func TestEC2EipProperties(t *testing.T) {
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
	svc := ec2.New(sess, aws.NewConfig().WithEndpoint(("http://localhost:5000")))
	obj := fuzzyec2.EC2EIPProperties(svc)
	assert := assert.New(t)
	obj_exists := assert.NotEqual(len(obj), 0)
	if obj_exists == false {
		fmt.Printf("No object was found")
		os.Exit(1)
	}
	assert.NotEmpty(obj[0]["Id"], "Id should not be empty")
}
