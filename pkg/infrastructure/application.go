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
	// START_IMPORT OMIT
	kms := awscdk.Fn_ImportValue(jsii.String("Key-ARN"))
	fmt.Println(&kms)
	awscdk.NewCfnOutput(app.Stack, nameBuilder(app.Stack, "output-test"), // HL1
		&awscdk.CfnOutputProps{ // HL1
			Value:      kms,
			ExportName: jsii.String("APP-KMS"),
		})
	// END_IMPORT OMIT
	return app
}
