# Personal Budget App

Personal Budget App is a project designed for learning and testing various technologies. It provides functionality to explore, implement, and experiment with features such as creating, retrieving, updating, and deleting budget entries. The application is built using Go, PostgreSQL

## Features

- Create, view, update, and delete budget entries.
- PostgreSQL database integration.
- REST API built with [Gin Gonic](https://gin-gonic.com/).
- Swagger documentation for API endpoints.
- Dockerized environment for ease of development and deployment.
- CI/CD pipeline with GitHub Actions.
- Linter and security checks integrated into the development workflow.

## Table of Contents

- [Getting Started](#getting-started)
- [Development](#development)
- [Testing](#testing)
- [API Documentation](#api-documentation)
- [Environment Variables](#environment-variables)
- [Contributing](#contributing)
- [License](#license)

## Getting Started

### Prerequisites

- [Go 1.23+](https://go.dev/)
- [Docker](https://www.docker.com/)
- [PostgreSQL](https://www.postgresql.org/)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/PatrykEL01/personal-budget-app.git
   cd personal-budget-app
   ```

2. Set up the environment variables (see [Environment Variables](#environment-variables)).

3. Start the application locally:
   ```bash
   go run main.go
   ```

4. Alternatively, use Docker:
   ```bash
   docker build -t personal-budget-app .
   docker run -p 8080:8080 --env-file .env personal-budget-app
   ```

## Development

### Run the Application

1. Ensure PostgreSQL is running locally or via Docker:
   ```bash
   docker run --name postgres -e POSTGRES_USER=testuser -e POSTGRES_PASSWORD=testpassword -e POSTGRES_DB=testdb -p 5432:5432 -d postgres:15
   ```

2. Run the app:
   ```bash
   go run main.go
   ```

### Linters and Code Formatting

Ensure your code adheres to Go best practices:
```bash
gofmt -s -w .
go vet ./...
golangci-lint run
```

## Testing

Run the tests with:
```bash
go test ./... -v
```

For integration tests, ensure PostgreSQL is available and seeded with test data.

## API Documentation

Swagger documentation is available for all endpoints. To access the documentation:

1. Start the application.
2. Open your browser and navigate to: `http://localhost:8080/swagger/index.html`.

## Environment Variables

The application requires the following environment variables:

| Variable       | Description                         | Example                                 |
|----------------|-------------------------------------|-----------------------------------------|
| `DATABASE_URL` | Connection string to PostgreSQL DB | `postgres://testuser:testpassword@localhost:5432/testdb` |

## CI/CD Pipeline

This repository uses GitHub Actions for continuous integration and deployment:

- **Linting and Code Quality**: Runs `gofmt`, `go vet`, and other linters.
- **Testing**: Executes unit and integration tests.
- **Docker Build and Scan**: Builds the Docker image and scans it for vulnerabilities using Trivy.



## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

