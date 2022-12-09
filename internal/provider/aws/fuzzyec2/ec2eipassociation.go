package fuzzyec2

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// Get properties for elastic ips
func EC2EIPAssociationProperties(svc *ec2.EC2) []map[string]interface{} {

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

			// An EIP is not necessarily associated with a resource
			if address.AssociationId != nil {
				obj["Id"] = aws.StringValue(address.AssociationId)
			}

			// An EIP is not necessarily associated with an instance
			if address.InstanceId != nil {
				obj["InstanceId"] = aws.StringValue(address.InstanceId)
			}

			// If the object does not have an ID, then it doesn't exist
			if _, ok := obj["Id"]; ok {
				properties = append(properties, obj)
			}

		}
	}
	return properties
}
