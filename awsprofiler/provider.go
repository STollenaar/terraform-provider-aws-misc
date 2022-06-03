package awsprofiler

import (
	"context"
	"fmt"
	"os"

	awsprofilerclient "github.com/STollenaar/aws-profiler-client"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

var stderr = os.Stderr

func New() tfsdk.Provider {
	return &provider{}
}

type provider struct {
	configured bool
	client     *awsprofilerclient.Client
}

// Provider schema struct
type providerData struct{}

func (p *provider) Configure(ctx context.Context, req tfsdk.ConfigureProviderRequest, resp *tfsdk.ConfigureProviderResponse) {
	fmt.Fprintf(stderr, "[DEBUG]- Already encountered an error")
	var config providerData
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		fmt.Fprint(stderr, "[DEBUG]- Already encountered an error")
		return
	}
	c, err := awsprofilerclient.NewClient()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create client",
			"Unable to create aws profiler client:\n\n"+err.Error(),
		)
		return
	}

	p.configured = true
	p.client = c
}

// GetSchema -
func (p *provider) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{},
	}, nil
}

// GetDataSources - Defines provider data sources
func (p *provider) GetDataSources(_ context.Context) (map[string]tfsdk.DataSourceType, diag.Diagnostics) {
	return map[string]tfsdk.DataSourceType{
		"awsprofiler_list": dataSourceListProfilesType{},
	}, nil
}

// GetResources - Defines provider resources
func (p *provider) GetResources(_ context.Context) (map[string]tfsdk.ResourceType, diag.Diagnostics) {
	return map[string]tfsdk.ResourceType{}, nil
}
