package main

import (
	"io"
	"os"

	"openapi-gen/gen"
)

func main() {
	if err := exec(); err != nil {
		os.Exit(1)
	}
}

func exec() error {
	specFile, err := os.Open("./spec.yml")
	if err != nil {
		return err
	}
	defer func() {
		_ = specFile.Close()
	}()
	b, err := io.ReadAll(specFile)
	if err != nil {
		return err
	}

	return gen.New(b, ".ts")
}
