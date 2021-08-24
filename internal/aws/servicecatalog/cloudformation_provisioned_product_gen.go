// Code generated by generators/resource/main.go; DO NOT EDIT.

package servicecatalog

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
	registry.AddResourceTypeFactory("awscc_servicecatalog_cloudformation_provisioned_product", cloudFormationProvisionedProductResourceType)
}

// cloudFormationProvisionedProductResourceType returns the Terraform awscc_servicecatalog_cloudformation_provisioned_product resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::ServiceCatalog::CloudFormationProvisionedProduct resource type.
func cloudFormationProvisionedProductResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"accept_language": {
			// Property: AcceptLanguage
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "en",
			//     "jp",
			//     "zh"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
		},
		"cloudformation_stack_arn": {
			// Property: CloudformationStackArn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:       types.StringType,
			Validators: []tfsdk.AttributeValidator{validate.StringLenBetween(1, 256)},
			Computed:   true,
		},
		"notification_arns": {
			// Property: NotificationArns
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "type": "string"
			//   },
			//   "maxItems": 5,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Type: types.ListType{ElemType: types.StringType},
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenBetween(0, 5),
				validate.UniqueItems(),
			},
			Optional: true,
			Computed: true,
			// NotificationArns is a force-new attribute.
		},
		"outputs": {
			// Property: Outputs
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "List of key-value pair outputs.",
			//   "patternProperties": {
			//     "": {
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "List of key-value pair outputs.",
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Computed: true,
		},
		"path_id": {
			// Property: PathId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 100,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:       types.StringType,
			Validators: []tfsdk.AttributeValidator{validate.StringLenBetween(1, 100)},
			Optional:   true,
		},
		"path_name": {
			// Property: PathName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 100,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:       types.StringType,
			Validators: []tfsdk.AttributeValidator{validate.StringLenBetween(1, 100)},
			Optional:   true,
		},
		"product_id": {
			// Property: ProductId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 100,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:       types.StringType,
			Validators: []tfsdk.AttributeValidator{validate.StringLenBetween(1, 100)},
			Optional:   true,
		},
		"product_name": {
			// Property: ProductName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:       types.StringType,
			Validators: []tfsdk.AttributeValidator{validate.StringLenBetween(1, 128)},
			Optional:   true,
		},
		"provisioned_product_id": {
			// Property: ProvisionedProductId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 50,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:       types.StringType,
			Validators: []tfsdk.AttributeValidator{validate.StringLenBetween(1, 50)},
			Computed:   true,
		},
		"provisioned_product_name": {
			// Property: ProvisionedProductName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:       types.StringType,
			Validators: []tfsdk.AttributeValidator{validate.StringLenBetween(1, 128)},
			Optional:   true,
			Computed:   true,
			// ProvisionedProductName is a force-new attribute.
		},
		"provisioning_artifact_id": {
			// Property: ProvisioningArtifactId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 100,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:       types.StringType,
			Validators: []tfsdk.AttributeValidator{validate.StringLenBetween(1, 100)},
			Optional:   true,
		},
		"provisioning_artifact_name": {
			// Property: ProvisioningArtifactName
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
		},
		"provisioning_parameters": {
			// Property: ProvisioningParameters
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 1000,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 4096,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:       types.StringType,
						Validators: []tfsdk.AttributeValidator{validate.StringLenBetween(1, 1000)},
						Required:   true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"provisioning_preferences": {
			// Property: ProvisioningPreferences
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "StackSetAccounts": {
			//       "items": {
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "type": "array",
			//       "uniqueItems": true
			//     },
			//     "StackSetFailureToleranceCount": {
			//       "type": "integer"
			//     },
			//     "StackSetFailureTolerancePercentage": {
			//       "type": "integer"
			//     },
			//     "StackSetMaxConcurrencyCount": {
			//       "type": "integer"
			//     },
			//     "StackSetMaxConcurrencyPercentage": {
			//       "type": "integer"
			//     },
			//     "StackSetOperationType": {
			//       "enum": [
			//         "CREATE",
			//         "UPDATE",
			//         "DELETE"
			//       ],
			//       "type": "string"
			//     },
			//     "StackSetRegions": {
			//       "items": {
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "type": "array",
			//       "uniqueItems": true
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"stack_set_accounts": {
						// Property: StackSetAccounts
						Type: types.ListType{ElemType: types.StringType},
						Validators: []tfsdk.AttributeValidator{
							validate.UniqueItems(),
						},
						Optional: true,
					},
					"stack_set_failure_tolerance_count": {
						// Property: StackSetFailureToleranceCount
						Type:     types.NumberType,
						Optional: true,
					},
					"stack_set_failure_tolerance_percentage": {
						// Property: StackSetFailureTolerancePercentage
						Type:     types.NumberType,
						Optional: true,
					},
					"stack_set_max_concurrency_count": {
						// Property: StackSetMaxConcurrencyCount
						Type:     types.NumberType,
						Optional: true,
					},
					"stack_set_max_concurrency_percentage": {
						// Property: StackSetMaxConcurrencyPercentage
						Type:     types.NumberType,
						Optional: true,
					},
					"stack_set_operation_type": {
						// Property: StackSetOperationType
						Type:     types.StringType,
						Optional: true,
					},
					"stack_set_regions": {
						// Property: StackSetRegions
						Type: types.ListType{ElemType: types.StringType},
						Validators: []tfsdk.AttributeValidator{
							validate.UniqueItems(),
						},
						Optional: true,
					},
				},
			),
			Optional: true,
		},
		"record_id": {
			// Property: RecordId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 50,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:       types.StringType,
			Validators: []tfsdk.AttributeValidator{validate.StringLenBetween(1, 50)},
			Computed:   true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:       types.StringType,
						Validators: []tfsdk.AttributeValidator{validate.StringLenBetween(1, 128)},
						Required:   true,
					},
					"value": {
						// Property: Value
						Type:       types.StringType,
						Validators: []tfsdk.AttributeValidator{validate.StringLenBetween(1, 256)},
						Required:   true,
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
		Description: "Resource Schema for AWS::ServiceCatalog::CloudFormationProvisionedProduct",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ServiceCatalog::CloudFormationProvisionedProduct").WithTerraformTypeName("awscc_servicecatalog_cloudformation_provisioned_product")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"accept_language":                        "AcceptLanguage",
		"cloudformation_stack_arn":               "CloudformationStackArn",
		"key":                                    "Key",
		"notification_arns":                      "NotificationArns",
		"outputs":                                "Outputs",
		"path_id":                                "PathId",
		"path_name":                              "PathName",
		"product_id":                             "ProductId",
		"product_name":                           "ProductName",
		"provisioned_product_id":                 "ProvisionedProductId",
		"provisioned_product_name":               "ProvisionedProductName",
		"provisioning_artifact_id":               "ProvisioningArtifactId",
		"provisioning_artifact_name":             "ProvisioningArtifactName",
		"provisioning_parameters":                "ProvisioningParameters",
		"provisioning_preferences":               "ProvisioningPreferences",
		"record_id":                              "RecordId",
		"stack_set_accounts":                     "StackSetAccounts",
		"stack_set_failure_tolerance_count":      "StackSetFailureToleranceCount",
		"stack_set_failure_tolerance_percentage": "StackSetFailureTolerancePercentage",
		"stack_set_max_concurrency_count":        "StackSetMaxConcurrencyCount",
		"stack_set_max_concurrency_percentage":   "StackSetMaxConcurrencyPercentage",
		"stack_set_operation_type":               "StackSetOperationType",
		"stack_set_regions":                      "StackSetRegions",
		"tags":                                   "Tags",
		"value":                                  "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(720).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(720)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_servicecatalog_cloudformation_provisioned_product", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
