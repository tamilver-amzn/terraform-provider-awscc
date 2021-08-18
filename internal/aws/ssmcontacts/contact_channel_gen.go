// Code generated by generators/resource/main.go; DO NOT EDIT.

package ssmcontacts

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
	registry.AddResourceTypeFactory("awscc_ssmcontacts_contact_channel", contactChannelResourceType)
}

// contactChannelResourceType returns the Terraform awscc_ssmcontacts_contact_channel resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::SSMContacts::ContactChannel resource type.
func contactChannelResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the engagement to a contact channel.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the engagement to a contact channel.",
			Type:        types.StringType,
			Computed:    true,
		},
		"channel_address": {
			// Property: ChannelAddress
			// CloudFormation resource type schema:
			// {
			//   "description": "The details that SSM Incident Manager uses when trying to engage the contact channel.",
			//   "type": "string"
			// }
			Description: "The details that SSM Incident Manager uses when trying to engage the contact channel.",
			Type:        types.StringType,
			Optional:    true,
		},
		"channel_name": {
			// Property: ChannelName
			// CloudFormation resource type schema:
			// {
			//   "description": "The device name. String of 6 to 50 alphabetical, numeric, dash, and underscore characters.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The device name. String of 6 to 50 alphabetical, numeric, dash, and underscore characters.",
			Type:        types.StringType,
			Optional:    true,
		},
		"channel_type": {
			// Property: ChannelType
			// CloudFormation resource type schema:
			// {
			//   "description": "Device type, which specify notification channel. Currently supported values: ?SMS?, ?VOICE?, ?EMAIL?, ?CHATBOT.",
			//   "enum": [
			//     "SMS",
			//     "VOICE",
			//     "EMAIL"
			//   ],
			//   "type": "string"
			// }
			Description: "Device type, which specify notification channel. Currently supported values: ?SMS?, ?VOICE?, ?EMAIL?, ?CHATBOT.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// ChannelType is a force-new attribute.
		},
		"contact_id": {
			// Property: ContactId
			// CloudFormation resource type schema:
			// {
			//   "description": "ARN of the contact resource",
			//   "maxLength": 2048,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "ARN of the contact resource",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// ContactId is a force-new attribute.
		},
		"defer_activation": {
			// Property: DeferActivation
			// CloudFormation resource type schema:
			// {
			//   "description": "If you want to activate the channel at a later time, you can choose to defer activation. SSM Incident Manager can't engage your contact channel until it has been activated.",
			//   "type": "boolean"
			// }
			Description: "If you want to activate the channel at a later time, you can choose to defer activation. SSM Incident Manager can't engage your contact channel until it has been activated.",
			Type:        types.BoolType,
			Optional:    true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Resource Type definition for AWS::SSMContacts::ContactChannel",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SSMContacts::ContactChannel").WithTerraformTypeName("awscc_ssmcontacts_contact_channel").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_ssmcontacts_contact_channel", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
