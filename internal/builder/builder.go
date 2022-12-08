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

func newIAMRole(sess *session.Session) []map[string]interface{} {
	svc := iam.New(sess)
	return fuzzyiam.IAMRoleProperties(svc)
}

func newEC2Instance(sess *session.Session) []map[string]interface{} {
	svc := ec2.New(sess)
	return fuzzyec2.EC2InstanceProperties(svc)
}

func newS3Bucket(sess *session.Session) []map[string]interface{} {
	svc := s3.New(sess)
	return fuzzys3.S3BucketProperties(svc)
}

func newEC2Subnet(sess *session.Session) []map[string]interface{} {
	svc := ec2.New(sess)
	return fuzzyec2.EC2SubnetProperties(svc)
}

func newEC2Vpc(sess *session.Session) []map[string]interface{} {
	svc := ec2.New(sess)
	return fuzzyec2.EC2VPCProperties(svc)
}

var objectsMap = map[string]interface{}{
	"aws_iam_role": newIAMRole,
	"aws_instance": newEC2Instance,
	"aws_bucket":   newS3Bucket,
	"aws_subnet":   newEC2Instance,
	"aws_vpc":      newEC2Vpc,
}

func GetObjects(resourceName string, sess *session.Session) ([]map[string]interface{}, bool) {
	if objectsMap[resourceName] != nil {
		f := objectsMap[resourceName].(func(sess *session.Session) []map[string]interface{})
		resourceIds := f(sess)
		return resourceIds, true
	} else {
		return nil, false
	}
}
