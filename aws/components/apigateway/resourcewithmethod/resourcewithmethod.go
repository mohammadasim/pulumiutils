package apigatewayresourcewithmethod

import (
	apigatewayresource "github.com/mohammadasim/pulumiutils/aws/components/apigateway/resources"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
)

type ApiGatewayResourceWithMethodArgs struct {
	ApiGatewayResource *apigatewayresource.ApiGatewayResourceComponent
	Httpmethod         string
	ParentId           string
	RestApiId          string
	IntegrationLambda  *lambda.Function
}
