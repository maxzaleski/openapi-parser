package main

import (
	"fmt"
	"io"
	"log"
	"os"

	openapi_parser "openapi-gen/openapi-parser"
)

type PreParsedDocument struct {
	Meta        *Meta                  `yaml:"info"`
	Definitions map[string]interface{} `yaml:"definitions"`
	Paths       map[string]interface{} `yaml:"paths"`
	Responses   map[string]interface{} `yaml:"responses"`
	Hosts       string                 `yaml:"host"`
}

type Meta struct {
	Title   string `yaml:"title"`
	Version string `yaml:"version"`
}

// Definition represents a type definition.
type Definition struct {
	Key         string
	Description string
	Required    []string
	Properties  []*DefinitionProperty
}

// DefinitionProperty represents a property of `Definition`.
type DefinitionProperty struct {
	// The model's name.
	Key string
	// The model's type.
	Type string
	// The model's description.
	Description string
	// The model's reference key.
	Ref string
	// Whether the property is required.
	Required bool
	// The model's validation properties.
	Validation *DefinitionPropertyValidation
}

// DefinitionPropertyValidation represents the validation properties of a `DefinitionProperty`.
type DefinitionPropertyValidation struct {
	// The property's regex pattern to match (string).
	Pattern string
	// The property's maximum length (string).
	MaxLength int
	// The property's minimum length (string).
	MinLength int
	// The property's maximum items (slice).
	MaxItems int
	// The property's minimum items (slice).
	MinItems int
}

func main() {
	specFile, err := os.Open("./spec.yml")
	if err != nil {
		log.Fatalln(err)
	}
	b, err := io.ReadAll(specFile)
	if err != nil {
		log.Fatalln(err)
	}

	doc, err := openapi_parser.NewDocument(b)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(doc)
}
