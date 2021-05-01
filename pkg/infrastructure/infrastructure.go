package infrastructure

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk"
)

type InfrastructureType string

type infrastructureSettings struct {
	app awscdk.App
	*awscdk.StackProps
}

func Factory(a awscdk.App, e string, it InfrastructureType) (awscdk.Stack, error) {
	var is infrastructureSettings

	is.app = a

	if it == "application" {
		return createApplication(is).Stack, nil
	}

	if it == "management" {
		return createManagement(is).Stack, nil
	}

	return nil, fmt.Errorf("%s infrastrucutre type not found", it)
}
