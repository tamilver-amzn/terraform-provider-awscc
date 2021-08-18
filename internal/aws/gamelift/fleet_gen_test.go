// Code generated by generators/resource/main.go; DO NOT EDIT.

package gamelift_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSGameLiftFleet_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::GameLift::Fleet", "awscc_gamelift_fleet", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyConfig(),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
			),
		},
	})
}

func TestAccAWSGameLiftFleet_disappears(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::GameLift::Fleet", "awscc_gamelift_fleet", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyConfig(),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
				td.DeleteResource(),
			),
			ExpectNonEmptyPlan: true,
		},
	})
}
