package infrastructure

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/jsii-runtime-go"
)

type application struct {
	awscdk.Stack
}

func createApplication(is *infrastructure) application {
	var app application
	app.Stack = awscdk.NewStack(*is.app, jsii.String(string(Application)), &is.prop.StackProps)

	return app
}
