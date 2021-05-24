package infrastructure

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/jsii-runtime-go"
)

type pipeline struct {
	awscdk.Stack
}

func createPipeline(is *infrastructure) pipeline {
	var pipeline pipeline

	pipeline.Stack = awscdk.NewStack(*is.app, jsii.String(string(Pipeline)), &is.prop.Stack)
	cbsp := awsiam.NewServicePrincipal(jsii.String("codebuild.amazonaws.com"), &awsiam.ServicePrincipalOpts{
		Region: is.prop.Stack.Env.Region,
	})
	is.prop.Codebuild.ServiceRole = awsiam.NewRole(pipeline.Stack, jsii.String("codebuild"), &awsiam.RoleProps{
		AssumedBy: cbsp,
	}).RoleName()

	cbap := awscodebuild.CfnProject_ArtifactsProperty{
		Type: jsii.String("NO_ARTIFACTS"),
	}
	is.prop.Codebuild.Artifacts = cbap

	cbep := awscodebuild.CfnProject_EnvironmentProperty{
		Type:        jsii.String("LINUX_CONTAINER"),
		ComputeType: jsii.String("BUILD_GENERAL1_SMALL"),
		Image:       jsii.String("aws/codebuild/standard:4.0"),
	}
	is.prop.Codebuild.Environment = cbep

	cbs := awscodebuild.CfnProject_SourceProperty{
		Type:             jsii.String("GITHUB"),
		Location:         jsii.String("https://github.com/dmatusiewicz/demo_cdk"),
		SourceIdentifier: jsii.String("test01"),
	}
	is.prop.Codebuild.Source = cbs

	// awsiam.NewCfnRole(pipeline.Stack, jsii.String(string(Pipeline)), &awsiam.CfnRoleProps{})
	awscodebuild.NewCfnProject(pipeline.Stack, jsii.String(string(Pipeline)), &is.prop.Codebuild)
	awscodebuild.NewGitHubSourceCredentials(pipeline.Stack, jsii.String("github"), &awscodebuild.GitHubSourceCredentialsProps{
		AccessToken: awscdk.NewSecretValue("asd", &awscdk.IntrinsicProps{
			StackTrace: jsii.Bool(false),
		}),
	})

	// {
	// 	"Version": "2012-10-17",
	// 	"Statement": [
	// 	  {
	// 		"Effect": "Allow",
	// 		"Principal": {
	// 		  "Service": "codebuild.amazonaws.com"
	// 		},
	// 		"Action": "sts:AssumeRole"
	// 	  }
	// 	]
	//   }
	return pipeline
}
