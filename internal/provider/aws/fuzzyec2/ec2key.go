package fuzzyec2

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// Get properties for keys
func EC2KeysProperties(svc *ec2.EC2) []map[string]interface{} {

	properties := []map[string]interface{}{}

	result, err := svc.DescribeKeyPairs(nil)
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

		for _, key := range result.KeyPairs {
			obj := make(map[string]interface{})
			obj["Id"] = aws.StringValue(key.KeyName)
			obj["Tags"] = FormatTags(key.Tags)

			properties = append(properties, obj)
		}
	}
	return properties
}
