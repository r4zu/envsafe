package internal

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"strings"
)

func deriveKey(passphrase string) []byte {
	hash := sha256.Sum256([]byte(passphrase))
	return hash[:]
}

func getPassphrase(prompt string) (string, error) {
	if key := os.Getenv("ENVSAFE_KEY"); key != "" {
		return key, nil
	}

	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	passphrase, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	passphrase = strings.TrimSpace(passphrase)
	if passphrase == "" {
		return "", fmt.Errorf("passphrase cannot be empty")
	}
	return passphrase, nil
}

func EncryptEnv(sourcePath string) error {
	if sourcePath == "" {
		sourcePath = ".env"
	}
	targetPath := sourcePath + ".enc"

	plaintext, err := os.ReadFile(sourcePath)
	if err != nil {
		return fmt.Errorf("failed to read %s: %w", sourcePath, err)
	}

	passphrase, err := getPassphrase(fmt.Sprintf("Enter passphrase to encrypt %s: ", Cyan(sourcePath)))
	if err != nil {
		return err
	}

	key := deriveKey(passphrase)

	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("failed to create cipher block: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return fmt.Errorf("failed to create GCM: %w", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return fmt.Errorf("failed to generate random nonce: %w", err)
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

	err = os.WriteFile(targetPath, ciphertext, 0600)
	if err != nil {
		return fmt.Errorf("failed to write encrypted file: %w", err)
	}

	fmt.Printf("🔒 %s! File encrypted and saved to %s\n", Bold(Green("Success")), Cyan(targetPath))
	return nil
}

func DecryptEnv(targetPath string) error {
	if targetPath == "" {
		targetPath = ".env"
	}
	sourcePath := targetPath + ".enc"

	if _, err := os.Stat(sourcePath); os.IsNotExist(err) {
		return fmt.Errorf("encrypted file '%s' not found", sourcePath)
	}

	ciphertext, err := os.ReadFile(sourcePath)
	if err != nil {
		return fmt.Errorf("failed to read %s: %w", sourcePath, err)
	}

	passphrase, err := getPassphrase(fmt.Sprintf("Enter passphrase to decrypt %s: ", Cyan(sourcePath)))
	if err != nil {
		return err
	}

	key := deriveKey(passphrase)

	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("failed to create cipher block: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return fmt.Errorf("failed to create GCM: %w", err)
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return fmt.Errorf("ciphertext is too short/corrupted")
	}

	nonce, ciphertextActual := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertextActual, nil)
	if err != nil {
		return fmt.Errorf("decryption failed: incorrect passphrase or corrupted file")
	}

	err = os.WriteFile(targetPath, plaintext, 0644)
	if err != nil {
		return fmt.Errorf("failed to write decrypted file: %w", err)
	}

	fmt.Printf("🔓 %s! Decrypted configuration written to %s\n", Bold(Green("Success")), Cyan(targetPath))
	return nil
}
