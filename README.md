# Project col-moda

One Paragraph of project description goes here

## 📂 Project Structure

### 📂 `cmd/api`
- **main.go**: Entry point of the application where the server is initialized.

### 📂 `internal`
Contains the core application logic and is structured as follows:

#### 📂 `controllers`
- ***.controller.go**: Handles HTTP requests and responses.

#### 📂 `database`
- **connection.go**: Manages database connections and configurations.

#### 📂 `domain`
Contains domain-specific logic and data structures:
- **dtos/**: Data Transfer Objects for request/response handling.
- **entities/**: Core business entities/models.
- **models/**: Database models used for ORM.

#### 📂 `repositories`
- **repository.go**: Defines the base repository structure.
- ***.repository.go**: Manages database queries.

#### 📂 `server`
- **server.go**: Initializes and configures the Gin server.
- **v1/routes.go**: Defines API endpoints for version 1.
- **v1/routes/*.go**: Contains group routes.
- **v1/routes/*.test.go**: Contains tests for API routes.

#### 📂 `services`
- ***.service.go**: Implements business logic.

#### 📂 `utils`
- **jwt.go**: Handles JWT-based authentication.

---

## 📌 Configuration Files

- **.air.toml**: Configuration file for live-reloading with `Air`.
- **.env**: Stores environment variables (e.g., database credentials, secrets).
- **.gitignore**: Specifies ignored files for Git.
- **go.mod / go.sum**: Manage dependencies using Go modules.
- **Makefile**: Provides shortcuts for building, testing, and running the application.
- **README.md**: Documentation for the project.

---

## 📌 Environment Variables

The application requires the following environment variables:

```env
PORT=8080
APP_ENV=local
JWT_SECRET="SECRET"
DSN="host=localhost user=admin password=123 dbname=test port=5433 sslmode=disable"
```

- `PORT`: Defines the port on which the server runs.
- `APP_ENV`: Specifies the environment (e.g., local, production).
- `JWT_SECRET`: Secret key for JWT authentication.
- `DSN`: Database connection string.

---

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## MakeFile

Run build make command with tests
```bash
make all
```

Build the application
```bash
make build
```

Run the application
```bash
make run
```

Live reload the application:
```bash
make watch
```

Run the test suite:
```bash
make test
```

Clean up binary from the last build:
```bash
make clean
```