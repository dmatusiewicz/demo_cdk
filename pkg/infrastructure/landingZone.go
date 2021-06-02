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

	// START_PARAMETER OMIT
	awsssm.NewStringParameter(lz.Stack, nameBuilder(lz.Stack, "ssm-keyArn"), // HL1
		&awsssm.StringParameterProps{ // HL1
			StringValue:   k.KeyArn(),
			ParameterName: nameBuilder(lz.Stack, "keyArn"),
			Description:   jsii.String("KeyARN parameter - reference it from another component. Loosely coupled."),
		})
	// END_PARAMETER OMIT
	// START_OUTPUT OMIT
	awscdk.NewCfnOutput(lz.Stack, nameBuilder(lz.Stack, "output-keyArn"), // HL1
		&awscdk.CfnOutputProps{ // HL1
			ExportName:  nameBuilder(lz.Stack, "keyArn"),
			Value:       k.KeyArn(),
			Description: jsii.String("KeyARN parameter - import it to another component. Mind tight components coupling."),
		})
	// END_OUTPUT OMIT

	return lz
}
