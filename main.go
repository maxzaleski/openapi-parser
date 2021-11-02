package main

import (
	"fmt"
	"io"
	"log"
	"os"

	openapi_parser "openapi-gen/openapi-parser"
)

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

	if false {
		fmt.Println(doc)
	}
}
