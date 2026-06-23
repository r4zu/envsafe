# 🔒 envsafe

A secure, collaborative, and zero-dependency CLI tool to manage your `.env` files.

`envsafe` helps you sync `.env` keys with your `.env.example` templates, detect missing variables, encrypt secrets for safe storage in git, and prevents you
from accidentally committing raw credentials.

---

## ✨ Features

- 🔄 **Auto-Sync:** Extract `.env` keys (without values) to update your `.env.example` file automatically.
- ⚠️ **Check Variables:** Compare `.env` with `.env.example` to find missing keys and add them interactively.
- 🔒 **AES-256-GCM Encryption:** Encrypt your `.env` to `.env.enc` to safely commit it to version control.
- 🚀 **Git Hooks Integration:** Installs a pre-commit hook that automatically blocks raw secrets from being committed.
- 🎨 **Beautiful UI:** Styled console output with emoji indicators and ANSI colors.  


---

## 📦 Installation

To download the binary directly, grab the latest release for your platform from the Releases page, or install it using Go:

```bash
go install github.com/r4zu/envsafe@latest
──────
## 🚀 Quick Start

### 1. Synchronize Example templates

Generate or update your  .env.example  template:

envsafe sync
# Or for framework-specific files:
envsafe sync .env.local

### 2. Verify variables

Check if your local configuration has all the keys defined in the example template:

envsafe check

### 3. Encrypt your secrets

Safely commit encrypted credentials (useful for sharing within team members or CI/CD):

envsafe encrypt

### 4. Install Git Hook protection

Prevent staging/committing raw  .env  files by installing the pre-commit hook:

envsafe hook install
```
