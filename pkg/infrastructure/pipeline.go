package infrastructure

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/jsii-runtime-go"
)

type pipeline struct {
	awscdk.Stack
}

func createPipeline(is *infrastructure) pipeline {
	var pipeline pipeline
	pipeline.Stack = awscdk.NewStack(*is.app, jsii.String(string(Pipeline)), &is.prop.StackProps)

	return pipeline
}
