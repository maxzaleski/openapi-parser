package main

import (
	"flag"
	"fmt"
	"io"
	"openapi-generator/gen"
	"os"
	"strings"
)

const VERSION = "0.1.0"

func main() {
	// README:
	// ---
	// Q: Why not simply include the execution code inside `main`?:
	// A: The execution code is contained inside a func (`run`) to enable a single point of return. By doing so,
	// we are able to define a centralised error/success handling pattern.
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1) // Exit with error.
	}
	fmt.Println("Done.")
}

const specFileName = ""

func run() error {
	// Parse file extension flag.
	var extnFlag string
	{
		flag.StringVar(&extnFlag, "extension", "", "Extension to use for output files")
		flag.Parse()

		if extnFlag == "" {
			return fmt.Errorf("error: an extension must be specified")
		} else if !strings.HasPrefix(extnFlag, ".") {
			extnFlag = "." + extnFlag
		}
	}

	// Open and read specification file.
	specFile, err := os.Open("../../openapi/" + specFileName)
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
	// Generate.
	return gen.New(b, VERSION, gen.Extension(extnFlag))
}
