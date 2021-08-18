// Code generated by generators/resource/main.go; DO NOT EDIT.

package mediaconnect

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
	registry.AddResourceTypeFactory("awscc_mediaconnect_flow", flowResourceType)
}

// flowResourceType returns the Terraform awscc_mediaconnect_flow resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::MediaConnect::Flow resource type.
func flowResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"availability_zone": {
			// Property: AvailabilityZone
			// CloudFormation resource type schema:
			// {
			//   "description": "The Availability Zone that you want to create the flow in. These options are limited to the Availability Zones within the current AWS.",
			//   "type": "string"
			// }
			Description: "The Availability Zone that you want to create the flow in. These options are limited to the Availability Zones within the current AWS.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// AvailabilityZone is a force-new attribute.
		},
		"flow_arn": {
			// Property: FlowArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN), a unique identifier for any AWS resource, of the flow.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN), a unique identifier for any AWS resource, of the flow.",
			Type:        types.StringType,
			Computed:    true,
		},
		"flow_availability_zone": {
			// Property: FlowAvailabilityZone
			// CloudFormation resource type schema:
			// {
			//   "description": "The Availability Zone that you want to create the flow in. These options are limited to the Availability Zones within the current AWS.(ReadOnly)",
			//   "type": "string"
			// }
			Description: "The Availability Zone that you want to create the flow in. These options are limited to the Availability Zones within the current AWS.(ReadOnly)",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the flow.",
			//   "type": "string"
			// }
			Description: "The name of the flow.",
			Type:        types.StringType,
			Required:    true,
			// Name is a force-new attribute.
		},
		"source": {
			// Property: Source
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The settings for the source of the flow.",
			//   "properties": {
			//     "Decryption": {
			//       "additionalProperties": false,
			//       "description": "Information about the encryption of the flow.",
			//       "properties": {
			//         "Algorithm": {
			//           "description": "The type of algorithm that is used for the encryption (such as aes128, aes192, or aes256).",
			//           "enum": [
			//             "aes128",
			//             "aes192",
			//             "aes256"
			//           ],
			//           "type": "string"
			//         },
			//         "ConstantInitializationVector": {
			//           "description": "A 128-bit, 16-byte hex value represented by a 32-character string, to be used with the key for encrypting content. This parameter is not valid for static key encryption.",
			//           "type": "string"
			//         },
			//         "DeviceId": {
			//           "description": "The value of one of the devices that you configured with your digital rights management (DRM) platform key provider. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
			//           "type": "string"
			//         },
			//         "KeyType": {
			//           "description": "The type of key that is used for the encryption. If no keyType is provided, the service will use the default setting (static-key).",
			//           "enum": [
			//             "speke",
			//             "static-key"
			//           ],
			//           "type": "string"
			//         },
			//         "Region": {
			//           "description": "The AWS Region that the API Gateway proxy endpoint was created in. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
			//           "type": "string"
			//         },
			//         "ResourceId": {
			//           "description": "An identifier for the content. The service sends this value to the key server to identify the current endpoint. The resource ID is also known as the content ID. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
			//           "type": "string"
			//         },
			//         "RoleArn": {
			//           "description": "The ARN of the role that you created during setup (when you set up AWS Elemental MediaConnect as a trusted entity).",
			//           "type": "string"
			//         },
			//         "SecretArn": {
			//           "description": " The ARN of the secret that you created in AWS Secrets Manager to store the encryption key. This parameter is required for static key encryption and is not valid for SPEKE encryption.",
			//           "type": "string"
			//         },
			//         "Url": {
			//           "description": "The URL from the API Gateway proxy that you set up to talk to your key server. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "Algorithm",
			//         "RoleArn"
			//       ],
			//       "type": "object"
			//     },
			//     "Description": {
			//       "description": "A description for the source. This value is not used or seen outside of the current AWS Elemental MediaConnect account.",
			//       "type": "string"
			//     },
			//     "EntitlementArn": {
			//       "description": "The ARN of the entitlement that allows you to subscribe to content that comes from another AWS account. The entitlement is set by the content originator and the ARN is generated as part of the originator's flow.",
			//       "type": "string"
			//     },
			//     "IngestIp": {
			//       "description": "The IP address that the flow will be listening on for incoming content.",
			//       "type": "string"
			//     },
			//     "IngestPort": {
			//       "description": "The port that the flow will be listening on for incoming content.",
			//       "type": "integer"
			//     },
			//     "MaxBitrate": {
			//       "description": "The smoothing max bitrate for RIST, RTP, and RTP-FEC streams.",
			//       "type": "integer"
			//     },
			//     "MaxLatency": {
			//       "description": "The maximum latency in milliseconds. This parameter applies only to RIST-based and Zixi-based streams.",
			//       "type": "integer"
			//     },
			//     "Name": {
			//       "description": "The name of the source.",
			//       "type": "string"
			//     },
			//     "Protocol": {
			//       "description": "The protocol that is used by the source or output.",
			//       "enum": [
			//         "zixi-push",
			//         "rtp-fec",
			//         "rtp",
			//         "rist"
			//       ],
			//       "type": "string"
			//     },
			//     "SourceArn": {
			//       "description": "The ARN of the source.",
			//       "type": "string"
			//     },
			//     "StreamId": {
			//       "description": "The stream ID that you want to use for this transport. This parameter applies only to Zixi-based streams.",
			//       "type": "string"
			//     },
			//     "VpcInterfaceName": {
			//       "description": "The name of the VPC Interface this Source is configured with.",
			//       "type": "string"
			//     },
			//     "WhitelistCidr": {
			//       "description": "The range of IP addresses that should be allowed to contribute content to your source. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The settings for the source of the flow.",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"decryption": {
						// Property: Decryption
						Description: "Information about the encryption of the flow.",
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"algorithm": {
									// Property: Algorithm
									Description: "The type of algorithm that is used for the encryption (such as aes128, aes192, or aes256).",
									Type:        types.StringType,
									Required:    true,
								},
								"constant_initialization_vector": {
									// Property: ConstantInitializationVector
									Description: "A 128-bit, 16-byte hex value represented by a 32-character string, to be used with the key for encrypting content. This parameter is not valid for static key encryption.",
									Type:        types.StringType,
									Optional:    true,
								},
								"device_id": {
									// Property: DeviceId
									Description: "The value of one of the devices that you configured with your digital rights management (DRM) platform key provider. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
									Type:        types.StringType,
									Optional:    true,
								},
								"key_type": {
									// Property: KeyType
									Description: "The type of key that is used for the encryption. If no keyType is provided, the service will use the default setting (static-key).",
									Type:        types.StringType,
									Optional:    true,
								},
								"region": {
									// Property: Region
									Description: "The AWS Region that the API Gateway proxy endpoint was created in. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
									Type:        types.StringType,
									Optional:    true,
								},
								"resource_id": {
									// Property: ResourceId
									Description: "An identifier for the content. The service sends this value to the key server to identify the current endpoint. The resource ID is also known as the content ID. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
									Type:        types.StringType,
									Optional:    true,
								},
								"role_arn": {
									// Property: RoleArn
									Description: "The ARN of the role that you created during setup (when you set up AWS Elemental MediaConnect as a trusted entity).",
									Type:        types.StringType,
									Required:    true,
								},
								"secret_arn": {
									// Property: SecretArn
									Description: " The ARN of the secret that you created in AWS Secrets Manager to store the encryption key. This parameter is required for static key encryption and is not valid for SPEKE encryption.",
									Type:        types.StringType,
									Optional:    true,
								},
								"url": {
									// Property: Url
									Description: "The URL from the API Gateway proxy that you set up to talk to your key server. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
									Type:        types.StringType,
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
					"description": {
						// Property: Description
						Description: "A description for the source. This value is not used or seen outside of the current AWS Elemental MediaConnect account.",
						Type:        types.StringType,
						Optional:    true,
					},
					"entitlement_arn": {
						// Property: EntitlementArn
						Description: "The ARN of the entitlement that allows you to subscribe to content that comes from another AWS account. The entitlement is set by the content originator and the ARN is generated as part of the originator's flow.",
						Type:        types.StringType,
						Optional:    true,
					},
					"ingest_ip": {
						// Property: IngestIp
						Description: "The IP address that the flow will be listening on for incoming content.",
						Type:        types.StringType,
						Computed:    true,
					},
					"ingest_port": {
						// Property: IngestPort
						Description: "The port that the flow will be listening on for incoming content.",
						Type:        types.NumberType,
						Optional:    true,
					},
					"max_bitrate": {
						// Property: MaxBitrate
						Description: "The smoothing max bitrate for RIST, RTP, and RTP-FEC streams.",
						Type:        types.NumberType,
						Optional:    true,
					},
					"max_latency": {
						// Property: MaxLatency
						Description: "The maximum latency in milliseconds. This parameter applies only to RIST-based and Zixi-based streams.",
						Type:        types.NumberType,
						Optional:    true,
					},
					"name": {
						// Property: Name
						Description: "The name of the source.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						// Name is a force-new attribute.
					},
					"protocol": {
						// Property: Protocol
						Description: "The protocol that is used by the source or output.",
						Type:        types.StringType,
						Optional:    true,
					},
					"source_arn": {
						// Property: SourceArn
						Description: "The ARN of the source.",
						Type:        types.StringType,
						Computed:    true,
					},
					"stream_id": {
						// Property: StreamId
						Description: "The stream ID that you want to use for this transport. This parameter applies only to Zixi-based streams.",
						Type:        types.StringType,
						Optional:    true,
					},
					"vpc_interface_name": {
						// Property: VpcInterfaceName
						Description: "The name of the VPC Interface this Source is configured with.",
						Type:        types.StringType,
						Optional:    true,
					},
					"whitelist_cidr": {
						// Property: WhitelistCidr
						Description: "The range of IP addresses that should be allowed to contribute content to your source. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Required: true,
		},
		"source_failover_config": {
			// Property: SourceFailoverConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The settings for source failover",
			//   "properties": {
			//     "RecoveryWindow": {
			//       "description": "Search window time to look for dash-7 packets",
			//       "type": "integer"
			//     },
			//     "State": {
			//       "enum": [
			//         "ENABLED",
			//         "DISABLED"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The settings for source failover",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"recovery_window": {
						// Property: RecoveryWindow
						Description: "Search window time to look for dash-7 packets",
						Type:        types.NumberType,
						Optional:    true,
					},
					"state": {
						// Property: State
						Type:     types.StringType,
						Optional: true,
					},
				},
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
		Description: "Resource schema for AWS::MediaConnect::Flow",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::MediaConnect::Flow").WithTerraformTypeName("awscc_mediaconnect_flow").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_mediaconnect_flow", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
