package infrastructure

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/jsii-runtime-go"
)

type application struct {
	awscdk.Stack
}

func createApplication(is *infrastructure) application {
	var app application
	app.Stack = awscdk.NewStack(*is.app, jsii.String(string(Application)), &is.prop.Stack)

	kms := awscdk.Fn_ImportValue(jsii.String("Key-ARN"))
	fmt.Println(&kms)
	awscdk.NewCfnOutput(app.Stack, nameBuilder(app.Stack, "output-test"), &awscdk.CfnOutputProps{
		Value:      kms,
		ExportName: jsii.String("APP-KMS"),
	})
	return app
}
