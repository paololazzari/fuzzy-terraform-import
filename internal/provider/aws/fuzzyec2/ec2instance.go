package fuzzyec2

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// Get properties for EC2 instances
func EC2InstanceProperties(svc *ec2.EC2) []map[string]interface{} {

	properties := []map[string]interface{}{}

	result, err := svc.DescribeInstances(nil)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
	} else {

		for _, reservation := range result.Reservations {
			for _, instance := range reservation.Instances {
				obj := make(map[string]interface{})
				obj["Id"] = *instance.InstanceId
				obj["Type"] = aws.StringValue(instance.InstanceType)
				obj["SubnetId"] = aws.StringValue(instance.SubnetId)
				obj["VpcId"] = aws.StringValue(instance.VpcId)
				obj["Tags"] = FormatTags(instance.Tags)

				properties = append(properties, obj)
			}
		}
	}
	return properties
}
