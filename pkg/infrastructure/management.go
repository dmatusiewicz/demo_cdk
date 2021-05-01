package infrastructure

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/jsii-runtime-go"
)

type management struct {
	awscdk.Stack
}

func createManagement(is infrastructureSettings) management {
	var mgmt management

	mgmt.Stack = awscdk.NewStack(is.app, jsii.String("management"), is.StackProps)

	return mgmt
}
