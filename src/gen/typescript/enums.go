package typescript

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"

	"openapi-generator/internal/parser"

	"openapi-generator/gen/typescript/templates"
)

// generateEnum generates a typescript enum from the given definition.
func generateEnum(def *parser.Definition) string {
	template := templates.Enum

	// Enum description.
	if defDesc := def.Description; defDesc != "" {
		template = toJSDoc("", defDesc) + template
	}
	// Enum entries.
	mappedEntries := make([]string, 0, len(def.EnumEntries))
	for _, entry := range def.EnumEntries {
		if entry == "" {
			continue
		}
		switch def.Key {
		case "Colour":
			mappedEntries = append(mappedEntries, generateEnumColourProperty(entry))
		default:
			mappedEntries = append(mappedEntries,
				fmt.Sprintf(templates.EnumStringProperty, strcase.ToScreamingSnake(entry), entry))
		}
	}

	return fmt.Sprintf(template, def.Key, strings.Join(mappedEntries, "\n"))
}

// generateEnumColourProperty generates a typescript enum property for the given colour.
func generateEnumColourProperty(entry string) string {
	colourName := ""
	switch entry {
	case "1":
		colourName = "AMBER"
	case "2":
		colourName = "ORANGE"
	case "3":
		colourName = "TOMATO"
	case "4":
		colourName = "RED"
	case "5":
		colourName = "CRIMSON"
	case "6":
		colourName = "PINK"
	case "7":
		colourName = "PLUM"
	case "8":
		colourName = "PURPLE"
	case "9":
		colourName = "VIOLET"
	case "10":
		colourName = "INDIGO"
	case "11":
		colourName = "BLUE"
	case "12":
		colourName = "CYAN"
	case "13":
		colourName = "TEAL"
	case "14":
		colourName = "GREEN"
	case "15":
		colourName = "GRASS"
	}
	return fmt.Sprintf(templates.EnumNumberProperty, colourName, entry)
}
