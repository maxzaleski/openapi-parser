package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"openapi-gen/gen"
)

func main() {
	if err := exec(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Done.")
}

func exec() error {
	var extnFlag string
	{
		flag.StringVar(&extnFlag, "extension", "", "Extension to use for output files")
		flag.Parse()

		if extnFlag == "" {
			return fmt.Errorf("extension must be specified")
		} else if !strings.HasPrefix(extnFlag, ".") {
			extnFlag = "." + extnFlag
		}
	}

	specFile, err := os.Open("../boardinghub-api.spec.yaml")
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

	return gen.New(b, gen.Extension(extnFlag))
}
