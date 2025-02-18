// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package s3objectlambda

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"regexp"
)

func init() {
	registry.AddResourceFactory("awscc_s3objectlambda_access_point", accessPointResource)
}

// accessPointResource returns the Terraform awscc_s3objectlambda_access_point resource.
// This Terraform resource corresponds to the CloudFormation AWS::S3ObjectLambda::AccessPoint resource.
func accessPointResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Alias
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "Status": {
		//	      "description": "The status of the Object Lambda alias.",
		//	      "pattern": "^[A-Z]*$",
		//	      "type": "string"
		//	    },
		//	    "Value": {
		//	      "description": "The value of the Object Lambda alias.",
		//	      "pattern": "^[a-z0-9\\-]*$",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"alias": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Status
				"status": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The status of the Object Lambda alias.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Value
				"value": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The value of the Object Lambda alias.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "arn:[^:]+:s3-object-lambda:[^:]*:\\d{12}:accesspoint/.*",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CreationDate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The date and time when the Object lambda Access Point was created.",
		//	  "type": "string"
		//	}
		"creation_date": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The date and time when the Object lambda Access Point was created.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name you want to assign to this Object lambda Access Point.",
		//	  "maxLength": 45,
		//	  "minLength": 3,
		//	  "pattern": "^[a-z0-9]([a-z0-9\\-]*[a-z0-9])?$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name you want to assign to this Object lambda Access Point.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(3, 45),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-z0-9]([a-z0-9\\-]*[a-z0-9])?$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ObjectLambdaConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The Object lambda Access Point Configuration that configures transformations to be applied on the objects on specified S3 Actions",
		//	  "properties": {
		//	    "AllowedFeatures": {
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "CloudWatchMetricsEnabled": {
		//	      "type": "boolean"
		//	    },
		//	    "SupportingAccessPoint": {
		//	      "maxLength": 2048,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    },
		//	    "TransformationConfigurations": {
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "Configuration to define what content transformation will be applied on which S3 Action.",
		//	        "properties": {
		//	          "Actions": {
		//	            "insertionOrder": false,
		//	            "items": {
		//	              "type": "string"
		//	            },
		//	            "type": "array",
		//	            "uniqueItems": true
		//	          },
		//	          "ContentTransformation": {
		//	            "properties": {
		//	              "AwsLambda": {
		//	                "additionalProperties": false,
		//	                "properties": {
		//	                  "FunctionArn": {
		//	                    "maxLength": 2048,
		//	                    "minLength": 1,
		//	                    "type": "string"
		//	                  },
		//	                  "FunctionPayload": {
		//	                    "type": "string"
		//	                  }
		//	                },
		//	                "required": [
		//	                  "FunctionArn"
		//	                ],
		//	                "type": "object"
		//	              }
		//	            },
		//	            "type": "object"
		//	          }
		//	        },
		//	        "required": [
		//	          "Actions",
		//	          "ContentTransformation"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    }
		//	  },
		//	  "required": [
		//	    "SupportingAccessPoint",
		//	    "TransformationConfigurations"
		//	  ],
		//	  "type": "object"
		//	}
		"object_lambda_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AllowedFeatures
				"allowed_features": schema.SetAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
						setplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: CloudWatchMetricsEnabled
				"cloudwatch_metrics_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: SupportingAccessPoint
				"supporting_access_point": schema.StringAttribute{ /*START ATTRIBUTE*/
					Required: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 2048),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
				// Property: TransformationConfigurations
				"transformation_configurations": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Actions
							"actions": schema.SetAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Required:    true,
							}, /*END ATTRIBUTE*/
							// Property: ContentTransformation
							"content_transformation": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: AwsLambda
									"aws_lambda": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: FunctionArn
											"function_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
												Required: true,
												Validators: []validator.String{ /*START VALIDATORS*/
													stringvalidator.LengthBetween(1, 2048),
												}, /*END VALIDATORS*/
											}, /*END ATTRIBUTE*/
											// Property: FunctionPayload
											"function_payload": schema.StringAttribute{ /*START ATTRIBUTE*/
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
								}, /*END SCHEMA*/
								Required: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Required: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The Object lambda Access Point Configuration that configures transformations to be applied on the objects on specified S3 Actions",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: PolicyStatus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "IsPublic": {
		//	      "description": "Specifies whether the Object lambda Access Point Policy is Public or not. Object lambda Access Points are private by default.",
		//	      "type": "boolean"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"policy_status": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: IsPublic
				"is_public": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Specifies whether the Object lambda Access Point Policy is Public or not. Object lambda Access Points are private by default.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PublicAccessBlockConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The PublicAccessBlock configuration that you want to apply to this Access Point. You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status 'The Meaning of Public' in the Amazon Simple Storage Service Developer Guide.",
		//	  "properties": {
		//	    "BlockPublicAcls": {
		//	      "description": "Specifies whether Amazon S3 should block public access control lists (ACLs) to this object lambda access point. Setting this element to TRUE causes the following behavior:\n- PUT Bucket acl and PUT Object acl calls fail if the specified ACL is public.\n - PUT Object calls fail if the request includes a public ACL.\n. - PUT Bucket calls fail if the request includes a public ACL.\nEnabling this setting doesn't affect existing policies or ACLs.",
		//	      "type": "boolean"
		//	    },
		//	    "BlockPublicPolicy": {
		//	      "description": "Specifies whether Amazon S3 should block public bucket policies for buckets in this account. Setting this element to TRUE causes Amazon S3 to reject calls to PUT Bucket policy if the specified bucket policy allows public access. Enabling this setting doesn't affect existing bucket policies.",
		//	      "type": "boolean"
		//	    },
		//	    "IgnorePublicAcls": {
		//	      "description": "Specifies whether Amazon S3 should ignore public ACLs for buckets in this account. Setting this element to TRUE causes Amazon S3 to ignore all public ACLs on buckets in this account and any objects that they contain. Enabling this setting doesn't affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set.",
		//	      "type": "boolean"
		//	    },
		//	    "RestrictPublicBuckets": {
		//	      "description": "Specifies whether Amazon S3 should restrict public bucket policies for this bucket. Setting this element to TRUE restricts access to this bucket to only AWS services and authorized users within this account if the bucket has a public policy.\nEnabling this setting doesn't affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked.",
		//	      "type": "boolean"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"public_access_block_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: BlockPublicAcls
				"block_public_acls": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Specifies whether Amazon S3 should block public access control lists (ACLs) to this object lambda access point. Setting this element to TRUE causes the following behavior:\n- PUT Bucket acl and PUT Object acl calls fail if the specified ACL is public.\n - PUT Object calls fail if the request includes a public ACL.\n. - PUT Bucket calls fail if the request includes a public ACL.\nEnabling this setting doesn't affect existing policies or ACLs.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: BlockPublicPolicy
				"block_public_policy": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Specifies whether Amazon S3 should block public bucket policies for buckets in this account. Setting this element to TRUE causes Amazon S3 to reject calls to PUT Bucket policy if the specified bucket policy allows public access. Enabling this setting doesn't affect existing bucket policies.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: IgnorePublicAcls
				"ignore_public_acls": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Specifies whether Amazon S3 should ignore public ACLs for buckets in this account. Setting this element to TRUE causes Amazon S3 to ignore all public ACLs on buckets in this account and any objects that they contain. Enabling this setting doesn't affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: RestrictPublicBuckets
				"restrict_public_buckets": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Specifies whether Amazon S3 should restrict public bucket policies for this bucket. Setting this element to TRUE restricts access to this bucket to only AWS services and authorized users within this account if the bucket has a public policy.\nEnabling this setting doesn't affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The PublicAccessBlock configuration that you want to apply to this Access Point. You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status 'The Meaning of Public' in the Amazon Simple Storage Service Developer Guide.",
			Computed:    true,
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
		Description: "The AWS::S3ObjectLambda::AccessPoint resource is an Amazon S3ObjectLambda resource type that you can use to add computation to S3 actions",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::S3ObjectLambda::AccessPoint").WithTerraformTypeName("awscc_s3objectlambda_access_point")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"actions":                           "Actions",
		"alias":                             "Alias",
		"allowed_features":                  "AllowedFeatures",
		"arn":                               "Arn",
		"aws_lambda":                        "AwsLambda",
		"block_public_acls":                 "BlockPublicAcls",
		"block_public_policy":               "BlockPublicPolicy",
		"cloudwatch_metrics_enabled":        "CloudWatchMetricsEnabled",
		"content_transformation":            "ContentTransformation",
		"creation_date":                     "CreationDate",
		"function_arn":                      "FunctionArn",
		"function_payload":                  "FunctionPayload",
		"ignore_public_acls":                "IgnorePublicAcls",
		"is_public":                         "IsPublic",
		"name":                              "Name",
		"object_lambda_configuration":       "ObjectLambdaConfiguration",
		"policy_status":                     "PolicyStatus",
		"public_access_block_configuration": "PublicAccessBlockConfiguration",
		"restrict_public_buckets":           "RestrictPublicBuckets",
		"status":                            "Status",
		"supporting_access_point":           "SupportingAccessPoint",
		"transformation_configurations":     "TransformationConfigurations",
		"value":                             "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
