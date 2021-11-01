package openapi_parser

type preParsedDocument struct {
	Meta        *DocumentMeta          `yaml:"info"`
	Definitions map[string]interface{} `yaml:"definitions"`
	Paths       map[string]interface{} `yaml:"paths"`
	Responses   map[string]interface{} `yaml:"responses"`
	Hosts       string                 `yaml:"host"`
}

// DocumentMeta represents the document's metadata.
type DocumentMeta struct {
	Title   string `yaml:"title"`
	Version string `yaml:"version"`
}

// Document represents a parsed document.
type Document struct {
	// The document's metadata.
	Meta *DocumentMeta `yaml:"info"`
	// The API's hosts.
	Hosts string `yaml:"host"`
	// The API's type definitions.
	Definitions map[string]*Definition
	// The API's responses.
	Responses map[string]*Response
}
