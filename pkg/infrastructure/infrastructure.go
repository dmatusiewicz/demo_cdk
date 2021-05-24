package infrastructure

import (
	"embed"
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscodebuild"
	"github.com/aws/jsii-runtime-go"
	"gopkg.in/yaml.v2"
)

type InfrastructureType string

const (
	Application      InfrastructureType = "application"
	Pipeline         InfrastructureType = "pipeline"
	LandingZone      InfrastructureType = "landingZone"
	fileType         string             = ".yaml"
	workingDirectory string             = "."
)

type infrastructure struct {
	app  *awscdk.App
	prop *infrastructureProperties
}

type infrastructureProperties struct {
	Stack             awscdk.StackProps            `yaml:"stack"`
	Codebuild         awscodebuild.CfnProjectProps `yaml:"codebuild"`
	LandingZoneConfig map[string]interface{}       `yaml:"landingZone"`
}

var infra infrastructure

//go:embed conf/*
var conf embed.FS

func Factory(app awscdk.App, envName string, infraType InfrastructureType) awscdk.Stack {

	infra = infrastructure{
		prop: getInfraProperties(envName),
		app:  &app,
	}
	iT := *infra.prop.Stack.Tags
	iT["Environment"] = jsii.String(envName)
	iT["InfastructureType"] = jsii.String(string(infraType))
	iT["DeploymentTime"] = jsii.String(time.Now().UTC().String())

	fmt.Println(infra.prop.Stack.Tags)
	if infraType == Application {
		return createApplication(&infra).Stack
	}

	if infraType == Pipeline {
		return createPipeline(&infra).Stack
	}
	if infraType == LandingZone {
		return createLandingZone(&infra).Stack
	}

	return nil
}

func nameBuilder(s awscdk.Stack, name string) *string {
	var sb strings.Builder
	sb.WriteString(*s.StackName())
	sb.WriteString("-")
	sb.WriteString(name)

	return jsii.String(sb.String())
}

func getInfraProperties(envName string) *infrastructureProperties {
	var fileToRead strings.Builder

	fileToRead.WriteString(envName)
	fileToRead.WriteString(fileType)

	d, err := conf.ReadDir(workingDirectory)
	if err != nil {
		log.Fatal(err)
	}

	// This is safe as long as embed loads only one directory.
	data, err := conf.ReadFile(filepath.Join(d[0].Name(), fileToRead.String()))
	if err != nil {
		log.Fatal(err)
	}

	ip := infrastructureProperties{}

	err = yaml.Unmarshal([]byte(data), &ip)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}

	return &ip
}
