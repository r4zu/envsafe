package main

import (
	"fmt"
	"os"

	"github.com/r4zu/envsafe/internal"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()

		return
	}

	command := os.Args[1]

	switch command {
	case "sync":
		sourceFile := ".env"
		if len(os.Args) >= 3 {
			sourceFile = os.Args[2]
		}

		fmt.Printf("Syncing %s with %s.example...\n", sourceFile, sourceFile)

		if err := internal.SyncEnv(sourceFile); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Successfully created/updated %s.example!\n", sourceFile)

	case "check":
		sourceFile := ".env"
		if len(os.Args) >= 3 {
			sourceFile = os.Args[2]
		}
		if err := internal.CheckEnv(sourceFile); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

	case "encrypt":
		sourceFile := ".env"
		if len(os.Args) >= 3 {
			sourceFile = os.Args[2]
		}
		if err := internal.EncryptEnv(sourceFile); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

	case "decrypt":
		sourceFile := ".env"
		if len(os.Args) >= 3 {
			sourceFile = os.Args[2]
		}
		if err := internal.DecryptEnv(sourceFile); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

	case "hook":
		if len(os.Args) > 3 {
			fmt.Println("Usage: envsafe hook [install | uninstall | verify ]")
			os.Exit(1)
		}
		subCommand := os.Args[2]
		switch subCommand {
		case "install":
			if err := internal.InstallHook(); err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
		case "uninstall":
			if err := internal.UninstallHook(); err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
		case "verify":
			if err := internal.VerifyCommit(); err != nil {
				os.Exit(1)
			}
		default:
			fmt.Printf("Error: unknown hook command '%s'\n", subCommand)
			os.Exit(1)
		}

	case "help", "-h", "--help":
		printHelp()

	default:
		fmt.Printf("Error: unknown command '%s'\n\n", command)

		printHelp()

		os.Exit(1)
	}
}

func printHelp() {
	helpText := `envsafe - Secure and collaborative .env file manager

Usage:
	envsafe <command> [arguments]

Commands:
    sync        Generate or update .env.example from your .env (without values)
    check       Compare your .env with .env.example to find missing variables
    encrypt     Encrypt your .env file to .env.enc
    decrypt     Decrypt .env.enc back to .env
    hook        Install or uninstall the git pre-commit hook
    help        Show this help text

Run a command to get started.`
	fmt.Println(helpText)
}
