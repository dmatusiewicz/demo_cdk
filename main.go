package main

import (
	"demo_cdk/pkg/infrastructure"

	"github.com/aws/aws-cdk-go/awscdk"
)

// This will be read from the environment variables.
var environmentName = "dev"

// START OMIT
func main() {
	app := awscdk.NewApp(nil)
	infrastructure.Factory(app, environmentName, infrastructure.Application) // HL1
	infrastructure.Factory(app, environmentName, infrastructure.Pipeline)    // HL1
	infrastructure.Factory(app, environmentName, infrastructure.LandingZone)
	app.Synth(nil) // OMIT
}

// END OMIT
