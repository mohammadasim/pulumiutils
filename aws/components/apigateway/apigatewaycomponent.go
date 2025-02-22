package apigatewaycomponent

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apigateway"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApigatewayComponentArgs struct {
	ApiName        pulumi.StringInput
	ApiDescription pulumi.StringInput
}

type ApiGatewayComponent struct {
	pulumi.ResourceState
	ApiGatewayArn            pulumi.StringOutput `pulumi:"apiGatewayArn"`
	ApiGatewayID             pulumi.StringOutput `pulumi:"apiGatewayID"`
	ApiGatewayRootResourceID pulumi.StringOutput `pulumi:"apiGatewayRootResourceID"`
}

func NewApiGatewayComponent(ctx *pulumi.Context, name string, args *ApigatewayComponentArgs, opts ...pulumi.ResourceOption) (*ApiGatewayComponent, error) {
	apigatewaycomponent := &ApiGatewayComponent{}
	err := ctx.RegisterComponentResource("mohammadasim:pulumiutils:ApiGatewayComponent", name, apigatewaycomponent, opts...)
	if err != nil {
		return nil, err
	}

	apigateway, err := apigateway.NewRestApi(ctx, fmt.Sprintf("%s-apigateway", args.ApiName), &apigateway.RestApiArgs{
		Name:        args.ApiName,
		Description: args.ApiDescription,
	})
	if err != nil {
		return nil, err
	}

	apigatewaycomponent.ApiGatewayArn = apigateway.Arn
	apigatewaycomponent.ApiGatewayID = apigateway.ID().ToStringOutput()
	apigatewaycomponent.ApiGatewayRootResourceID = apigateway.RootResourceId

	return apigatewaycomponent, nil
}
