package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/squat/terraform-provider-lambda/internal/provider/types"
	"github.com/squat/terraform-provider-lambda/internal/sdk/models/shared"
)

func (r *InstanceResourceModel) ToSharedLaunch() *shared.Launch {
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = nil
	}

	quantity := int64(1)
	out := shared.Launch{
		Name:             name,
		Quantity:         &quantity,
		RegionName:       r.RegionName.ValueString(),
		InstanceTypeName: r.InstanceTypeName.ValueString(),
	}
	out.SSHKeyNames = []string{}
	for _, v := range r.SSHKeyNames {
		out.SSHKeyNames = append(out.SSHKeyNames, v.ValueString())
	}

	out.FileSystemNames = []string{}
	for _, v := range r.FileSystemNames {
		out.FileSystemNames = append(out.FileSystemNames, v.ValueString())
	}

	return &out
}

func (r *InstanceResourceModel) ToSharedUpdate() *shared.Update {
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = nil
	}

	out := shared.Update{
		ID:   r.ID.ValueString(),
		Name: name,
	}

	return &out
}

func (r *InstanceResourceModel) RefreshFromSharedInstance(resp *shared.Instance) {
	r.FileSystemNames = []types.String{}
	for _, v := range resp.FileSystemNames {
		r.FileSystemNames = append(r.FileSystemNames, types.StringValue(v))
	}
	r.Hostname = types.StringPointerValue(resp.Hostname)
	r.ID = types.StringValue(resp.ID)
	if resp.InstanceType == nil {
		r.InstanceType = nil
	} else {
		r.InstanceType = &tfTypes.InstanceType{}
		r.InstanceType.Description = types.StringValue(resp.InstanceType.Description)
		r.InstanceType.GpuDescription = types.StringValue(resp.InstanceType.GpuDescription)
		r.InstanceType.Name = types.StringValue(resp.InstanceType.Name)
		r.InstanceTypeName = types.StringValue(resp.InstanceType.Name)
		r.InstanceType.PriceCentsPerHour = types.Int64Value(resp.InstanceType.PriceCentsPerHour)
		r.InstanceType.Specs.Gpus = types.Int64Value(resp.InstanceType.Specs.Gpus)
		r.InstanceType.Specs.MemoryGib = types.Int64Value(resp.InstanceType.Specs.MemoryGib)
		r.InstanceType.Specs.StorageGib = types.Int64Value(resp.InstanceType.Specs.StorageGib)
		r.InstanceType.Specs.Vcpus = types.Int64Value(resp.InstanceType.Specs.Vcpus)
	}
	r.IP = types.StringPointerValue(resp.IP)
	r.IsReserved = types.BoolPointerValue(resp.IsReserved)
	r.JupyterToken = types.StringPointerValue(resp.JupyterToken)
	r.JupyterURL = types.StringPointerValue(resp.JupyterURL)
	r.Name = types.StringPointerValue(resp.Name)
	r.PrivateIP = types.StringPointerValue(resp.PrivateIP)
	if resp.Region == nil {
		r.Region = nil
	} else {
		r.Region = &tfTypes.Region{}
		r.Region.Description = types.StringValue(resp.Region.Description)
		r.Region.Name = types.StringValue(resp.Region.Name)
		r.RegionName = types.StringValue(resp.Region.Name)
	}
	r.SSHKeyNames = []types.String{}
	for _, v := range resp.SSHKeyNames {
		r.SSHKeyNames = append(r.SSHKeyNames, types.StringValue(v))
	}
	r.Status = types.StringValue(string(resp.Status))
}
