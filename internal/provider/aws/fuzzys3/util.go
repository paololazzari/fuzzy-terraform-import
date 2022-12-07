package fuzzys3

import (
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// format tags for fuzzyfinder menu
func FormatTags(bucketName string, svc *s3.S3) string {

	tagInput := &s3.GetBucketTaggingInput{
		Bucket: aws.String(bucketName),
	}
	tags, _ := svc.GetBucketTagging(tagInput)

	formattedTags := ""

	// Not all buckets have tags
	if len(tags.TagSet) != 0 {
		t := make(map[string]string)
		keys := []string{}
		for _, tag := range tags.TagSet {
			t[*tag.Key] = *tag.Value
			keys = append(keys, *tag.Key)
		}
		sort.Strings(keys)
		for _, k := range keys {
			formattedTags += k
			formattedTags += ":"
			formattedTags += t[k]
			formattedTags += "\n      "
		}

		formattedTags = strings.TrimRight(formattedTags, " \n")
	}
	return formattedTags
}
