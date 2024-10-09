package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/squat/terraform-provider-lambda/internal/provider/types"
	"github.com/squat/terraform-provider-lambda/internal/sdk"
	"github.com/squat/terraform-provider-lambda/internal/sdk/models/operations"
	"github.com/squat/terraform-provider-lambda/internal/sdk/models/shared"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &InstanceResource{}
var _ resource.ResourceWithImportState = &InstanceResource{}

func NewInstanceResource() resource.Resource {
	return &InstanceResource{}
}

// InstanceResource defines the resource implementation.
type InstanceResource struct {
	client *sdk.Lambda
}

// InstanceResourceModel describes the resource data model.
type InstanceResourceModel struct {
	RegionName       types.String          `tfsdk:"region_name"`
	InstanceTypeName types.String          `tfsdk:"instance_type_name"`
	FileSystemNames  []types.String        `tfsdk:"file_system_names"`
	Hostname         types.String          `tfsdk:"hostname"`
	ID               types.String          `tfsdk:"id"`
	InstanceType     *tfTypes.InstanceType `tfsdk:"instance_type"`
	IP               types.String          `tfsdk:"ip"`
	IsReserved       types.Bool            `tfsdk:"is_reserved"`
	JupyterToken     types.String          `tfsdk:"jupyter_token"`
	JupyterURL       types.String          `tfsdk:"jupyter_url"`
	Name             types.String          `tfsdk:"name"`
	PrivateIP        types.String          `tfsdk:"private_ip"`
	Region           *tfTypes.Region       `tfsdk:"region"`
	SSHKeyNames      []types.String        `tfsdk:"ssh_key_names"`
	Status           types.String          `tfsdk:"status"`
}

func (r *InstanceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_instance"
}

func (r *InstanceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Instance Resource",
		Attributes: map[string]schema.Attribute{
			"region_name": schema.StringAttribute{
				Required:    true,
				Description: `Short name of the reguin in which to laucn the instance.`,
			},
			"instance_type_name": schema.StringAttribute{
				Required:    true,
				Description: `Name of an instance type.`,
			},
			"file_system_names": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
				Description: `Names of the file systems, if any, attached to the instance`,
			},
			"hostname": schema.StringAttribute{
				Computed:    true,
				Description: `Hostname assigned to this instance, which resolves to the instance's IP.`,
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `The unique identifier (ID) of the instance`,
			},
			"instance_type": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"description": schema.StringAttribute{
						Computed:    true,
						Description: `Long name of the instance type`,
					},
					"gpu_description": schema.StringAttribute{
						Computed:    true,
						Description: `Description of the GPU(s) in the instance type`,
					},
					"name": schema.StringAttribute{
						Computed:    true,
						Description: `Name of an instance type`,
					},
					"price_cents_per_hour": schema.Int64Attribute{
						Computed:    true,
						Description: `Price of the instance type, in US cents per hour`,
					},
					"specs": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"gpus": schema.Int64Attribute{
								Computed:    true,
								Description: `Number of GPUs`,
							},
							"memory_gib": schema.Int64Attribute{
								Computed:    true,
								Description: `Amount of RAM, in gibibytes (GiB)`,
							},
							"storage_gib": schema.Int64Attribute{
								Computed:    true,
								Description: `Amount of storage, in gibibytes (GiB).`,
							},
							"vcpus": schema.Int64Attribute{
								Computed:    true,
								Description: `Number of virtual CPUs`,
							},
						},
					},
				},
				Description: `Hardware configuration and pricing of an instance type`,
			},
			"ip": schema.StringAttribute{
				Computed:    true,
				Description: `IPv4 address of the instance`,
			},
			"is_reserved": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether the instance is reserved.`,
			},
			"jupyter_token": schema.StringAttribute{
				Computed:    true,
				Description: `Secret token used to log into the jupyter lab server hosted on the instance.`,
			},
			"jupyter_url": schema.StringAttribute{
				Computed:    true,
				Description: `URL that opens a jupyter lab notebook on the instance.`,
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `User-provided name for the instance`,
			},
			"private_ip": schema.StringAttribute{
				Computed:    true,
				Description: `Private IPv4 address of the instance`,
			},
			"region": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"description": schema.StringAttribute{
						Computed:    true,
						Description: `Long name of a region`,
					},
					"name": schema.StringAttribute{
						Computed:    true,
						Description: `Short name of a region`,
					},
				},
			},
			"ssh_key_names": schema.ListAttribute{
				Required:    true,
				ElementType: types.StringType,
				Description: `Names of the SSH keys allowed to access the instance`,
			},
			"status": schema.StringAttribute{
				Computed:    true,
				Description: `The current status of the instance`,
			},
		},
	}
}

func (r *InstanceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.Lambda)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.Lambda, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *InstanceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *InstanceResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	request := *data.ToSharedLaunch()
	println(data.InstanceTypeName.ValueString())
	println(data.RegionName.ValueString())
	res, err := r.client.LaunchInstance(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.Launch != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}

	getRequest := operations.GetInstanceRequest{
		ID: res.Launch.Data.InstanceIds[0],
	}
	getResponse, err := r.client.GetInstance(ctx, getRequest)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if getResponse != nil && getResponse.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(getResponse.RawResponse))
		}
		return
	}
	if getResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", getResponse))
		return
	}
	if getResponse.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if getResponse.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", getResponse.StatusCode), debugResponse(getResponse.RawResponse))
		return
	}
	if !(getResponse.Instance != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(getResponse.RawResponse))
		return
	}
	data.RefreshFromSharedInstance(&getResponse.Instance.Data)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *InstanceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *InstanceResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	id := data.ID.ValueString()

	request := operations.GetInstanceRequest{
		ID: id,
	}
	res, err := r.client.GetInstance(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.Instance != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedInstance(&res.Instance.Data)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *InstanceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *InstanceResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; all attributes marked as RequiresReplace

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *InstanceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *InstanceResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	id := data.ID.ValueString()

	request := shared.Terminate{
		InstanceIds: []string{id},
	}
	res, err := r.client.TerminateInstance(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *InstanceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource ssh_key.")
}
