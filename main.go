package main

import (
	"demo_cdk/pkg/infrastructure"

	"github.com/aws/aws-cdk-go/awscdk"
)

// This will be read from the environment variables.
var environmentName = "dev"

func main() {
	app := awscdk.NewApp(nil)

	infrastructure.Factory(app, environmentName, infrastructure.Application)
	infrastructure.Factory(app, environmentName, infrastructure.Pipeline)

	app.Synth(nil)
}
