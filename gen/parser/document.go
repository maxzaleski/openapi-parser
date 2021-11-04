package parser

import (
	"strings"

	"gopkg.in/yaml.v2"
)

type preParsedDocument struct {
	Meta        *DocumentMeta          `yaml:"info"`
	Definitions map[string]interface{} `yaml:"definitions"`
	Paths       map[string]interface{} `yaml:"paths"`
	Responses   map[string]interface{} `yaml:"responses"`
	Hosts       string                 `yaml:"host"`
	BasePath    string                 `yaml:"basePath"`
}

// DocumentMeta represents the document's metadata.
type DocumentMeta struct {
	Title   string `yaml:"title"`
	Version string `yaml:"version"`
}

// Document represents a parsed document.
type Document struct {
	// The document's metadata.
	Meta *DocumentMeta
	// The API's hosts.
	Hosts []string
	// The API's base path.
	BasePath string
	// The API's type definitions.
	Definitions map[string]*Definition
	// The API's responses.
	Responses map[string]*Response
	// The API's paths.
	Paths map[string]*Path
}

// NewDocument returns a new instance of `Document`.
func NewDocument(b []byte) (*Document, error) {
	var doc preParsedDocument
	if err := yaml.Unmarshal(b, &doc); err != nil {
		return nil, err
	}
	return &Document{
		Meta:        doc.Meta,
		BasePath:    doc.BasePath,
		Hosts:       strings.Split(doc.Hosts, ", "),
		Definitions: parseIntoDefinitions(doc.Definitions),
		Responses:   parseIntoResponses(doc.Responses),
		Paths:       parseIntoPaths(doc.Paths),
	}, nil
}
