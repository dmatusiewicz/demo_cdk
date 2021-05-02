package infrastructure

import (
	"embed"
	"log"
	"path/filepath"
	"strings"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscodebuild"
	"gopkg.in/yaml.v2"
)

type InfrastructureType string

const (
	Application      InfrastructureType = "application"
	Pipeline         InfrastructureType = "pipeline"
	fileType         string             = ".yaml"
	workingDirectory string             = "."
)

type infrastructure struct {
	app  *awscdk.App
	prop *infrastructureProperties
}

type infrastructureProperties struct {
	Stack      awscdk.StackProps            `yaml:"stack"`
	Codebuild  awscodebuild.CfnProjectProps `yaml:"codebuild"`
	CustomTest string                       `yaml:"customTestKey"`
}

type cbp struct {
}

//go:embed conf/*
var conf embed.FS

func Factory(app awscdk.App, envName string, infraType InfrastructureType) (awscdk.Stack, error) {

	i := infrastructure{
		prop: getInfraProperties(envName, app, infraType),
		app:  &app,
	}

	if infraType == Application {
		return createApplication(&i).Stack, nil
	}

	if infraType == Pipeline {
		return createPipeline(&i).Stack, nil
	}

	return nil, nil
}

func getInfraProperties(envName string, app awscdk.App, it InfrastructureType) *infrastructureProperties {
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
