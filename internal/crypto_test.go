package internal

import (
	"os"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	// Set the environment variable for testing
	testKey := "my-super-secret-test-passphrase-123!"
	os.Setenv("ENVSAFE_KEY", testKey)
	defer os.Unsetenv("ENVSAFE_KEY")

	// 1. Setup: Create a temporary test env file
	originalContent := "SECRET_TOKEN=superSecretTokenValue123\nDB_PASS=mypass123\n"
	tmpFile := "crypto_test.env"
	err := os.WriteFile(tmpFile, []byte(originalContent), 0644)
	if err != nil {
		t.Fatalf("failed to create test env file: %v", err)
	}
	defer func() {
		os.Remove(tmpFile)
		os.Remove(tmpFile + ".enc")
	}()

	// 2. Execute Encryption
	err = EncryptEnv(tmpFile)
	if err != nil {
		t.Fatalf("encryption failed: %v", err)
	}

	// Check that the encrypted file was created
	if _, err := os.Stat(tmpFile + ".enc"); os.IsNotExist(err) {
		t.Fatalf("encrypted file was not created")
	}

	// Remove the plain file to ensure decryption actually restores it
	os.Remove(tmpFile)

	// 3. Execute Decryption
	err = DecryptEnv(tmpFile)
	if err != nil {
		t.Fatalf("decryption failed: %v", err)
	}

	// 4. Verify: Read decrypted file and compare with original content
	decryptedBytes, err := os.ReadFile(tmpFile)
	if err != nil {
		t.Fatalf("failed to read decrypted file: %v", err)
	}
	decryptedContent := string(decryptedBytes)

	if decryptedContent != originalContent {
		t.Errorf("decrypted content does not match original.\nExpected: %q\nGot: %q", originalContent, decryptedContent)
	}
}
