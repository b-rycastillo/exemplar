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

The project includes GitHub Actions workflows for automated testing and deployment across three environments:

- **Development (development)** - Runs on every push to develop branch
- **Testing (test)** - Runs on successful completion in dev, tests all changes
- **Production (master)** - Runs on successful completion in test, deploys to production

### Workflows

#### 1. Development Pipeline (.github/workflows/dev.yml)
- Triggers on: Push to `develop` branch
- Runs: Unit tests, integration tests
- On success: Triggers test workflow

#### 2. Testing Pipeline (.github/workflows/test.yml)
- Triggers on: Successful completion of dev pipeline
- Runs: Full test suite, code quality checks
- On success: Triggers production workflow

#### 3. Production Pipeline (.github/workflows/main.yml)
- Triggers on: Successful completion of test pipeline
- Runs: Final verification, builds release binary
- On success: Deploys to production

### Setting Up CI/CD

1. **Push to develop branch:**
```bash
git checkout develop
git add .
git commit -m "New feature"
git push origin develop
```

2. **Monitor CI/CD:**
Visit the "Actions" tab in your GitHub repository to see workflow progress.

3. **Merge to main:**
Once all workflows pass, your code is automatically merged to main and deployed.

### Environment Variables

Configure these in GitHub repository settings under Secrets:

- `DEPLOY_KEY` - SSH key for production deployment
- `SERVER_ADDR` - Production server address
- `API_PORT` - Port for production API

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
