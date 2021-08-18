// Code generated by generators/resource/main.go; DO NOT EDIT.

package detective_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSDetectiveMemberInvitation_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Detective::MemberInvitation", "awscc_detective_member_invitation", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
