// Code generated by generators/resource/main.go; DO NOT EDIT.

package cloudformation_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSCloudFormationModuleVersion_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::CloudFormation::ModuleVersion", "awscc_cloudformation_module_version", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}