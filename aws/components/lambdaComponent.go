package lambdacomponent

import (
	"encoding/json"

	"github.com/pulumi/pulumi-archive/sdk/go/archive"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LambdaComponentArgs struct {
	Name       pulumi.StringInput
	Handler    pulumi.StringInput
	Runtime    pulumi.StringInput
	SourceFile string
	OutputPath string
	EnvVars    pulumi.StringMap
	RoleName   pulumi.StringInput
}

type LambdaComponent struct {
	pulumi.ResourceState
	//Output properties for Component
	LambdaID  pulumi.StringOutput `pulumi:"lambdaID"`
	RoleID    pulumi.StringOutput `pulumi:"roleID"`
	RoleArn   pulumi.StringOutput `pulumi:"roleArn"`
	LambdaArn pulumi.StringOutput `pulumi:"lambdaArn"`
	Role      *iam.Role           `pulumi:"role"`
	Function  *lambda.Function    `pulumi:"function"`
}

func NewLambdaComponent(ctx *pulumi.Context, name string, args *LambdaComponentArgs, opts ...pulumi.ResourceOption) (*LambdaComponent, error) {
	lambdaComponent := &LambdaComponent{}

	err := ctx.RegisterComponentResource("mohammadasim:pulumiutils:LambdaComponent", name, lambdaComponent, opts...)
	if err != nil {
		return nil, err
	}
	// Create a trust policy for the role
	trustPolicyMap := map[string]interface{}{
		"Version": "2012-10-17",
		"Statement": []map[string]interface{}{
			{
				"Effect": "Allow",
				"Principal": map[string]interface{}{
					"Service": "lambda.amazonaws.com",
				},
				"Action": "sts:AssumeRole",
			},
		},
	}
	trustPolicy, err := json.Marshal(trustPolicyMap)
	if err != nil {
		return nil, err
	}
	// Get lambda execution policy
	policy, err := iam.LookupPolicy(ctx, &iam.LookupPolicyArgs{
		Name: pulumi.StringRef("AWSLambdaBasicExecutionRole"),
	})
	if err != nil {
		return nil, err
	}

	// Create lambda role
	role, err := iam.NewRole(ctx, "lambdaRole", &iam.RoleArgs{
		AssumeRolePolicy: pulumi.String(trustPolicy),
	})
	if err != nil {
		return nil, err
	}
	// Attach the policy to the role
	_, err = iam.NewRolePolicyAttachment(ctx, "lambdaPolicyAttachment", &iam.RolePolicyAttachmentArgs{
		Role:      role.Name,
		PolicyArn: pulumi.String(policy.Arn),
	})
	if err != nil {
		return nil, err
	}
	// If using the provider provided.al2023, the build file must be called bootstrap
	lambdaArchive, err := archive.LookupFile(ctx, &archive.LookupFileArgs{
		Type:       "zip",
		SourceFile: pulumi.StringRef(args.SourceFile),
		OutputPath: args.OutputPath,
	})
	if err != nil {
		return nil, err
	}
	// Create lambda function
	lambdaFunction, err := lambda.NewFunction(ctx, "lambdaFunction", &lambda.FunctionArgs{
		Code:    pulumi.NewFileArchive(lambdaArchive.OutputPath),
		Name:    args.Name,
		Handler: args.Handler,
		Runtime: args.Runtime,
		Role:    role.Arn,
		Environment: &lambda.FunctionEnvironmentArgs{
			Variables: args.EnvVars,
		},
	})
	if err != nil {
		return nil, err
	}
	lambdaComponent.LambdaArn = lambdaFunction.Arn.ToStringOutput()
	lambdaComponent.LambdaID = lambdaFunction.ID().ToStringOutput()
	lambdaComponent.RoleArn = role.Arn.ToStringOutput()
	lambdaComponent.RoleID = role.ID().ToStringOutput()
	lambdaComponent.Role = role
	lambdaComponent.Function = lambdaFunction

	return lambdaComponent, err

}
