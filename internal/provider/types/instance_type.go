// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type InstanceType struct {
	Description       types.String `tfsdk:"description"`
	GpuDescription    types.String `tfsdk:"gpu_description"`
	Name              types.String `tfsdk:"name"`
	PriceCentsPerHour types.Int64  `tfsdk:"price_cents_per_hour"`
	Specs             Specs        `tfsdk:"specs"`
}
