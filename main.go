package main

import (
	"demo_cdk/pkg/infrastructure"

	"github.com/aws/aws-cdk-go/awscdk"
)

var environmentName = "dev"

const (
	application infrastructure.InfrastructureType = "application"
	managemeent infrastructure.InfrastructureType = "management"
)

func main() {
	app := awscdk.NewApp(nil)

	infrastructure.Factory(app, environmentName, application)
	// fmt.Println(s)
	app.Synth(nil)
}
