// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/squat/terraform-provider-lambda/internal/sdk"
	"github.com/squat/terraform-provider-lambda/internal/sdk/models/shared"
	"net/http"
	"os"
)

var _ provider.Provider = &LambdaProvider{}

type LambdaProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// LambdaProviderModel describes the provider data model.
type LambdaProviderModel struct {
	ServerURL  types.String `tfsdk:"server_url"`
	BearerAuth types.String `tfsdk:"bearer_auth"`
}

func (p *LambdaProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "lambda"
	resp.Version = p.version
}

func (p *LambdaProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: `Lambda Provider: API for interacting with the Lambda GPU Cloud`,
		Attributes: map[string]schema.Attribute{
			"server_url": schema.StringAttribute{
				MarkdownDescription: "Server URL (defaults to https://cloud.lambdalabs.com/api/v1/)",
				Optional:            true,
				Required:            false,
			},
			"bearer_auth": schema.StringAttribute{
				Sensitive: true,
				Optional:  true,
			},
		},
	}
}

func (p *LambdaProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data LambdaProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	ServerURL := data.ServerURL.ValueString()

	if ServerURL == "" {
		ServerURL = "https://cloud.lambdalabs.com/api/v1/"
	}

	bearerAuth := new(string)
	if !data.BearerAuth.IsUnknown() && !data.BearerAuth.IsNull() {
		*bearerAuth = data.BearerAuth.ValueString()
	} else {
		if len(os.Getenv("LAMBDA_API_KEY")) > 0 {
			*bearerAuth = os.Getenv("LAMBDA_API_KEY")
		} else {
			bearerAuth = nil
		}
	}
	security := shared.Security{
		BearerAuth: bearerAuth,
	}

	opts := []sdk.SDKOption{
		sdk.WithServerURL(ServerURL),
		sdk.WithSecurity(security),
		sdk.WithClient(http.DefaultClient),
	}
	client := sdk.New(opts...)

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *LambdaProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewSSHKeyResource,
		NewInstanceResource,
	}
}

func (p *LambdaProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewInstanceDataSource,
		NewInstanceTypesDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &LambdaProvider{
			version: version,
		}
	}
}
