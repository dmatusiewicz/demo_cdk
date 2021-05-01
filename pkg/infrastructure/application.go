package infrastructure

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/jsii-runtime-go"
)

type application struct {
	awscdk.Stack
}

func createApplication(is infrastructureSettings) application {
	var app application

	app.Stack = awscdk.NewStack(is.app, jsii.String("application"), is.StackProps)
	awscdk.NewCfnOutput(app.Stack, jsii.String("exportFromApp"), &awscdk.CfnOutputProps{
		Value:      jsii.String("val1"),
		ExportName: jsii.String("exportFromApp"),
	})
	return app
}
