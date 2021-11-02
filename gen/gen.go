package gen

import (
	"errors"
	"log"
	"os"

	"openapi-gen/gen/parser"
)

const outDir = "tmp"

func New(b []byte, extn Extension) error {
	logger := log.New(os.Stdout, "[gen] ", 0)

	doc, err := parser.NewDocument(b)
	if err != nil {
		logger.Println("parser.NewDocument: ", err)
		return err
	}

	// Create directory if it doesn't exist.
	if err = os.MkdirAll(outDir, 0755); err != nil {
		logger.Println("os.MkdirAll", err)
		return err
	}

	fileName := outDir + "/api" + extn.String()
	var f *os.File
	f, err = os.Open(fileName)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			logger.Println("os.Open: ", err)
			return err
		}
		f, err = os.Create(fileName)
		if err != nil {
			logger.Println("os.Create: ", err)
			return err
		}
	}
	defer func() {
		_ = f.Close()
	}()

	generated := generateFromTemplate(doc, extn, logger)
	if _, err = f.WriteString(generated); err != nil {
		logger.Println("file.WriteString: ", err)
		return err
	}

	return nil
}

// Extension represents a file extension.
type Extension string

func (e Extension) String() string {
	return string(e)
}

const (
	ExtensionTypescript Extension = ".ts"
)
