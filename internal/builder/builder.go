package builder

import (
	"github.com/paololazzari/fuzzy-terraform-import/internal/provider/aws/fuzzyec2"
	"github.com/paololazzari/fuzzy-terraform-import/internal/provider/aws/fuzzyiam"
	"github.com/paololazzari/fuzzy-terraform-import/internal/provider/aws/fuzzys3"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GetObjects(resourceName string, sess *session.Session) ([]map[string]interface{}, bool) {
	switch resourceName {
	case "aws_iam_role":
		svc := iam.New(sess)
		resourceIds := fuzzyiam.IAMRoleProperties(svc)
		return resourceIds, true
	case "aws_instance":
		svc := ec2.New(sess)
		resourceIds := fuzzyec2.EC2InstanceProperties(svc)
		return resourceIds, true
	case "aws_s3_bucket":
		svc := s3.New(sess)
		resourceIds := fuzzys3.S3BucketProperties(svc)
		return resourceIds, true
	case "aws_subnet":
		svc := ec2.New(sess)
		resourceIds := fuzzyec2.EC2SubnetProperties(svc)
		return resourceIds, true
	case "aws_vpc":
		svc := ec2.New(sess)
		resourceIds := fuzzyec2.EC2VPCProperties(svc)
		return resourceIds, true
	}
	return nil, false
}
