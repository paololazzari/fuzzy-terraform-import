package fuzzyec2

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// Get properties for elastic ips
func EC2EIPProperties(svc *ec2.EC2) []map[string]interface{} {

	properties := []map[string]interface{}{}

	result, err := svc.DescribeAddresses(nil)
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

		for _, address := range result.Addresses {
			obj := make(map[string]interface{})
			obj["Id"] = aws.StringValue(address.AllocationId)

			// Not all eips have an association id
			if address.AssociationId != nil {
				obj["AssociationId"] = aws.StringValue(address.AssociationId)
			}

			// Not all eips have a private ip
			if address.PrivateIpAddress != nil {
				obj["PrivateIp"] = aws.StringValue(address.PrivateIpAddress)
			}

			obj["PublicIp"] = aws.StringValue(address.PublicIp)
			obj["Tags"] = FormatTags(address.Tags)

			properties = append(properties, obj)
		}
	}
	return properties
}
