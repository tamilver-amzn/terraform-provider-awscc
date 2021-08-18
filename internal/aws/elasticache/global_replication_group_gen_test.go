// Code generated by generators/resource/main.go; DO NOT EDIT.

package elasticache_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSElastiCacheGlobalReplicationGroup_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::ElastiCache::GlobalReplicationGroup", "awscc_elasticache_global_replication_group", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
