// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package forecast

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_forecast_dataset", datasetDataSourceType)
}

// datasetDataSourceType returns the Terraform awscc_forecast_dataset data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Forecast::Dataset resource type.
func datasetDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 256,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"data_frequency": {
			// Property: DataFrequency
			// CloudFormation resource type schema:
			// {
			//   "description": "Frequency of data collection. This parameter is required for RELATED_TIME_SERIES",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Frequency of data collection. This parameter is required for RELATED_TIME_SERIES",
			Type:        types.StringType,
			Computed:    true,
		},
		"dataset_name": {
			// Property: DatasetName
			// CloudFormation resource type schema:
			// {
			//   "description": "A name for the dataset",
			//   "maxLength": 63,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "A name for the dataset",
			Type:        types.StringType,
			Computed:    true,
		},
		"dataset_type": {
			// Property: DatasetType
			// CloudFormation resource type schema:
			// {
			//   "description": "The dataset type",
			//   "enum": [
			//     "TARGET_TIME_SERIES",
			//     "RELATED_TIME_SERIES",
			//     "ITEM_METADATA"
			//   ],
			//   "type": "string"
			// }
			Description: "The dataset type",
			Type:        types.StringType,
			Computed:    true,
		},
		"domain": {
			// Property: Domain
			// CloudFormation resource type schema:
			// {
			//   "description": "The domain associated with the dataset",
			//   "enum": [
			//     "RETAIL",
			//     "CUSTOM",
			//     "INVENTORY_PLANNING",
			//     "EC2_CAPACITY",
			//     "WORK_FORCE",
			//     "WEB_TRAFFIC",
			//     "METRICS"
			//   ],
			//   "type": "string"
			// }
			Description: "The domain associated with the dataset",
			Type:        types.StringType,
			Computed:    true,
		},
		"encryption_config": {
			// Property: EncryptionConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "KmsKeyArn": {
			//       "description": "KMS key used to encrypt the Dataset data",
			//       "maxLength": 256,
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "RoleArn": {
			//       "description": "The ARN of the IAM role that Amazon Forecast can assume to access the AWS KMS key.",
			//       "maxLength": 256,
			//       "pattern": "",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"kms_key_arn": {
						// Property: KmsKeyArn
						Description: "KMS key used to encrypt the Dataset data",
						Type:        types.StringType,
						Computed:    true,
					},
					"role_arn": {
						// Property: RoleArn
						Description: "The ARN of the IAM role that Amazon Forecast can assume to access the AWS KMS key.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"schema": {
			// Property: Schema
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "Attributes": {
			//       "insertionOrder": true,
			//       "items": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "AttributeName": {
			//             "description": "Name of the dataset field",
			//             "pattern": "",
			//             "type": "string"
			//           },
			//           "AttributeType": {
			//             "description": "Data type of the field",
			//             "enum": [
			//               "string",
			//               "integer",
			//               "float",
			//               "timestamp",
			//               "geolocation"
			//             ],
			//             "type": "string"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "maxItems": 100,
			//       "minItems": 1,
			//       "type": "array"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"attributes": {
						// Property: Attributes
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"attribute_name": {
									// Property: AttributeName
									Description: "Name of the dataset field",
									Type:        types.StringType,
									Computed:    true,
								},
								"attribute_type": {
									// Property: AttributeType
									Description: "Data type of the field",
									Type:        types.StringType,
									Computed:    true,
								},
							},
							tfsdk.ListNestedAttributesOptions{},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": true,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 200,
			//   "minItems": 0,
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::Forecast::Dataset",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Forecast::Dataset").WithTerraformTypeName("awscc_forecast_dataset")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":               "Arn",
		"attribute_name":    "AttributeName",
		"attribute_type":    "AttributeType",
		"attributes":        "Attributes",
		"data_frequency":    "DataFrequency",
		"dataset_name":      "DatasetName",
		"dataset_type":      "DatasetType",
		"domain":            "Domain",
		"encryption_config": "EncryptionConfig",
		"key":               "Key",
		"kms_key_arn":       "KmsKeyArn",
		"role_arn":          "RoleArn",
		"schema":            "Schema",
		"tags":              "Tags",
		"value":             "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
