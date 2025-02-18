// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package dynamodb

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_dynamodb_table", tableResource)
}

// tableResource returns the Terraform awscc_dynamodb_table resource.
// This Terraform resource corresponds to the CloudFormation AWS::DynamoDB::Table resource.
func tableResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AttributeDefinitions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "AttributeName": {
		//	        "type": "string"
		//	      },
		//	      "AttributeType": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "AttributeName",
		//	      "AttributeType"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"attribute_definitions": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: AttributeName
					"attribute_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
					}, /*END ATTRIBUTE*/
					// Property: AttributeType
					"attribute_type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.UniqueValues(),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: BillingMode
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"billing_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ContributorInsightsSpecification
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "Enabled": {
		//	      "type": "boolean"
		//	    }
		//	  },
		//	  "required": [
		//	    "Enabled"
		//	  ],
		//	  "type": "object"
		//	}
		"contributor_insights_specification": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Enabled
				"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Required: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DeletionProtectionEnabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"deletion_protection_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: GlobalSecondaryIndexes
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "ContributorInsightsSpecification": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "Enabled": {
		//	            "type": "boolean"
		//	          }
		//	        },
		//	        "required": [
		//	          "Enabled"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "IndexName": {
		//	        "type": "string"
		//	      },
		//	      "KeySchema": {
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "AttributeName": {
		//	              "type": "string"
		//	            },
		//	            "KeyType": {
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "KeyType",
		//	            "AttributeName"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "type": "array",
		//	        "uniqueItems": true
		//	      },
		//	      "Projection": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "NonKeyAttributes": {
		//	            "items": {
		//	              "type": "string"
		//	            },
		//	            "type": "array",
		//	            "uniqueItems": false
		//	          },
		//	          "ProjectionType": {
		//	            "type": "string"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "ProvisionedThroughput": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "ReadCapacityUnits": {
		//	            "type": "integer"
		//	          },
		//	          "WriteCapacityUnits": {
		//	            "type": "integer"
		//	          }
		//	        },
		//	        "required": [
		//	          "WriteCapacityUnits",
		//	          "ReadCapacityUnits"
		//	        ],
		//	        "type": "object"
		//	      }
		//	    },
		//	    "required": [
		//	      "IndexName",
		//	      "Projection",
		//	      "KeySchema"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"global_secondary_indexes": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: ContributorInsightsSpecification
					"contributor_insights_specification": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Enabled
							"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
								Required: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
							objectplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: IndexName
					"index_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
					}, /*END ATTRIBUTE*/
					// Property: KeySchema
					"key_schema": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: AttributeName
								"attribute_name": schema.StringAttribute{ /*START ATTRIBUTE*/
									Required: true,
								}, /*END ATTRIBUTE*/
								// Property: KeyType
								"key_type": schema.StringAttribute{ /*START ATTRIBUTE*/
									Required: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						Required: true,
						Validators: []validator.List{ /*START VALIDATORS*/
							listvalidator.UniqueValues(),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Projection
					"projection": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: NonKeyAttributes
							"non_key_attributes": schema.ListAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
									listplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: ProjectionType
							"projection_type": schema.StringAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Required: true,
					}, /*END ATTRIBUTE*/
					// Property: ProvisionedThroughput
					"provisioned_throughput": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: ReadCapacityUnits
							"read_capacity_units": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Required: true,
							}, /*END ATTRIBUTE*/
							// Property: WriteCapacityUnits
							"write_capacity_units": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Required: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
							objectplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ImportSourceSpecification
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "InputCompressionType": {
		//	      "type": "string"
		//	    },
		//	    "InputFormat": {
		//	      "type": "string"
		//	    },
		//	    "InputFormatOptions": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "Csv": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Delimiter": {
		//	              "type": "string"
		//	            },
		//	            "HeaderList": {
		//	              "items": {
		//	                "type": "string"
		//	              },
		//	              "type": "array",
		//	              "uniqueItems": true
		//	            }
		//	          },
		//	          "type": "object"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "S3BucketSource": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "S3Bucket": {
		//	          "type": "string"
		//	        },
		//	        "S3BucketOwner": {
		//	          "type": "string"
		//	        },
		//	        "S3KeyPrefix": {
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "S3Bucket"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "required": [
		//	    "S3BucketSource",
		//	    "InputFormat"
		//	  ],
		//	  "type": "object"
		//	}
		"import_source_specification": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: InputCompressionType
				"input_compression_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: InputFormat
				"input_format": schema.StringAttribute{ /*START ATTRIBUTE*/
					Required: true,
				}, /*END ATTRIBUTE*/
				// Property: InputFormatOptions
				"input_format_options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Csv
						"csv": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Delimiter
								"delimiter": schema.StringAttribute{ /*START ATTRIBUTE*/
									Optional: true,
									Computed: true,
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: HeaderList
								"header_list": schema.ListAttribute{ /*START ATTRIBUTE*/
									ElementType: types.StringType,
									Optional:    true,
									Computed:    true,
									Validators: []validator.List{ /*START VALIDATORS*/
										listvalidator.UniqueValues(),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
										listplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: S3BucketSource
				"s3_bucket_source": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: S3Bucket
						"s3_bucket": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
						}, /*END ATTRIBUTE*/
						// Property: S3BucketOwner
						"s3_bucket_owner": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: S3KeyPrefix
						"s3_key_prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Required: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// ImportSourceSpecification is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: KeySchema
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "object"
		//	}
		"key_schema": schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: KinesisStreamSpecification
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "StreamArn": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "StreamArn"
		//	  ],
		//	  "type": "object"
		//	}
		"kinesis_stream_specification": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: StreamArn
				"stream_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Required: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LocalSecondaryIndexes
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "IndexName": {
		//	        "type": "string"
		//	      },
		//	      "KeySchema": {
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "AttributeName": {
		//	              "type": "string"
		//	            },
		//	            "KeyType": {
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "KeyType",
		//	            "AttributeName"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "type": "array",
		//	        "uniqueItems": true
		//	      },
		//	      "Projection": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "NonKeyAttributes": {
		//	            "items": {
		//	              "type": "string"
		//	            },
		//	            "type": "array",
		//	            "uniqueItems": false
		//	          },
		//	          "ProjectionType": {
		//	            "type": "string"
		//	          }
		//	        },
		//	        "type": "object"
		//	      }
		//	    },
		//	    "required": [
		//	      "IndexName",
		//	      "Projection",
		//	      "KeySchema"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"local_secondary_indexes": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: IndexName
					"index_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
					}, /*END ATTRIBUTE*/
					// Property: KeySchema
					"key_schema": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: AttributeName
								"attribute_name": schema.StringAttribute{ /*START ATTRIBUTE*/
									Required: true,
								}, /*END ATTRIBUTE*/
								// Property: KeyType
								"key_type": schema.StringAttribute{ /*START ATTRIBUTE*/
									Required: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						Required: true,
						Validators: []validator.List{ /*START VALIDATORS*/
							listvalidator.UniqueValues(),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Projection
					"projection": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: NonKeyAttributes
							"non_key_attributes": schema.ListAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
									listplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: ProjectionType
							"projection_type": schema.StringAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Required: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PointInTimeRecoverySpecification
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "PointInTimeRecoveryEnabled": {
		//	      "type": "boolean"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"point_in_time_recovery_specification": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: PointInTimeRecoveryEnabled
				"point_in_time_recovery_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ProvisionedThroughput
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "ReadCapacityUnits": {
		//	      "type": "integer"
		//	    },
		//	    "WriteCapacityUnits": {
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "required": [
		//	    "WriteCapacityUnits",
		//	    "ReadCapacityUnits"
		//	  ],
		//	  "type": "object"
		//	}
		"provisioned_throughput": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ReadCapacityUnits
				"read_capacity_units": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Required: true,
				}, /*END ATTRIBUTE*/
				// Property: WriteCapacityUnits
				"write_capacity_units": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Required: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SSESpecification
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "KMSMasterKeyId": {
		//	      "type": "string"
		//	    },
		//	    "SSEEnabled": {
		//	      "type": "boolean"
		//	    },
		//	    "SSEType": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "SSEEnabled"
		//	  ],
		//	  "type": "object"
		//	}
		"sse_specification": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: KMSMasterKeyId
				"kms_master_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: SSEEnabled
				"sse_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Required: true,
				}, /*END ATTRIBUTE*/
				// Property: SSEType
				"sse_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: StreamArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"stream_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: StreamSpecification
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "StreamViewType": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "StreamViewType"
		//	  ],
		//	  "type": "object"
		//	}
		"stream_specification": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: StreamViewType
				"stream_view_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Required: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TableClass
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"table_class": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TableName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"table_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TimeToLiveSpecification
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "AttributeName": {
		//	      "type": "string"
		//	    },
		//	    "Enabled": {
		//	      "type": "boolean"
		//	    }
		//	  },
		//	  "required": [
		//	    "Enabled",
		//	    "AttributeName"
		//	  ],
		//	  "type": "object"
		//	}
		"time_to_live_specification": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AttributeName
				"attribute_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Required: true,
				}, /*END ATTRIBUTE*/
				// Property: Enabled
				"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Required: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Version: None. Resource Type definition for AWS::DynamoDB::Table",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::DynamoDB::Table").WithTerraformTypeName("awscc_dynamodb_table")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                                  "Arn",
		"attribute_definitions":                "AttributeDefinitions",
		"attribute_name":                       "AttributeName",
		"attribute_type":                       "AttributeType",
		"billing_mode":                         "BillingMode",
		"contributor_insights_specification":   "ContributorInsightsSpecification",
		"csv":                                  "Csv",
		"deletion_protection_enabled":          "DeletionProtectionEnabled",
		"delimiter":                            "Delimiter",
		"enabled":                              "Enabled",
		"global_secondary_indexes":             "GlobalSecondaryIndexes",
		"header_list":                          "HeaderList",
		"import_source_specification":          "ImportSourceSpecification",
		"index_name":                           "IndexName",
		"input_compression_type":               "InputCompressionType",
		"input_format":                         "InputFormat",
		"input_format_options":                 "InputFormatOptions",
		"key":                                  "Key",
		"key_schema":                           "KeySchema",
		"key_type":                             "KeyType",
		"kinesis_stream_specification":         "KinesisStreamSpecification",
		"kms_master_key_id":                    "KMSMasterKeyId",
		"local_secondary_indexes":              "LocalSecondaryIndexes",
		"non_key_attributes":                   "NonKeyAttributes",
		"point_in_time_recovery_enabled":       "PointInTimeRecoveryEnabled",
		"point_in_time_recovery_specification": "PointInTimeRecoverySpecification",
		"projection":                           "Projection",
		"projection_type":                      "ProjectionType",
		"provisioned_throughput":               "ProvisionedThroughput",
		"read_capacity_units":                  "ReadCapacityUnits",
		"s3_bucket":                            "S3Bucket",
		"s3_bucket_owner":                      "S3BucketOwner",
		"s3_bucket_source":                     "S3BucketSource",
		"s3_key_prefix":                        "S3KeyPrefix",
		"sse_enabled":                          "SSEEnabled",
		"sse_specification":                    "SSESpecification",
		"sse_type":                             "SSEType",
		"stream_arn":                           "StreamArn",
		"stream_specification":                 "StreamSpecification",
		"stream_view_type":                     "StreamViewType",
		"table_class":                          "TableClass",
		"table_name":                           "TableName",
		"tags":                                 "Tags",
		"time_to_live_specification":           "TimeToLiveSpecification",
		"value":                                "Value",
		"write_capacity_units":                 "WriteCapacityUnits",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/ImportSourceSpecification",
	})
	opts = opts.WithCreateTimeoutInMinutes(720).WithDeleteTimeoutInMinutes(720)

	opts = opts.WithUpdateTimeoutInMinutes(720)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
