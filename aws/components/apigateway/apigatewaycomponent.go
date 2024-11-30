package apigatewaycomponent

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apigateway"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApigatewayComponentArgs struct {
	apiName        pulumi.StringInput
	apiDescription pulumi.StringInput
}

type ApiGatewayComponent struct {
	pulumi.ResourceState
	apiGatewayArn            pulumi.StringOutput `pulumi:"apiGatewayArn"`
	apiGatewayID             pulumi.StringOutput `pulumi:"apiGatewayID"`
	apiGatewayRootResourceID pulumi.StringOutput `pulumi:"apiGatewayRootResourceID"`
}

func NewApiGatewayComponent(ctx *pulumi.Context, name string, args *ApigatewayComponentArgs, opts ...pulumi.ResourceOption) (*ApiGatewayComponent, error) {
	apigatewaycomponent := &ApiGatewayComponent{}
	err := ctx.RegisterComponentResource("apigatewaycomponent:index:ApiGatewayComponent", name, apigatewaycomponent, opts...)
	if err != nil {
		return nil, err
	}

	apigateway, err := apigateway.NewRestApi(ctx, fmt.Sprintf("%s-apigateway", args.apiName), &apigateway.RestApiArgs{
		Name:        args.apiName,
		Description: args.apiDescription,
	})
	if err != nil {
		return nil, err
	}

	apigatewaycomponent.apiGatewayArn = apigateway.Arn
	apigatewaycomponent.apiGatewayID = apigateway.ID().ToStringOutput()
	apigatewaycomponent.apiGatewayRootResourceID = apigateway.RootResourceId

	return apigatewaycomponent, nil
}
