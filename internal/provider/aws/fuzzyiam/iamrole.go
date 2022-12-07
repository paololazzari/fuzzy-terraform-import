package fuzzyiam

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/iam"
)

// Get properties for IAM roles
func IAMRoleProperties(svc *iam.IAM) []map[string]interface{} {

	properties := []map[string]interface{}{}
	_ = properties

	result, err := svc.ListRoles(nil)
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

		for _, role := range result.Roles {
			obj := make(map[string]interface{})
			obj["CreateDate"] = (*role.CreateDate).String()
			obj["Name"] = *role.RoleName
			obj["Tags"] = FormatTags(*role.RoleName, svc)

			// Not all roles have descriptions
			if role.Description != nil {
				obj["Description"] = *role.Description
			}
			properties = append(properties, obj)

		}
	}
	return properties
}
