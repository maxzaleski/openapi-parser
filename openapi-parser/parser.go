package openapi_parser

import (
	"gopkg.in/yaml.v2"
)

// NewDocument returns a new instance of `Document`.
func NewDocument(b []byte) (*Document, error) {
	var doc preParsedDocument
	if err := yaml.Unmarshal(b, &doc); err != nil {
		return nil, err
	}
	return &Document{
		Meta:        doc.Meta,
		Hosts:       doc.Hosts,
		Definitions: parseIntoDefinitions(doc.Definitions),
		Responses:   parseIntoResponses(doc.Responses),
	}, nil
}
