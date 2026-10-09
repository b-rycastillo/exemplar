# Contributing to Exemplar

Thank you for your interest in contributing to Exemplar! This document provides guidelines and instructions for contributing to the project.

## Code of Conduct

We are committed to providing a welcoming and inspiring community for all. Please read and adhere to our Code of Conduct.

## Getting Started

1. Fork the repository
2. Clone your fork: `git clone https://github.com/YOUR_USERNAME/exemplar.git`
3. Create a new branch: `git checkout -b feature/your-feature-name`
4. Make your changes
5. Test your changes: `make test`
6. Commit with clear messages: `git commit -m "Add your feature"`
7. Push to your fork: `git push origin feature/your-feature-name`
8. Create a Pull Request

## Development Setup

```bash
# Clone the repository
git clone https://github.com/example/exemplar.git
cd exemplar

# Download dependencies
make deps

# Run tests
make test

# Build the application
make build
```

## Coding Standards

### Go Style Guide

- Follow the [Effective Go](https://golang.org/doc/effective_go) guidelines
- Use `gofmt` for code formatting: `make format`
- Run `go vet` to check for common errors: `make vet`
- Use meaningful variable names
- Add comments for exported functions

### Code Review Checklist

Before submitting a PR, ensure:

- [ ] Code follows Go style guidelines
- [ ] All tests pass: `make test`
- [ ] New features have tests
- [ ] Code coverage doesn't decrease
- [ ] Documentation is updated
- [ ] No hardcoded credentials or secrets
- [ ] Commit messages are clear and descriptive

## Testing Requirements

### Unit Tests

All business logic must have unit tests:

```bash
make test-unit
```

### Integration Tests

All API endpoints must have integration tests:

```bash
make test-integration
```

### Test Coverage

Maintain at least 70% code coverage:

```bash
make test-coverage
```

### Race Detection

Run tests with race detector enabled:

```bash
make test-race
```

## Commit Message Guidelines

Format commit messages as follows:

```
[Type] Brief description (50 chars max)

Longer explanation if needed. Wrap at 72 characters.
- Point 1
- Point 2

Closes #123
```

Types:
- `feat:` New feature
- `fix:` Bug fix
- `docs:` Documentation changes
- `style:` Code style changes (formatting, missing semicolons, etc.)
- `refactor:` Code refactoring without feature/fix changes
- `test:` Adding or updating tests
- `chore:` Build process, dependencies, etc.

## Pull Request Process

1. Update the README.md with details of changes if applicable
2. Update the CHANGELOG.md with notes on your changes
3. Increase version numbers following [Semantic Versioning](https://semver.org/)
4. Ensure all tests pass
5. Request review from maintainers
6. Address review comments
7. After approval, your PR will be merged automatically by CI/CD

## Reporting Issues

When reporting bugs, please include:

- Clear description of the issue
- Steps to reproduce
- Expected behavior
- Actual behavior
- Go version: `go version`
- Operating system and version
- Error messages and stack traces
- Code examples if applicable

## Feature Requests

Include:

- Clear description of the feature
- Use cases and benefits
- Example API or behavior
- Any potential breaking changes

## Documentation

- Update README.md for major changes
- Add code comments for complex logic
- Document exported functions and types
- Update API documentation for endpoint changes

## Questions?

- Open a GitHub issue with the `question` label
- Check existing issues and discussions
- Review the README and documentation

## License

By contributing to Exemplar, you agree that your contributions will be licensed under the MIT License.

---

Thank you for contributing to Exemplar!
