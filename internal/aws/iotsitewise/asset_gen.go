// Code generated by generators/resource/main.go; DO NOT EDIT.

package iotsitewise

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("awscc_iotsitewise_asset", assetResourceType)
}

// assetResourceType returns the Terraform awscc_iotsitewise_asset resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::IoTSiteWise::Asset resource type.
func assetResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"asset_arn": {
			// Property: AssetArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the asset",
			//   "type": "string"
			// }
			Description: "The ARN of the asset",
			Type:        types.StringType,
			Computed:    true,
		},
		"asset_hierarchies": {
			// Property: AssetHierarchies
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A hierarchy specifies allowed parent/child asset relationships.",
			//     "properties": {
			//       "ChildAssetId": {
			//         "description": "The ID of the child asset to be associated.",
			//         "type": "string"
			//       },
			//       "LogicalId": {
			//         "description": "The LogicalID of a hierarchy in the parent asset's model.",
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "LogicalId",
			//       "ChildAssetId"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			// Multiset.
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"child_asset_id": {
						// Property: ChildAssetId
						Description: "The ID of the child asset to be associated.",
						Type:        types.StringType,
						Required:    true,
					},
					"logical_id": {
						// Property: LogicalId
						Description: "The LogicalID of a hierarchy in the parent asset's model.",
						Type:        types.StringType,
						Required:    true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"asset_id": {
			// Property: AssetId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the asset",
			//   "type": "string"
			// }
			Description: "The ID of the asset",
			Type:        types.StringType,
			Computed:    true,
		},
		"asset_model_id": {
			// Property: AssetModelId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the asset model from which to create the asset.",
			//   "type": "string"
			// }
			Description: "The ID of the asset model from which to create the asset.",
			Type:        types.StringType,
			Required:    true,
		},
		"asset_name": {
			// Property: AssetName
			// CloudFormation resource type schema:
			// {
			//   "description": "A unique, friendly name for the asset.",
			//   "type": "string"
			// }
			Description: "A unique, friendly name for the asset.",
			Type:        types.StringType,
			Required:    true,
		},
		"asset_properties": {
			// Property: AssetProperties
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "The asset property's definition, alias, and notification state.",
			//     "properties": {
			//       "Alias": {
			//         "description": "The property alias that identifies the property.",
			//         "type": "string"
			//       },
			//       "LogicalId": {
			//         "description": "Customer provided ID for property.",
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "NotificationState": {
			//         "description": "The MQTT notification state (ENABLED or DISABLED) for this asset property.",
			//         "enum": [
			//           "ENABLED",
			//           "DISABLED"
			//         ],
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "LogicalId"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			// Multiset.
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"alias": {
						// Property: Alias
						Description: "The property alias that identifies the property.",
						Type:        types.StringType,
						Optional:    true,
					},
					"logical_id": {
						// Property: LogicalId
						Description: "Customer provided ID for property.",
						Type:        types.StringType,
						Required:    true,
					},
					"notification_state": {
						// Property: NotificationState
						Description: "The MQTT notification state (ENABLED or DISABLED) for this asset property.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of key-value pairs that contain metadata for the asset.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "type": "string"
			//       },
			//       "Value": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "A list of key-value pairs that contain metadata for the asset.",
			// Multiset.
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Resource schema for AWS::IoTSiteWise::Asset",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTSiteWise::Asset").WithTerraformTypeName("awscc_iotsitewise_asset").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_iotsitewise_asset", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
