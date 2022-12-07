package fuzzyiam

import (
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
)

// format tags for fuzzyfinder menu
func FormatTags(roleName string, svc *iam.IAM) string {

	tagInput := &iam.ListRoleTagsInput{
		RoleName: aws.String(roleName),
	}
	tags, _ := svc.ListRoleTags(tagInput)

	formattedTags := ""

	// Not all roles have tags
	if len(tags.Tags) != 0 {
		t := make(map[string]string)
		keys := []string{}
		for _, tag := range tags.Tags {
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
