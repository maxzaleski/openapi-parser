package output

import (
	"errors"
	"os"
	"strings"

	"openapi-gen/internal/slog"
)

// CreateFiles will create the given files.
func CreateFiles(outDir string, files []*File, extn string, logger slog.Logger) error {
	logger.SetPrefix("[CreateFiles] ")
	logger.Println("Generating output files...")

	// Append path suffix if missing.
	if outDir != "" && !strings.HasSuffix(outDir, "/") {
		outDir += "/"
	}
	for _, fileToBeCreated := range files {
		// Append path suffix if missing.
		if fileToBeCreated.Directory != "" && !strings.HasSuffix(fileToBeCreated.Directory, "/") {
			fileToBeCreated.Directory += "/"
		}
		// Create the file's directory if it doesn't exist.
		if err := os.MkdirAll(outDir+fileToBeCreated.Directory, 0755); err != nil {
			logger.Println("os.MkdirAll:", err)
			return err
		}

		filePath := outDir + fileToBeCreated.Directory + fileToBeCreated.Name + extn
		logger.Printf("Seen '%s'", filePath)

		// Attempt to open the file; otherwise, create it.
		f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				logger.Println("os.OpenFile:", err)
				return err
			}
			f, err = os.Create(filePath)
			if err != nil {
				logger.Println("os.Create:", err)
				return err
			}
		}
		if _, err = f.WriteString(fileToBeCreated.Body); err != nil {
			return err
		}

		logger.Printf("Created '%s'", filePath)
		_ = f.Close()
	}

	return nil
}
