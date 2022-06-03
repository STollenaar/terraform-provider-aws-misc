package awsprofiler

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ProfileDetails -
type ProfileDetails struct {
	Name               types.String `tfsdk:"name"`
	AWSAccessKeyId     types.String `tfsdk:"aws_access_key_id"`
	AWSSecretAccessKey types.String `tfsdk:"aws_secret_access_key"`
	AWSSessionToken    types.String `tfsdk:"aws_session_token"`
	Region             types.String `tfsdk:"region"`
	Output             types.String `tfsdk:"output"`
}
