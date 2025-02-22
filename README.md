# Go Clean Architecture Boilerplate with FastHTTP

A robust Go application boilerplate implementing Clean Architecture principles with FastHTTP framework.

## Project Structure

```txt
├── api                 # HTTP API layer
│   ├── controller     # Request handlers
│   ├── middleware     # HTTP middleware
│   └── route         # Route definitions
├── bootstrap          # Application bootstrapping
├── cmd               # Main applications
├── domain            # Enterprise business rules and entities
├── infrastructure    # External tools and libraries
├── pkg               # Common utilities and helpers
├── repository        # Data access layer
└── usecase           # Application business rules
```

## Architecture Overview

This project follows Clean Architecture principles with the following layers:

1. **Entities/Domain Layer** (`domain`)

   - Contains enterprise-wide business rules
   - Defines core business entities and interfaces
   - Has no dependencies on external layers

2. **Use Cases Layer** (`usecase`)

   - Contains application-specific business rules
   - Orchestrates data flow between entities
   - Implements business logic

3. **Interface Adapters Layer** (`api`, `repository`)

   - Contains adapters that convert data between use cases and external agencies
   - Implements repository interfaces
   - Handles HTTP routing and controllers

4. **Infrastructure Layer** (`infrastructure`, `pkg`)
   - Contains frameworks, tools, and helpers
   - Database implementations
   - External services integration

## Design Patterns Used

1. **Repository Pattern**

   - Abstracts data persistence operations
   - Defined in `repository`
   - Interfaces in `domain`

2. **Dependency Injection**

   - Used throughout the application
   - Dependencies are initialized in `bootstrap`
   - Promotes loose coupling and testability

3. **Factory Pattern**

   - Used in bootstrapping components
   - Creates instances of repositories and use cases

4. **Middleware Pattern**
   - HTTP middleware implementations
   - Request/Response processing

## Getting Started

### Prerequisites

- Go 1.23.4 or higher
- PostgreSQL
- Docker (optional)

### Environment Setup

```bash
cp .env.example .env
# Edit .env with your configurations
```

### Running the Application

1. Using Go directly:

```bash
make run
```

2. Using Docker:

```bash
docker-compose up
```

### Building

```bash
make build
```

## Features

- Clean Architecture implementation
- FastHTTP for high-performance HTTP handling
- PostgreSQL integration with transaction support
- Structured logging with Logrus
- Environment configuration management
- Encryption support for PII data
- Request/Response middleware
- Graceful shutdown

## Project Dependencies

- `fasthttp` - High-performance HTTP server
- `logrus` - Structured logging
- `pq` - PostgreSQL driver
- `encryption-pii` - PII data encryption
- `go-envconfig` - Environment configuration

## Development

1. **Adding New Features**

   - Add domain entities in `domain`
   - Implement use cases in `usecase`
   - Add repository implementations in `repository`
   - Create API routes in `api/route`

2. **Testing**

```bash
make test
```

## License

MIT License
