package internal

import (
	"os"
	"strings"
	"testing"
)

func TestSyncEnv(t *testing.T) {
	// 1. Setup: Create a temporary test env file
	testEnvContent := `# Database Config                                                                                                                       
    DB_PORT=5432                                                                                                                                                 
    DB_HOST=localhost                                                                                                                                            
                                                                                                                                                                 
    # Keys                                                                                                                                                       
    API_KEY=xyz123                                                                                                                                               
    `
	tmpEnvFile := "test.env"
	err := os.WriteFile(tmpEnvFile, []byte(testEnvContent), 0644)
	if err != nil {
		t.Fatalf("failed to create test env file: %v", err)
	}
	defer func() {
		os.Remove(tmpEnvFile)
		os.Remove(tmpEnvFile + ".example")
	}()

	// 2. Execute: Run SyncEnv on our temp file
	err = SyncEnv(tmpEnvFile)
	if err != nil {
		t.Fatalf("SyncEnv returned an error: %v", err)
	}

	// 3. Verify: Read the generated .example file
	exampleBytes, err := os.ReadFile(tmpEnvFile + ".example")
	if err != nil {
		t.Fatalf("failed to read generated example file: %v", err)
	}
	exampleContent := string(exampleBytes)

	// Check if keys exist but values are stripped
	if !strings.Contains(exampleContent, "DB_PORT=") || strings.Contains(exampleContent, "5432") {
		t.Errorf("DB_PORT value was not stripped correctly. Got: %s", exampleContent)
	}
	if !strings.Contains(exampleContent, "DB_HOST=") || strings.Contains(exampleContent, "localhost") {
		t.Errorf("DB_HOST value was not stripped correctly")
	}
	if !strings.Contains(exampleContent, "API_KEY=") || strings.Contains(exampleContent, "xyz123") {
		t.Errorf("API_KEY value was not stripped correctly")
	}

	// Verify comments are preserved
	if !strings.Contains(exampleContent, "# Database Config") {
		t.Errorf("comments were not preserved in the example file")
	}
}
