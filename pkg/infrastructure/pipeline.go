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

	// awsiam.NewCfnRole(pipeline.Stack, jsii.String(string(Pipeline)), &awsiam.CfnRoleProps{})
	awscodebuild.NewCfnProject(pipeline.Stack, jsii.String(string(Pipeline)), &is.prop.Codebuild)
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
