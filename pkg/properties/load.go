package properties

import (
	"embed"
	"log"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

const (
	fileType         string = ".yaml"
	workingDirectory string = "."
)

//go:embed files/*
var files embed.FS
var cnf Config

type Config struct {
	config interface{}
}

func (c Config) Get() interface{} {
	return c.config
}

func New(id *string) Config {
	var fileToRead strings.Builder

	fileToRead.WriteString(*id)
	fileToRead.WriteString(fileType)

	d, err := files.ReadDir(workingDirectory)
	if err != nil {
		log.Fatal(err)
	}

	// This is safe as long as embed loads only one directory.
	data, err := files.ReadFile(filepath.Join(d[0].Name(), fileToRead.String()))
	if err != nil {
		log.Fatal(err)
	}
	var yamlData interface{}
	err = yaml.Unmarshal([]byte(data), &yamlData)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}
	cnf.config = yamlData
	return cnf
}
