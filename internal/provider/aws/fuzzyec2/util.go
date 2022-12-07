package fuzzyec2

import (
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go/service/ec2"
)

// format tags for fuzzyfinder menu
func FormatTags(rawTags []*ec2.Tag) string {

	tags := make(map[string]string)
	keys := []string{}
	for _, tag := range rawTags {
		tags[*tag.Key] = *tag.Value
		keys = append(keys, *tag.Key)
	}

	sort.Strings(keys)
	formattedTags := ""
	for _, k := range keys {
		formattedTags += k
		formattedTags += ":"
		formattedTags += tags[k]
		formattedTags += "\n      "
	}
	formattedTags = strings.TrimRight(formattedTags, " \n")
	return formattedTags
}
