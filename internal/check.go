package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getEnvKeys(filePath string) (map[string]bool, error) {
	keys := make(map[string]bool)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) >= 1 {
			key := strings.TrimSpace(parts[0])
			if key != "" {
				keys[key] = true
			}
		}
	}

	return keys, scanner.Err()
}

func CheckEnv(sourcePath string) error {
	if sourcePath == "" {
		sourcePath = ".env"
	}
	targetPath := sourcePath + ".example"

	if _, err := os.Stat(sourcePath); os.IsNotExist(err) {
		return fmt.Errorf("local file '%s' not found", sourcePath)
	}
	if _, err := os.Stat(targetPath); os.IsNotExist(err) {
		return fmt.Errorf("example file '%s' not found. Run 'sync' to generate it first", targetPath)
	}

	sourceKeys, err := getEnvKeys(sourcePath)
	if err != nil {
		return fmt.Errorf("failed to read %s: %w", sourcePath, err)
	}

	exampleKeys, err := getEnvKeys(targetPath)
	if err != nil {
		return fmt.Errorf("failed to read %s: %w", targetPath, err)
	}

	var missingKeys []string
	for key := range exampleKeys {
		if !sourceKeys[key] {
			missingKeys = append(missingKeys, key)
		}
	}

	if len(missingKeys) > 0 {
		fmt.Printf("\n⚠️  %s: The following keys are defined in %s but missing from %s:\n",
			Bold(Yellow("Warning")),
			Cyan(targetPath),
			Cyan(sourcePath),
		)
		for _, key := range missingKeys {
			fmt.Printf("   - %s\n", Bold(key))
		}

		err := promptAndAddKeys(sourcePath, missingKeys)
		if err != nil {
			return err
		}
		return nil
	}

	fmt.Printf("✅ %s: All keys from %s are present in %s!\n",
		Bold(Green("Success")),
		Cyan(targetPath),
		Cyan(sourcePath),
	)
	return nil
}

func promptAndAddKeys(sourcePath string, missingKeys []string) error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\nDo you want to add these missing keys to %s? (Y/n): ", Cyan(sourcePath))

	response, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	response = strings.ToLower(strings.TrimSpace(response))
	if response == "" {
		response = "y"
	}

	if response != "y" && response != "yes" {
		fmt.Println(Yellow("Skipped adding keys."))
		return fmt.Errorf("configuration is incomplete")
	}

	file, err := os.OpenFile(sourcePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open %s for writing: %w", sourcePath, err)
	}
	defer file.Close()

	file.WriteString("\n")

	for _, key := range missingKeys {
		fmt.Printf("Enter value for %s: ", Bold(Cyan(key)))
		val, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		val = strings.TrimSpace(val)

		_, err = file.WriteString(fmt.Sprintf("%s=%s\n", key, val))
		if err != nil {
			return fmt.Errorf("failed to write key %s: %w", key, err)
		}
	}

	fmt.Println(Green("✅ Keys successfully added to your local file!"))
	return nil
}
