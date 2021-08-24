// Code generated by generators/resource/main.go; DO NOT EDIT.

package location

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"

	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_location_geofence_collection", geofenceCollectionResourceType)
}

// geofenceCollectionResourceType returns the Terraform awscc_location_geofence_collection resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Location::GeofenceCollection resource type.
func geofenceCollectionResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"collection_arn": {
			// Property: CollectionArn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1600,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:       types.StringType,
			Validators: []tfsdk.AttributeValidator{validate.StringLenBetween(0, 1600)},
			Computed:   true,
		},
		"collection_name": {
			// Property: CollectionName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 100,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:       types.StringType,
			Validators: []tfsdk.AttributeValidator{validate.StringLenBetween(1, 100)},
			Required:   true,
			// CollectionName is a force-new attribute.
		},
		"create_time": {
			// Property: CreateTime
			// CloudFormation resource type schema:
			// {
			//   "description": "The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ)",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ)",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1000,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Type:       types.StringType,
			Validators: []tfsdk.AttributeValidator{validate.StringLenBetween(0, 1000)},
			Optional:   true,
			Computed:   true,
			// Description is a force-new attribute.
		},
		"kms_key_id": {
			// Property: KmsKeyId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 2048,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:       types.StringType,
			Validators: []tfsdk.AttributeValidator{validate.StringLenBetween(1, 2048)},
			Optional:   true,
		},
		"pricing_plan": {
			// Property: PricingPlan
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "RequestBasedUsage",
			//     "MobileAssetTracking",
			//     "MobileAssetManagement"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			// PricingPlan is a force-new attribute.
		},
		"pricing_plan_data_source": {
			// Property: PricingPlanDataSource
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			// PricingPlanDataSource is a force-new attribute.
		},
		"update_time": {
			// Property: UpdateTime
			// CloudFormation resource type schema:
			// {
			//   "description": "The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ)",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ)",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Definition of AWS::Location::GeofenceCollection Resource Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Location::GeofenceCollection").WithTerraformTypeName("awscc_location_geofence_collection")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"collection_arn":           "CollectionArn",
		"collection_name":          "CollectionName",
		"create_time":              "CreateTime",
		"description":              "Description",
		"kms_key_id":               "KmsKeyId",
		"pricing_plan":             "PricingPlan",
		"pricing_plan_data_source": "PricingPlanDataSource",
		"update_time":              "UpdateTime",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_location_geofence_collection", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
