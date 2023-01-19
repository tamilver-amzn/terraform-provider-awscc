// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ce

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ce_anomaly_subscription", anomalySubscriptionDataSource)
}

// anomalySubscriptionDataSource returns the Terraform awscc_ce_anomaly_subscription data source.
// This Terraform data source corresponds to the CloudFormation AWS::CE::AnomalySubscription resource.
func anomalySubscriptionDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AccountId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The accountId",
		//	  "maxLength": 1024,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"account_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The accountId",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Frequency
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The frequency at which anomaly reports are sent over email. ",
		//	  "enum": [
		//	    "DAILY",
		//	    "IMMEDIATE",
		//	    "WEEKLY"
		//	  ],
		//	  "type": "string"
		//	}
		"frequency": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The frequency at which anomaly reports are sent over email. ",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MonitorArnList
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of cost anomaly monitors.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "description": "Subscription ARN",
		//	    "pattern": "^arn:aws[-a-z0-9]*:[a-z0-9]+:[-a-z0-9]*:[0-9]{12}:[-a-zA-Z0-9/:_]+$",
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"monitor_arn_list": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A list of cost anomaly monitors.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResourceTags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Tags to assign to subscription.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name for the tag.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 200,
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"resource_tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name for the tag.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Tags to assign to subscription.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Subscribers
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of subscriber",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Address": {
		//	        "pattern": "(^[a-zA-Z0-9.!#$%\u0026'*+=?^_‘{|}~-]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$)|(^arn:(aws[a-zA-Z-]*):sns:[a-zA-Z0-9-]+:[0-9]{12}:[a-zA-Z0-9_-]+(\\.fifo)?$)",
		//	        "type": "string"
		//	      },
		//	      "Status": {
		//	        "enum": [
		//	          "CONFIRMED",
		//	          "DECLINED"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "Type": {
		//	        "enum": [
		//	          "EMAIL",
		//	          "SNS"
		//	        ],
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Address",
		//	      "Type"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"subscribers": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Address
					"address": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Status
					"status": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Type
					"type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of subscriber",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SubscriptionArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Subscription ARN",
		//	  "pattern": "^arn:aws[-a-z0-9]*:[a-z0-9]+:[-a-z0-9]*:[0-9]{12}:[-a-zA-Z0-9/:_]+$",
		//	  "type": "string"
		//	}
		"subscription_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Subscription ARN",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SubscriptionName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the subscription.",
		//	  "maxLength": 1024,
		//	  "minLength": 0,
		//	  "pattern": "[\\S\\s]*",
		//	  "type": "string"
		//	}
		"subscription_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the subscription.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Threshold
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The dollar value that triggers a notification if the threshold is exceeded. ",
		//	  "minimum": 0,
		//	  "type": "number"
		//	}
		"threshold": schema.Float64Attribute{ /*START ATTRIBUTE*/
			Description: "The dollar value that triggers a notification if the threshold is exceeded. ",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ThresholdExpression
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An Expression object in JSON String format used to specify the anomalies that you want to generate alerts for.",
		//	  "type": "string"
		//	}
		"threshold_expression": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An Expression object in JSON String format used to specify the anomalies that you want to generate alerts for.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::CE::AnomalySubscription",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CE::AnomalySubscription").WithTerraformTypeName("awscc_ce_anomaly_subscription")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_id":           "AccountId",
		"address":              "Address",
		"frequency":            "Frequency",
		"key":                  "Key",
		"monitor_arn_list":     "MonitorArnList",
		"resource_tags":        "ResourceTags",
		"status":               "Status",
		"subscribers":          "Subscribers",
		"subscription_arn":     "SubscriptionArn",
		"subscription_name":    "SubscriptionName",
		"threshold":            "Threshold",
		"threshold_expression": "ThresholdExpression",
		"type":                 "Type",
		"value":                "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
