// Code generated by generators/resource/main.go; DO NOT EDIT.

package resourcegroups

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
	registry.AddResourceTypeFactory("awscc_resourcegroups_group", groupResourceType)
}

// groupResourceType returns the Terraform awscc_resourcegroups_group resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::ResourceGroups::Group resource type.
func groupResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Resource Group ARN.",
			//   "type": "string"
			// }
			Description: "The Resource Group ARN.",
			Type:        types.StringType,
			Computed:    true,
		},
		"configuration": {
			// Property: Configuration
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "properties": {
			//       "Parameters": {
			//         "items": {
			//           "properties": {
			//             "Name": {
			//               "type": "string"
			//             },
			//             "Values": {
			//               "items": {
			//                 "type": "string"
			//               },
			//               "type": "array"
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "type": "array"
			//       },
			//       "Type": {
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"parameters": {
						// Property: Parameters
						Attributes: schema.ListNestedAttributes(
							map[string]schema.Attribute{
								"name": {
									// Property: Name
									Type:     types.StringType,
									Optional: true,
								},
								"values": {
									// Property: Values
									Type:     types.ListType{ElemType: types.StringType},
									Optional: true,
								},
							},
							schema.ListNestedAttributesOptions{},
						),
						Optional: true,
					},
					"type": {
						// Property: Type
						Type:     types.StringType,
						Optional: true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "The description of the resource group",
			//   "maxLength": 512,
			//   "type": "string"
			// }
			Description: "The description of the resource group",
			Type:        types.StringType,
			Optional:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the resource group",
			//   "maxLength": 128,
			//   "type": "string"
			// }
			Description: "The name of the resource group",
			Type:        types.StringType,
			Required:    true,
			// Name is a force-new attribute.
		},
		"resource_query": {
			// Property: ResourceQuery
			// CloudFormation resource type schema:
			// {
			//   "properties": {
			//     "Query": {
			//       "properties": {
			//         "ResourceTypeFilters": {
			//           "items": {
			//             "type": "string"
			//           },
			//           "type": "array"
			//         },
			//         "StackIdentifier": {
			//           "type": "string"
			//         },
			//         "TagFilters": {
			//           "items": {
			//             "properties": {
			//               "Key": {
			//                 "type": "string"
			//               },
			//               "Values": {
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "type": "array"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "Type": {
			//       "enum": [
			//         "TAG_FILTERS_1_0",
			//         "CLOUDFORMATION_STACK_1_0"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"query": {
						// Property: Query
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"resource_type_filters": {
									// Property: ResourceTypeFilters
									Type:     types.ListType{ElemType: types.StringType},
									Optional: true,
								},
								"stack_identifier": {
									// Property: StackIdentifier
									Type:     types.StringType,
									Optional: true,
								},
								"tag_filters": {
									// Property: TagFilters
									Attributes: schema.ListNestedAttributes(
										map[string]schema.Attribute{
											"key": {
												// Property: Key
												Type:     types.StringType,
												Optional: true,
											},
											"values": {
												// Property: Values
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
											},
										},
										schema.ListNestedAttributesOptions{},
									),
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"type": {
						// Property: Type
						Type:     types.StringType,
						Optional: true,
					},
				},
			),
			Optional: true,
		},
		"resources": {
			// Property: Resources
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Optional: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "properties": {
			//       "Key": {
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Optional: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Optional: true,
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
		Description: "Schema for ResourceGroups::Group",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ResourceGroups::Group").WithTerraformTypeName("awscc_resourcegroups_group").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_resourcegroups_group", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
