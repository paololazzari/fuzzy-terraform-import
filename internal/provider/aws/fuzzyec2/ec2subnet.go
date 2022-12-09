package fuzzyec2

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// Get properties for subnets
func EC2SubnetProperties(svc *ec2.EC2) []map[string]interface{} {

	properties := []map[string]interface{}{}

	result, err := svc.DescribeSubnets(nil)
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

		for _, subnet := range result.Subnets {
			obj := make(map[string]interface{})
			obj["Id"] = *subnet.SubnetId
			obj["VpcId"] = *subnet.VpcId
			obj["CidrBlock"] = *subnet.CidrBlock
			obj["AvailabilityZone"] = *subnet.AvailabilityZone
			obj["Tags"] = FormatTags(subnet.Tags)

			properties = append(properties, obj)
		}
	}
	return properties
}
