package apigatewayresource

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apigateway"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApiGatewayResourceArgs struct {
	ParentId  string
	PathPart  string
	restApiId string
}

type ApiGatewayResourceComponent struct {
	pulumi.ResourceState
	resourceID pulumi.StringOutput `pulumi:"ResourceID"`
}

func NewApiGatewayResourceComponent(ctx *pulumi.Context, name string, args *ApiGatewayResourceArgs, opts ...pulumi.ResourceOption) (*ApiGatewayResourceComponent, error) {
	apigatewayResourceComponent := &ApiGatewayResourceComponent{}
	err := ctx.RegisterComponentResource("mohammadasim:pulumiutils:ApiGatewayResource", name, apigatewayResourceComponent, opts...)
	if err != nil {
		return nil, err
	}

	resource, err := apigateway.NewResource(ctx, fmt.Sprintf("%s-resource", name), &apigateway.ResourceArgs{
		ParentId: pulumi.String(args.ParentId),
		PathPart: pulumi.String(args.PathPart),
		RestApi:  pulumi.String(args.restApiId),
	})
	if err != nil {
		return nil, err
	}

	apigatewayResourceComponent.resourceID = resource.ID().ToStringOutput()
	return apigatewayResourceComponent, nil
}
