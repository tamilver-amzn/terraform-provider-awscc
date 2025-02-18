// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package glue

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_glue_schema_version_metadata", schemaVersionMetadataDataSource)
}

// schemaVersionMetadataDataSource returns the Terraform awscc_glue_schema_version_metadata data source.
// This Terraform data source corresponds to the CloudFormation AWS::Glue::SchemaVersionMetadata resource.
func schemaVersionMetadataDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Key
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Metadata key",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"key": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Metadata key",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SchemaVersionId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Represents the version ID associated with the schema version.",
		//	  "pattern": "[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}",
		//	  "type": "string"
		//	}
		"schema_version_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Represents the version ID associated with the schema version.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Value
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Metadata value",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"value": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Metadata value",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Glue::SchemaVersionMetadata",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Glue::SchemaVersionMetadata").WithTerraformTypeName("awscc_glue_schema_version_metadata")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"key":               "Key",
		"schema_version_id": "SchemaVersionId",
		"value":             "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
