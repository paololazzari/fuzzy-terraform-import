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

func newS3Bucket(sess *session.Session) []map[string]interface{} {
	svc := s3.New(sess)
	return fuzzys3.S3BucketProperties(svc)
}

/* The following functions are used to instantiate ec2 objects */

func newEC2EIP(sess *session.Session) []map[string]interface{} {
	svc := ec2.New(sess)
	return fuzzyec2.EC2EIPProperties(svc)
}

func newEC2EIPAssociation(sess *session.Session) []map[string]interface{} {
	svc := ec2.New(sess)
	return fuzzyec2.EC2EIPAssociationProperties(svc)
}

func newEC2Instance(sess *session.Session) []map[string]interface{} {
	svc := ec2.New(sess)
	return fuzzyec2.EC2InstanceProperties(svc)
}

func newEC2KeyPair(sess *session.Session) []map[string]interface{} {
	svc := ec2.New(sess)
	return fuzzyec2.EC2KeysProperties(svc)
}

func newEC2LaunchTemplate(sess *session.Session) []map[string]interface{} {
	svc := ec2.New(sess)
	return fuzzyec2.EC2LaunchTemplateProperties(svc)
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
	"aws_bucket":   newS3Bucket,
	// EC2 function handles
	"aws_eip":             newEC2EIP,
	"aws_eip_association": newEC2EIPAssociation,
	"aws_instance":        newEC2Instance,
	"aws_key_pair":        newEC2KeyPair,
	"aws_launch_template": newEC2LaunchTemplate,
	"aws_subnet":          newEC2Subnet,
	"aws_vpc":             newEC2Vpc,
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
