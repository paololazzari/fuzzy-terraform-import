package fuzzyec2

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// Get properties for placement groups
func EC2PlacementGroupProperties(svc *ec2.EC2) []map[string]interface{} {

	properties := []map[string]interface{}{}

	result, err := svc.DescribePlacementGroups(nil)
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

		for _, placementGroup := range result.PlacementGroups {
			obj := make(map[string]interface{})
			obj["Name"] = aws.StringValue(placementGroup.GroupName)
			obj["Strategy"] = aws.StringValue(placementGroup.Strategy)

			// A placement group may not have a spread level
			if placementGroup.SpreadLevel != nil {
				obj["SpreadLevel"] = aws.StringValue(placementGroup.SpreadLevel)
			}

			// A placement group may not have a partition count
			if placementGroup.PartitionCount != nil {
				obj["PartitionCount"] = strconv.FormatInt(aws.Int64Value(placementGroup.PartitionCount), 10)
			}
			obj["Tags"] = FormatTags(placementGroup.Tags)

			properties = append(properties, obj)
		}
	}
	return properties
}
