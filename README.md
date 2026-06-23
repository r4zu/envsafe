# 🔒 envsafe

A secure, collaborative, and zero-dependency CLI tool to manage your `.env` files.

`envsafe` helps you sync `.env` keys with your `.env.example` templates, detect missing variables, encrypt secrets for safe storage in Git, and prevents you from accidentally committing raw credentials.

---

## ✨ Features

- 🔄 **Auto-Sync:** Extract `.env` keys (without values) to update your `.env.example` file automatically. Supports custom env files (e.g. `.env.local` -> `.env.local.example`).
- ⚠️ **Check & Auto-add:** Compare `.env` with `.env.example` to find missing keys and add them interactively from the command line with default options.
- 🔒 **AES-256-GCM Encryption:** Encrypt your `.env` to `.env.enc` to safely commit it to version control. Supports standard decryption using environment variables for CI/CD pipelines.
- 🚀 **Git Hooks Integration:** Installs a pre-commit hook that automatically blocks raw secrets from being committed and verifies your `.gitignore` configuration.
- 🎨 **Beautiful UI:** Styled console output with emoji indicators and ANSI colors.
- ⚡ **Zero Dependencies:** Written in pure Go with no runtime dependencies.

---

## 📦 Installation

To download the binary directly, grab the latest release for your platform from the Releases page, or install it using Go:

```bash
go install github.com/r4zu/envsafe@latest
```

---

## 🚀 Quick Start

### 1. Synchronize Example templates
Generate or update your `.env.example` template:
```bash
envsafe sync
# Or for framework-specific files:
envsafe sync .env.local
```

### 2. Verify variables
Check if your local configuration has all the keys defined in the example template, and add missing ones interactively:
```bash
envsafe check
```

### 3. Encrypt your secrets
Encrypt your raw environment file into an encrypted, shareable `.enc` version:
```bash
envsafe encrypt
```
And decrypt it back whenever needed:
```bash
envsafe decrypt
```

### 4. Install Git Hook protection
Prevent staging/committing raw `.env` files by installing the pre-commit hook:
```bash
envsafe hook install
```

---

## 🔒 Security

`envsafe` uses industry-standard **AES-256-GCM** (Galois/Counter Mode) encryption to protect your secrets. For automated environments (like CI/CD), you can set the `ENVSAFE_KEY` environment variable instead of using the interactive password prompt:

```bash
export ENVSAFE_KEY="your-strong-passphrase"
envsafe decrypt
```

---

## 📄 License

Distributed under the MIT License. See [LICENSE](LICENSE) for more information.
