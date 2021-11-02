package main

import (
	"io"
	"os"

	"openapi-gen/gen"
)

func main() {
	if err := main_(); err != nil {
		os.Exit(1)
	}
}

func main_() error {
	specFile, err := os.Open("./spec.yml")
	if err != nil {
		return err
	}
	b, err := io.ReadAll(specFile)
	if err != nil {
		return err
	}

	return gen.New(b, ".ts")
}
