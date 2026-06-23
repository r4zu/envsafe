package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// syncEnv reads a source env file and generates/updates its .example equivalent.
func SyncEnv(sourcePath string) error {
	// If sourcePath is empty, default to ".env"
	if sourcePath == "" {
		sourcePath = ".env"
	}

	// The example file will match the source name + ".example" (e.g. "env.local.example")
	targetPath := sourcePath + ".example"

	// Check if source file exists
	if _, err := os.Stat(sourcePath); os.IsNotExist(err) {
		return fmt.Errorf(".env file not found in current directory. Please create it first")
	}

	// Open the source file
	inFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer inFile.Close()

	// Create or overwrite the destination example file
	outFile, err := os.Create(targetPath)
	if err != nil {
		return fmt.Errorf("failed to create example file: %w", err)
	}
	defer outFile.Close()

	scanner := bufio.NewScanner(inFile)
	writer := bufio.NewWriter(outFile)

	for scanner.Scan() {
		line := scanner.Text()

		trimmedLine := strings.TrimSpace(line)

		// Preserve empty lines and comments
		if trimmedLine == "" || strings.HasPrefix(trimmedLine, "#") {
			_, err := writer.WriteString(line + "\n")
			if err != nil {
				return err
			}

			continue
		}

		// Split the line at the first "=" sign
		parts := strings.SplitN(line, "=", 2)

		if len(parts) == 2 {
			key := parts[0]
			_, err := writer.WriteString(key + "=\n")
			if err != nil {
				return err
			}
		} else {
			_, err := writer.WriteString(line + "\n")
			if err != nil {
				return err
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading source file: %w", err)
	}

	err = writer.Flush()
	if err != nil {
		return fmt.Errorf("failed to save example file: %w", err)
	}

	return nil
}
