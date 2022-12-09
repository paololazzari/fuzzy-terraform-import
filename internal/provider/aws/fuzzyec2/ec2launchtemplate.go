package fuzzyec2

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// Get properties for launch templates
func EC2LaunchTemplateProperties(svc *ec2.EC2) []map[string]interface{} {

	properties := []map[string]interface{}{}

	result, err := svc.DescribeLaunchTemplates(nil)
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

		for _, template := range result.LaunchTemplates {
			obj := make(map[string]interface{})
			obj["Id"] = aws.StringValue(template.LaunchTemplateId)
			obj["Name"] = aws.StringValue(template.LaunchTemplateName)
			obj["CreatedBy"] = aws.StringValue(template.CreatedBy)
			obj["DefaultVersionNumber"] = strconv.FormatInt(aws.Int64Value(template.DefaultVersionNumber), 10)
			obj["LatestVersionNumber"] = strconv.FormatInt(aws.Int64Value(template.LatestVersionNumber), 10)
			obj["Tags"] = FormatTags(template.Tags)

			properties = append(properties, obj)
		}
	}
	return properties
}
