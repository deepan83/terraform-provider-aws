// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package transfer

import (
	"context"

	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceServer,
			TypeName: "aws_transfer_server",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceAccess,
			TypeName: "aws_transfer_access",
		},
		{
			Factory:  ResourceServer,
			TypeName: "aws_transfer_server",
		},
		{
			Factory:  ResourceSSHKey,
			TypeName: "aws_transfer_ssh_key",
		},
		{
			Factory:  ResourceTag,
			TypeName: "aws_transfer_tag",
		},
		{
			Factory:  ResourceUser,
			TypeName: "aws_transfer_user",
		},
		{
			Factory:  ResourceWorkflow,
			TypeName: "aws_transfer_workflow",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Transfer
}

var ServicePackage = &servicePackage{}
