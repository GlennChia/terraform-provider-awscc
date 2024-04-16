// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package planmodifiers

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type objectDefaultValueAttributePlanModifier struct {
	defaultVal types.Object
}

// ObjectDefaultValue return an AttributePlanModifier that sets the specified value if the planned value is Null.
func ObjectDefaultValue(defaultVal types.Object) planmodifier.Object {
	return objectDefaultValueAttributePlanModifier{
		defaultVal: defaultVal,
	}
}

func (m objectDefaultValueAttributePlanModifier) Description(_ context.Context) string {
	return fmt.Sprintf("value defaults to %v", m.defaultVal)
}

func (m objectDefaultValueAttributePlanModifier) MarkdownDescription(ctx context.Context) string {
	return fmt.Sprintf("value defaults to `%v`", m.defaultVal)
}

func (m objectDefaultValueAttributePlanModifier) PlanModifyObject(ctx context.Context, request planmodifier.ObjectRequest, response *planmodifier.ObjectResponse) {
	// NoOp.
	response.PlanValue = request.PlanValue
}
