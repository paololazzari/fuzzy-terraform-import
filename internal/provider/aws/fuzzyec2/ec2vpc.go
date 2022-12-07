package fuzzyec2

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// Get properties for VPCs
func EC2VPCProperties(svc *ec2.EC2) []map[string]interface{} {

	properties := []map[string]interface{}{}
	_ = properties

	result, err := svc.DescribeVpcs(nil)
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

		for _, vpc := range result.Vpcs {
			obj := make(map[string]interface{})
			obj["Id"] = *vpc.VpcId
			obj["CidrBlock"] = *vpc.CidrBlock
			obj["State"] = *vpc.State
			obj["IsDefault"] = strconv.FormatBool(*vpc.IsDefault)
			obj["Tags"] = FormatTags(vpc.Tags)

			properties = append(properties, obj)
		}
	}
	return properties
}
