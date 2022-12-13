package fuzzyec2

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// Get properties for availability zone groups
func EC2AvailabilityZoneGroupProperties(svc *ec2.EC2) []map[string]interface{} {

	properties := []map[string]interface{}{}

	result, err := svc.DescribeAvailabilityZones(nil)
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

		for _, availabilityZone := range result.AvailabilityZones {
			obj := make(map[string]interface{})
			obj["Name"] = aws.StringValue(availabilityZone.GroupName)
			obj["OptInStatus"] = aws.StringValue(availabilityZone.OptInStatus)
		}
	}
	return properties
}
