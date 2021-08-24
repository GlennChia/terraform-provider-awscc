// Code generated by generators/resource/main.go; DO NOT EDIT.

package route53recoveryreadiness

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
	registry.AddResourceTypeFactory("awscc_route53recoveryreadiness_cell", cellResourceType)
}

// cellResourceType returns the Terraform awscc_route53recoveryreadiness_cell resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Route53RecoveryReadiness::Cell resource type.
func cellResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"cell_arn": {
			// Property: CellArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the cell.",
			//   "maxLength": 256,
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the cell.",
			Type:        types.StringType,
			Computed:    true,
		},
		"cell_name": {
			// Property: CellName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the cell to create.",
			//   "maxLength": 64,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the cell to create.",
			Type:        types.StringType,
			Required:    true,
			// CellName is a force-new attribute.
		},
		"cells": {
			// Property: Cells
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of cell Amazon Resource Names (ARNs) contained within this cell, for use in nested cells. For example, Availability Zones within specific Regions.",
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "maxItems": 5,
			//   "type": "array"
			// }
			Description: "A list of cell Amazon Resource Names (ARNs) contained within this cell, for use in nested cells. For example, Availability Zones within specific Regions.",
			Type:        types.ListType{ElemType: types.StringType},
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenBetween(0, 5),
			},
			Optional: true,
		},
		"parent_readiness_scopes": {
			// Property: ParentReadinessScopes
			// CloudFormation resource type schema:
			// {
			//   "description": "The readiness scope for the cell, which can be a cell Amazon Resource Name (ARN) or a recovery group ARN. This is a list but currently can have only one element.",
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "maxItems": 5,
			//   "type": "array"
			// }
			Description: "The readiness scope for the cell, which can be a cell Amazon Resource Name (ARN) or a recovery group ARN. This is a list but currently can have only one element.",
			Type:        types.ListType{ElemType: types.StringType},
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenBetween(0, 5),
			},
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A collection of tags associated with a resource",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "type": "string"
			//       },
			//       "Value": {
			//         "insertionOrder": false,
			//         "items": {
			//           "maxItems": 50,
			//           "type": "string"
			//         },
			//         "type": "array"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "A collection of tags associated with a resource",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						Type:     types.ListType{ElemType: types.StringType},
						Required: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "The API Schema for AWS Route53 Recovery Readiness Cells.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53RecoveryReadiness::Cell").WithTerraformTypeName("awscc_route53recoveryreadiness_cell")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"cell_arn":                "CellArn",
		"cell_name":               "CellName",
		"cells":                   "Cells",
		"key":                     "Key",
		"parent_readiness_scopes": "ParentReadinessScopes",
		"tags":                    "Tags",
		"value":                   "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_route53recoveryreadiness_cell", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
