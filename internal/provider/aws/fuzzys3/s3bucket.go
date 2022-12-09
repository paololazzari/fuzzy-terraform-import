package fuzzys3

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Get properties for S3 buckets
func S3BucketProperties(svc *s3.S3) []map[string]interface{} {

	properties := []map[string]interface{}{}

	result, err := svc.ListBuckets(nil)
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

		for _, bucket := range result.Buckets {
			obj := make(map[string]interface{})
			obj["CreationDate"] = (bucket.CreationDate).String()
			obj["Name"] = *bucket.Name
			obj["Tags"] = FormatTags(*bucket.Name, svc)

			properties = append(properties, obj)
		}
	}
	return properties
}
