package apigatewayresourcewithmethod

import (
	apigatewayresource "github.com/mohammadasim/pulumiutils/aws/components/apigateway/resources"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apigateway"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApiGatewayResourceWithMethodArgs struct {
	ApiGatewayResourceArgs *apigatewayresource.ApiGatewayResourceArgs
	Httpmethod             string
	IntegrationLambda      *lambda.Function
}

type ApiGatewayResourceWithMethod struct {
	pulumi.ResourceState
	ResourceID pulumi.StringOutput `pulumi:"ResourceID"`
}

func NewApiGatewayResourceWithMethodComponent(ctx *pulumi.Context, name string, args *ApiGatewayResourceWithMethodArgs, opts ...pulumi.ResourceOption) (*ApiGatewayResourceWithMethod, error) {
	// Created an instance of the component
	apigatewayResourceWithMethodComponent := &ApiGatewayResourceWithMethod{}
	// Registered the component with the pulumi context
	err := ctx.RegisterComponentResource("mohammadasim:pulumiutils:ApiGatewayResourceWithMethod", name, apigatewayResourceWithMethodComponent, opts...)
	if err != nil {
		return nil, err
	}
	// Create an ApiGatewayResourceComponent
	apigatewayresourcecomponent, err := apigatewayresource.NewApiGatewayResourceComponent(ctx, name, args.ApiGatewayResourceArgs)
	if err != nil {
		return nil, err
	}
	// Create a method for the resource
	_, err = apigateway.NewMethod(ctx, name, &apigateway.MethodArgs{
		Authorization: pulumi.String("NONE"),
		HttpMethod:    pulumi.String(args.Httpmethod),
		ResourceId:    apigatewayresourcecomponent.ResourceID.ApplyT(func(id string) (string, error) { return id, nil }).(pulumi.StringOutput),
		RestApi:       pulumi.String(args.ApiGatewayResourceArgs.RestApiId),
	}, pulumi.Parent(apigatewayresourcecomponent))
	if err != nil {
		return nil, err
	}

	apigatewayResourceWithMethodComponent.ResourceID = apigatewayresourcecomponent.ResourceID

	return apigatewayResourceWithMethodComponent, nil
}
