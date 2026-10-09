# Exemplar - Demo User Management API

A comprehensive example Go application demonstrating best practices for building scalable APIs with full test coverage, CI/CD integration, and multi-environment deployment.

## Features

- **User Management API** - Create, read, update, and delete users
- **Input Validation** - Email and password validation
- **Unit Tests** - Comprehensive unit test coverage for services and utilities
- **Integration Tests** - Full API endpoint integration tests
- **Continuous Deployment** - Automated deployment pipeline from development to production
- **Health Check Endpoint** - Service health monitoring
- **Professional Structure** - Clean code organization following Go best practices

## Project Structure

```
exemplar/
├── cmd/
│   └── main.go              # Application entry point
├── internal/
│   ├── handler/
│   │   ├── handler.go       # HTTP request handlers
│   │   └── handler_test.go  # Integration tests
│   └── service/
│       ├── service.go       # Business logic
│       ├── utils.go         # Utility functions
│       ├── service_test.go  # Unit tests for service
│       └── utils_test.go    # Unit tests for utils
├── .github/
│   └── workflows/           # GitHub Actions CI/CD pipelines
├── go.mod                   # Go module definition
├── README.md               # This file
└── Makefile               # Build commands
```

## Prerequisites

- Go 1.21 or higher
- Git
- Make (optional, for using Makefile commands)

## Installation

1. **Clone the repository:**
```bash
git clone https://github.com/example/exemplar.git
cd exemplar
```

2. **Download dependencies:**
```bash
go mod download
```

3. **Build the application:**
```bash
go build -o exemplar ./cmd
```

## Running the Application

### Using the built binary:
```bash
./exemplar
```

### Using go run:
```bash
go run ./cmd/main.go
```

### With custom port:
```bash
PORT=8080 go run ./cmd/main.go
```

The server will start on `http://localhost:3000` (or the specified PORT).

## API Endpoints

### Health Check
- **GET** `/health`
- Returns the server health status

```bash
curl http://localhost:3000/health
```

### Create User
- **POST** `/users`
- Creates a new user

```bash
curl -X POST http://localhost:3000/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }'
```

Requirements:
- `name`: Non-empty string
- `email`: Valid email format
- `password`: Minimum 8 characters

### Get All Users
- **GET** `/users`
- Returns list of all users

```bash
curl http://localhost:3000/users
```

### Get User by ID
- **GET** `/users/{id}`
- Returns a specific user

```bash
curl http://localhost:3000/users/user_123456789
```

### Update User
- **PATCH** `/users/{id}`
- Updates user information

```bash
curl -X PATCH http://localhost:3000/users/user_123456789 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Jane Doe",
    "email": "jane@example.com"
  }'
```

### Delete User
- **DELETE** `/users/{id}`
- Deletes a user

```bash
curl -X DELETE http://localhost:3000/users/user_123456789
```

## Testing

### Run All Tests
```bash
go test ./...
```

### Run Tests with Coverage
```bash
go test -cover ./...
```

### Generate Coverage Report
```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Run Unit Tests Only
```bash
go test ./internal/service/...
```

### Run Integration Tests Only
```bash
go test ./internal/handler/...
```

### Run Tests with Verbose Output
```bash
go test -v ./...
```

## Development

### Code Structure

**internal/service/:**
- `service.go` - UserService with CRUD operations
- `utils.go` - Email validation, password validation, and ID generation
- `*_test.go` - Comprehensive unit tests

**internal/handler/:**
- `handler.go` - HTTP request handlers and route setup
- `*_test.go` - Integration tests using httptest package

**cmd/:**
- `main.go` - Application entry point

### Running with Hot Reload

Install `air` for hot reloading during development:
```bash
go install github.com/cosmtrek/air@latest
air
```

## Continuous Deployment

The branch flow is `task-001 → development → test → master`.

1. Open a task branch PR targeting `development`. **Pull Request CI** (`.github/workflows/pr.yml`) runs only **Run Tests and Build**. Promotion jobs are absent from this workflow.
2. After review and successful checks, merge the PR into `development`. **Development CI and Promotion** (`.github/workflows/dev.yml`) validates the merged commit and fast-forwards `test` to that commit.
3. The reusable test workflow (`.github/workflows/test.yml`) runs coverage, API integration tests, a build, and a health check. Once these pass, it fast-forwards `master` to the same commit.
4. The reusable master workflow (`.github/workflows/master.yml`) runs final checks and publishes Linux, macOS, and Windows release binaries from that commit.

### Finding the testing and release jobs

The Actions tab shows a PR validation run and, after merging, a development push run. The testing and release workflows execute as nested jobs inside the development push run; they do not create separate runs for the `test` or `master` branches. Open that run to see **Validate Test Build**, **Promote Test to Master**, and **Build and Publish Release**.

This pipeline publishes GitHub Releases. A production server deployment requires an additional deployment step.

### Repository setup

- Create `development`, `test`, and `master` with shared history. Promotions use fast-forward merges and fail if the destination has diverged.
- Require **Run Tests and Build** for PRs into `development`. Task PRs are merged after review; the workflow promotes the resulting development commit automatically.
- Allow workflow contents writes and ensure branch rules permit the workflow actor to push to `test` and `master`.
- Restrict direct writes to `test` and `master` to the promotion process. The current pipeline uses `GITHUB_TOKEN` and needs no deployment secrets.
- Development push runs serialize the full promotion chain. New PR commits cancel outdated PR validation runs.

## Performance & Scalability

- In-memory user storage (for demo purposes)
- RESTful API design
- Fast JSON encoding/decoding with Go's standard library
- Gorilla mux for efficient routing
- Goroutine support for concurrent requests

For production use, consider:
- Database integration (PostgreSQL, MongoDB)
- Caching layer (Redis)
- Load balancing
- Containerization (Docker)

## Error Handling

The API returns appropriate HTTP status codes:

- `200 OK` - Successful GET request
- `201 Created` - User successfully created
- `400 Bad Request` - Invalid input or validation error
- `404 Not Found` - User not found
- `500 Internal Server Error` - Server error

All error responses follow the format:
```json
{
  "error": "Error message here"
}
```

## Validation Rules

### Email
- Must contain `@` symbol
- Must have domain extension (e.g., `.com`)
- No spaces allowed

### Password
- Minimum 8 characters required
- Case sensitive

### Name
- Cannot be empty
- Maximum length: unlimited (but recommended < 255 characters)

## Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/new-feature`
3. Make your changes and add tests
4. Ensure all tests pass: `go test ./...`
5. Commit your changes: `git commit -am 'Add new feature'`
6. Push to the branch: `git push origin feature/new-feature`
7. Submit a pull request

## Testing Requirements

Before submitting a PR:
- All tests must pass
- New features must include tests
- Code coverage should not decrease
- Follow Go naming conventions and best practices

## Deployment

### Local Development
```bash
go run ./cmd/main.go
```

### Docker Deployment
```bash
docker build -t exemplar .
docker run -p 3000:3000 exemplar
```

### Production
The CI/CD pipeline handles production deployment automatically when code is merged to main.

## Troubleshooting

### Build Errors
```bash
# Clean build cache
go clean -cache
go build ./...
```

### Test Failures
```bash
# Run tests with verbose output
go test -v ./...

# Check for race conditions
go test -race ./...
```

### Port Already in Use
```bash
# Use different port
PORT=8080 go run ./cmd/main.go

# Kill existing process (macOS/Linux)
lsof -i :3000 | grep LISTEN | awk '{print $2}' | xargs kill -9
```

## Support

For issues and questions:
- Check existing GitHub issues
- Create a new issue with detailed description
- Include error messages and steps to reproduce

## Changelog

### Version 1.0.0
- Initial release
- User CRUD operations
- Full test coverage
- CI/CD pipeline setup
- Health check endpoint

