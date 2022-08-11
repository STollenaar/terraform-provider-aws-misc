package profiler

import (
	"context"
	"fmt"
	"os"

	awsmiscclient "github.com/STollenaar/aws-misc-client"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DataSourceListProfilesType struct{}

func (r DataSourceListProfilesType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"profiles": {
				Computed: true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Type:     types.StringType,
						Computed: true,
					},
					"aws_access_key_id": {
						Type:     types.StringType,
						Computed: true,
					},
					"aws_secret_access_key": {
						Type:      types.StringType,
						Computed:  true,
						Sensitive: true,
					},
					"aws_session_token": {
						Type:      types.StringType,
						Computed:  true,
						Sensitive: true,
					},
					"region": {
						Type:     types.StringType,
						Computed: true,
					},
					"output": {
						Type:     types.StringType,
						Computed: true,
					},
				}, tfsdk.ListNestedAttributesOptions{}),
			},
		},
	}, nil
}

func (r DataSourceListProfilesType) NewDataSource(ctx context.Context, p tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	c, err := awsmiscclient.NewClient()
	if err != nil {
		d := diag.Diagnostics{}
		d.AddError(
			"Unable to create client",
			"Unable to create aws profiler client:\n\n"+err.Error(),
		)
		return nil, d
	}

	return dataSourceProfiles{
		p:      p,
		client: c,
	}, nil
}

type dataSourceProfiles struct {
	p      tfsdk.Provider
	client *awsmiscclient.Client
}

func (r dataSourceProfiles) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	// Declare struct that this function will set to this data source's state
	var resourceState struct {
		Profile []ProfileDetails `tfsdk:"profiles"`
	}

	profiles, err := r.client.Profiler.GetProfiles()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error retrieving profiles",
			err.Error(),
		)
		return
	}
	for _, profile := range profiles {
		p := ProfileDetails{
			Name:               types.String{Value: profile.Name},
			AWSAccessKeyId:     types.String{Value: profile.AWSAccessKeyId},
			AWSSecretAccessKey: types.String{Value: profile.AWSSecretAccessKey},
			AWSSessionToken:    types.String{Value: profile.AWSSessionToken},
		}
		if profile.Region != "" {
			p.Region = types.String{Value: profile.Region}
		}
		if profile.Output != "" {
			p.Output = types.String{Value: profile.Output}
		}
		resourceState.Profile = append(resourceState.Profile, p)
	}

	fmt.Fprintf(os.Stderr, "[DEBUG]-Resource State:%+v", resourceState)

	// Set state
	diags := resp.State.Set(ctx, &resourceState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		fmt.Fprint(os.Stderr, "[DEBUG]- Encountered an error while reading")
		return
	}
}
