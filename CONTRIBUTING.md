# Contributing to envsafe

Thank you for your interest in contributing to `envsafe`! Contributions from the community help make this tool better for everyone.

Here are some guidelines to help you get started.

## 🚀 Getting Started

1. **Fork the repository** on GitHub.
2. **Clone your fork** locally:
   ```bash
   git clone https://github.com/YOUR_USERNAME/envsafe.git
   cd envsafe
   ```
3. Ensure you have **Go 1.18 or higher** installed.

## 🛠️ Development Workflow

### Coding Standards
- Follow standard Go formatting guidelines by running:
  ```bash
  go fmt ./...
  ```
- Keep codebase structure clean: command routing goes into `main.go`, core logic goes into packages inside the `internal/` directory.

### Running Tests
Before submitting any changes, make sure all unit tests compile and pass:
```bash
go test -v ./...
```
If you are adding a new feature, please write corresponding test files under the `internal/` directory matching the suffix `_test.go`.

## 📥 Submitting a Pull Request

1. Create a new branch for your feature or bugfix:
   ```bash
   git checkout -b feature/your-feature-name
   ```
2. Commit your changes with clear, descriptive commit messages.
3. Push the branch to your fork:
   ```bash
   git push origin feature/your-feature-name
   ```
4. Open a **Pull Request** against the `main` branch of the original repository.
5. Provide a clear description of the changes and references to any open issues.

Thank you for contributing!
