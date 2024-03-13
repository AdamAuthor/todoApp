# Go Todo List REST API
This is a sample project showcasing the development of a Todo List REST API in Go. It follows the principles of Clean Architecture and utilizes the gin-gonic/gin framework for web development. The application interacts with a PostgreSQL database, runs within Docker containers, and utilizes migration files for managing database schema changes.

## Features
- Clean Architecture: The project follows the Clean Architecture principles, separating concerns into layers for better maintainability and testability.
- Dependency Injection: Dependency Injection is employed to manage component dependencies, making the application more modular and easier to test.
- PostgreSQL Database: Data is stored in a PostgreSQL database, and the application uses the sqlx library for database interaction.
- Configuration Management: Application configuration is handled using the spf13/viper library, supporting environment variables for configuration flexibility.
- Authentication and Authorization: User registration, authentication, and JWT-based authorization are implemented to secure the API endpoints.
- Middleware: Middleware functions are utilized for tasks such as authentication and logging.
- SQL Queries: SQL queries are written to interact with the database efficiently.
- Graceful Shutdown: Graceful shutdown mechanisms are implemented to ensure the application shuts down cleanly.
- 
## Prerequisites

Before running the application, ensure you have the following installed:

- Go programming language (v1.15 or later)
- Docker
- PostgreSQL

## Setup
1. Clone the repository:

```bash
git clone <repository-url>
```
2. Navigate to the project directory:

```bash
cd <project-directory>
```
3. Build the application:

```bash
make build
```
4. Run database migrations:

```bash
make migrate
```

## Usage
To start the application, run:

```bash
make run
```
This command will build and start the application. Once running, you can interact with the API using HTTP requests.
