package infrastructure

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
	"github.com/aws/aws-cdk-go/awscdk/awsssm"
	"github.com/aws/jsii-runtime-go"
)

type landingZone struct {
	awscdk.Stack
}

func createLandingZone(is *infrastructure) landingZone {
	var lz landingZone
	lz.Stack = awscdk.NewStack(*is.app, jsii.String("landingZone"), &is.prop.Stack)

	fmt.Println(*lz.Stack.StackName())
	// fmt.Println(lz.Stack.ParentStack().StackName())

	k := awskms.NewKey(lz.Stack, nameBuilder(lz.Stack, "kms"), &awskms.KeyProps{})
	k.AddAlias(nameBuilder(lz.Stack, "kms"))

	awsssm.NewStringParameter(lz.Stack, nameBuilder(lz.Stack, "ssm-kms"), &awsssm.StringParameterProps{
		StringValue:   k.KeyArn(),
		ParameterName: nameBuilder(lz.Stack, "ssm-kms"),
		Description:   jsii.String("Landing Zone KMS Key ARN."),
	})

	awscdk.NewCfnOutput(lz.Stack, nameBuilder(lz.Stack, "Key_ARN"), &awscdk.CfnOutputProps{
		ExportName: jsii.String("Key-ARN"),
		Value:      k.KeyArn(),
	})

	return lz
}
