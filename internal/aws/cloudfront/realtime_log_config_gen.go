// Code generated by generators/resource/main.go; DO NOT EDIT.

package cloudfront

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
	registry.AddResourceTypeFactory("awscc_cloudfront_realtime_log_config", realtimeLogConfigResourceType)
}

// realtimeLogConfigResourceType returns the Terraform awscc_cloudfront_realtime_log_config resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::CloudFront::RealtimeLogConfig resource type.
func realtimeLogConfigResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"end_points": {
			// Property: EndPoints
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "KinesisStreamConfig": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "RoleArn": {
			//             "type": "string"
			//           },
			//           "StreamArn": {
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "RoleArn",
			//           "StreamArn"
			//         ],
			//         "type": "object"
			//       },
			//       "StreamType": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "KinesisStreamConfig",
			//       "StreamType"
			//     ],
			//     "type": "object"
			//   },
			//   "minItems": 1,
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"kinesis_stream_config": {
						// Property: KinesisStreamConfig
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"role_arn": {
									// Property: RoleArn
									Type:     types.StringType,
									Required: true,
								},
								"stream_arn": {
									// Property: StreamArn
									Type:     types.StringType,
									Required: true,
								},
							},
						),
						Required: true,
					},
					"stream_type": {
						// Property: StreamType
						Type:     types.StringType,
						Required: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{
					MinItems: 1,
				},
			),
			Required: true,
		},
		"fields": {
			// Property: Fields
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "type": "string"
			//   },
			//   "minItems": 1,
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Type: types.ListType{ElemType: types.StringType},
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtLeast(1),
			},
			Required: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			// Name is a force-new attribute.
		},
		"sampling_rate": {
			// Property: SamplingRate
			// CloudFormation resource type schema:
			// {
			//   "type": "number"
			// }
			Type:     types.NumberType,
			Required: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::CloudFront::RealtimeLogConfig",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFront::RealtimeLogConfig").WithTerraformTypeName("awscc_cloudfront_realtime_log_config")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                   "Arn",
		"end_points":            "EndPoints",
		"fields":                "Fields",
		"kinesis_stream_config": "KinesisStreamConfig",
		"name":                  "Name",
		"role_arn":              "RoleArn",
		"sampling_rate":         "SamplingRate",
		"stream_arn":            "StreamArn",
		"stream_type":           "StreamType",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_cloudfront_realtime_log_config", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
