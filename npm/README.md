# 🔒 envsafe (NPM Wrapper)

A secure, collaborative, and zero-dependency CLI tool to manage your `.env` files.

This is the NPM wrapper package for `envsafe`, allowing Javascript/Node.js developers to run `envsafe` via `npx` or install it globally with `npm`.

---

## 📦 Installation

To install globally:
```bash
npm install -g @r4zu/envsafe
```

Or run it instantly without installing:
```bash
npx @r4zu/envsafe <command>
```

---

## 🚀 Quick Start

### 1. Sync Example templates
Generate or update your `.env.example` template:
```bash
npx @r4zu/envsafe sync
```

### 2. Verify variables
Check if your local configuration matches the example template and add missing ones:
```bash
npx @r4zu/envsafe check
```

### 3. Encrypt secrets
Safely encrypt your credentials into a shareable encrypted format:
```bash
npx @r4zu/envsafe encrypt
```

### 4. Git hooks setup
Prevent committing raw `.env` credentials:
```bash
npx @r4zu/envsafe hook install
```

---

## 📄 Documentation

For the full documentation and Go source code, visit the main repository:
[github.com/r4zu/envsafe](https://github.com/r4zu/envsafe)
