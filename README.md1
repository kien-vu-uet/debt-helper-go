# Debt Helper Project Template: Go Backend with MySQL/Postgres and Redis

## Table of contents

- [Introduction](#introduction)
- [Project Overview and Core Features](#project-overview-and-core-features)
- [Core Features](#core-features)
  - [User Management](#user-management)
  - [Group Management](#group-management)
  - [Transaction Logging](#transaction-logging)
  - [Debt Calculation](#debt-calculation)
- [Architectural Approach](#architectural-approach)
  - [Project structure](#project-structure)
  - [Database Schema Considerations (MySQL/Postgres)](#database-schema-considerations-mysqlpostgres)
  - [Redis Usage](#redis-usage)
  - [Key Technologies](#key-technologies)
- [Local Development Setup](#local-development-setup)
- [Conclusions & Recommendations](#conclusions--recommendations)

## Introduction
This document outlines a comprehensive project template for a "Debt Helper" application, designed to assist users in managing shared credits and debts within defined groups. The application will enable members to log transactions, automatically calculate group debts, and allow group administrators to mark debts as settled. This template leverages best practices for building scalable, maintainable, and robust Go backend services, integrating relational databases (MySQL/Postgres) for persistent storage and Redis for enhanced performance and real-time capabilities.

## Project Overview and Core Features
The Debt Helper application aims to simplify debt management among groups of people (e.g., friends, roommates, travel groups).

## Core Features:

### User Management:
- Register, login, and manage user profiles

### Group Management:
- Create new debt groups

- Invite and manage group members.

- Assign an administrator role to a group member.

- View group details and current debt status.

- Admin can mark a group's debts as settled.

### Transaction Logging:

- Members can log transactions, specifying who paid and who owes (one-to-one or one-to-many).

- Support for different currencies (optional, but good for future-proofing).

### Debt Calculation:

- Automatically calculate the net debt for each member within a group.

- Optimize debt settlement paths (e.g., minimize the number of transactions required to settle).

- Real-time Updates (Optional, via Redis Pub/Sub): Notify group members of new transactions or debt changes.

- Persistence: Store all group, member, transaction, and debt data reliably.

## Architectural Approach

The template will adhere to a Clean Architecture or Hexagonal Architecture pattern, emphasizing clear separation of concerns and dependency inversion. This ensures that the core business logic remains independent of external frameworks, databases, or third-party services, leading to a highly testable, maintainable, and adaptable system.

The internal/ directory will be crucial for enforcing these architectural boundaries, preventing unintended coupling between layers. Dependency Injection will be used to wire concrete implementations (e.g., database repositories, Redis clients) into the application's use cases and handlers at startup.

### Project structure
The following directory structure adapts the general Go backend template to the specific needs of the Debt Helper project:

```py
debt-helper/
├── cmd/                    # Application entry points for different executables
│   ├── api/                # Main entry point for the HTTP API server
│   │   └── main.go         # Initializes server, dependencies, and starts listeners
│   └── worker/             # (Optional) Entry point for background workers (e.g., for complex debt calculation, notifications)
│       └── main.go
│
├── internal/               # Private application and library code (not importable by other modules)
│   ├── config/             # Application configuration loading and parsing
│   │   └── config.go       # Defines configuration struct and loading logic
│   │
│   ├── domain/             # Core business entities, value objects, and domain interfaces
│   │   ├── user.go         # Defines User entity (ID, Name, Email, PasswordHash)
│   │   ├── group.go        # Defines Group entity (ID, Name, AdminUserID, Status)
│   │   ├── member.go       # Defines GroupMember entity (UserID, GroupID, Role)
│   │   ├── transaction.go  # Defines Transaction entity (ID, GroupID, PayerID, Amount, Description, Date)
│   │   ├── debt.go         # Defines Debt entity (ID, GroupID, OwerID, LenderID, Amount)
│   │   └── errors.go       # Custom domain-specific errors (e.g., ErrGroupNotFound, ErrUnauthorized)
│   │
│   ├── usecase/            # Business logic / Application layer
│   │   ├── user/           # User-related business rules
│   │   │   └── service.go  # RegisterUser, AuthenticateUser, GetUserByID
│   │   ├── group/          # Group management business rules
│   │   │   └── service.go  # CreateGroup, AddMemberToGroup, GetGroupDebts, SettleGroupDebts (admin only)
│   │   ├── transaction/    # Transaction logging business rules
│   │   │   └── service.go  # LogTransaction, GetTransactionsByGroup
│   │   └── debt_calculator/# Debt calculation logic
│   │       └── service.go  # CalculateGroupDebts (complex algorithm to minimize transactions)
│   │
│   ├── adapter/            # Interface Adapters (HTTP handlers, Repository implementations)
│   │   ├── http/           # HTTP handlers (Driving Adapter)
│   │   │   ├── user_handler.go     # Handles user registration, login
│   │   │   ├── group_handler.go    # Handles group creation, member management, debt viewing, settling
│   │   │   ├── transaction_handler.go # Handles logging new transactions
│   │   │   └── routes.go           # Defines HTTP routes and connects them to handlers
│   │   ├── repository/     # Database and external service repository implementations (Driven Adapter)
│   │   │   ├── mysql/      # MySQL-specific repository implementations
│   │   │   │   ├── user_repository.go      # Implements domain.UserRepository
│   │   │   │   ├── group_repository.go     # Implements domain.GroupRepository
│   │   │   │   ├── transaction_repository.go # Implements domain.TransactionRepository
│   │   │   │   └── debt_repository.go      # Implements domain.DebtRepository
│   │   │   ├── postgres/   # PostgreSQL-specific repository implementations (alternative to mysql)
│   │   │   │   ├── user_repository.go
│   │   │   │   └── ...
│   │   │   └── redis/      # Redis-backed implementations for caching or pub/sub
│   │   │       └── cache_repository.go     # Implements domain.CacheRepository (e.g., for debt calculations)
│   │   │       └── pubsub_adapter.go       # Implements domain.MessagePublisher (for real-time updates)
│   │   └── auth/           # Authentication and authorization logic
│   │       └── jwt_manager.go      # JWT token generation and validation
│   │       └── middleware.go       # Authentication and authorization middleware
│   │
│   ├── infrastructure/     # External concerns, concrete implementations (Frameworks & Drivers)
│   │   ├── database/       # Database connection setup, connection pooling
│   │   │   ├── mysql.go    # MySQL client initialization and pool configuration
│   │   │   └── postgres.go # PostgreSQL client (pgxpool) initialization and pool configuration
│   │   ├── redis/          # Redis client initialization and connection management
│   │   │   └── client.go   # Initializes go-redis client with pooling
│   │   ├── logger/         # Centralized logging setup (e.g., Logrus, Zap)
│   │   │   └── logger.go
│   │   └── validator/      # Data validation utilities (e.g., Go Playground Validator)
│   │       └── validator.go
│   │
│   └── shared/             # Common utilities, constants, and helpers
│       ├── constants.go
│       └── utils.go        # General utility functions
│       └── password.go     # Password hashing (e.g., bcrypt)
│
├── api/                    # API definitions (e.g., OpenAPI YAML)
│   └── openapi.yaml        # REST API specification for Debt Helper
│
├── db/                     # Database-related files
│   └── migrations/         # Database migration scripts (e.g., using golang-migrate or goose)
│       ├── 000001_create_users_table.up.sql
│       ├── 000002_create_groups_table.up.sql
│       ├── 000003_create_group_members_table.up.sql
│       ├── 000004_create_transactions_table.up.sql
│       ├── 000005_create_debts_table.up.sql
│       └── ...
│
├── scripts/                # Helper scripts for development, build, test, and deployment
│   ├── run.sh              # Script to run the application locally
│   ├── test.sh             # Script to run all tests
│   └── build.sh            # Script to build binaries
│
├── .env.example            # Example environment variables for local development
├── Dockerfile              # Docker build instructions for the Go application
├── docker-compose.yaml     # Docker Compose configuration for local development environment
├── go.mod                  # Go module definition file
├── go.sum                  # Go module checksums file
└── README.md               # Project documentation and setup instructions
```

### Database Schema Considerations (MySQL/Postgres)

Here are the essential tables and their relationships for the Debt Helper project:

* **`users` Table:**
    * `id` (PK, UUID/BIGINT, auto-generated)
    * `username` (VARCHAR, unique)
    * `email` (VARCHAR, unique)
    * `password_hash` (VARCHAR)
    * `created_at` (TIMESTAMP)
    * `updated_at` (TIMESTAMP)

* **`groups` Table:**
    * `id` (PK, UUID/BIGINT, auto-generated)
    * `name` (VARCHAR)
    * `admin_user_id` (FK to `users.id`)
    * `status` (VARCHAR, e.g., 'active', 'settled')
    * `created_at` (TIMESTAMP)
    * `updated_at` (TIMESTAMP)

* **`group_members` Table:** (Many-to-many relationship between users and groups)
    * `group_id` (PK, FK to `groups.id`)
    * `user_id` (PK, FK to `users.id`)
    * `role` (VARCHAR, e.g., 'member', 'admin')
    * `joined_at` (TIMESTAMP)

* **`transactions` Table:**
    * `id` (PK, UUID/BIGINT, auto-generated)
    * `group_id` (FK to `groups.id`)
    * `payer_user_id` (FK to `users.id`)
    * `amount` (DECIMAL/NUMERIC)
    * `description` (TEXT)
    * `transaction_date` (TIMESTAMP)
    * `created_at` (TIMESTAMP)
    * `updated_at` (TIMESTAMP)

* **`transaction_participants` Table:** (For one-to-many transactions)
    * `transaction_id` (PK, FK to `transactions.id`)
    * `participant_user_id` (PK, FK to `users.id`)
    * `share_amount` (DECIMAL/NUMERIC, optional, if not equal split)

* **`debts` Table:** (Represents the current net debts within a group)
    * `id` (PK, UUID/BIGINT, auto-generated)
    * `group_id` (FK to `groups.id`)
    * `ower_user_id` (FK to `users.id`)
    * `lender_user_id` (FK to `users.id`)
    * `amount` (DECIMAL/NUMERIC)
    * `status` (VARCHAR, e.g., 'outstanding', 'settled')
    * `last_calculated_at` (TIMESTAMP)
    * `created_at` (TIMESTAMP)
    * `updated_at` (TIMESTAMP)

**Migration Strategy:**
Use `golang-migrate` or `goose` to manage schema evolution. Each table creation, alteration, and index addition should be a separate, atomic migration script.

### Redis Usage

Redis can significantly enhance the Debt Helper application's performance and functionality:

* **Caching Debt Calculations:**
    * Complex debt calculation results for a group can be cached in Redis. When a new transaction is logged, the cache for that group is invalidated.
    * Use Redis Hashes to store debt summaries per group, or JSON strings for more complex structures.
    * Set appropriate TTLs for cached data to ensure freshness.
* **Session Management:**
    * Store user session tokens (e.g., JWT refresh tokens) in Redis for centralized session invalidation and management across multiple API instances.
    * Implement secure session ID generation, expiration, and secure cookie practices.
* **Real-time Notifications (Pub/Sub):**
    * When a new transaction is added or debts are settled, publish an event to a Redis Pub/Sub channel specific to the group.
    * Frontend clients (e.g., via WebSockets) can subscribe to these channels to receive real-time updates without constant polling.
* **Rate Limiting:** Implement API rate limiting using Redis counters to prevent abuse and ensure fair resource usage.

### Key Technologies

To build a robust, scalable, and maintainable Debt Helper backend in Go, the following technologies are recommended, chosen for their performance, community support, and alignment with Go's idioms and best practices:

* **Go Web Framework:**
    * **Gin (`github.com/gin-gonic/gin`)**: A high-performance HTTP web framework. Gin is known for its speed, middleware support, and a simple API, making it an excellent choice for building RESTful APIs. It's widely adopted and has a large, active community.
    * **Fiber (`github.com/gofiber/fiber/v2`)**: An Express.js-inspired web framework built on top of Fasthttp, known for its extreme speed and low memory footprint. Fiber is a strong alternative to Gin if raw performance is a top priority, offering a similar middleware-based approach.

* **Database Driver:**
    * **MySQL: `github.com/go-sql-driver/mysql`**: The most popular and well-maintained MySQL driver for Go. It's robust, performant, and adheres to the `database/sql` interface.
    * **PostgreSQL: `github.com/jackc/pgx/v5` (with `pgxpool` for connection pooling)**: `pgx` is the modern, high-performance PostgreSQL driver for Go. It offers superior performance, native PostgreSQL feature support (e.g., notifications, COPY protocol, rich type support), and `pgxpool` provides efficient connection pooling, which is crucial for high-concurrency applications. It is generally preferred over `github.com/lib/pq` for new projects due to its features and performance.

* **ORM/Query Builder (Choose one or a combination based on project needs):**
    * **GORM (`gorm.io/gorm`)**: A developer-friendly ORM that maps Go structs to database tables, simplifying CRUD operations. It offers a rich feature set including associations, hooks, and automatic migrations (though `AutoMigrate` is generally not recommended for production). GORM accelerates development for common database interactions.
    * **`sqlc` (`github.com/sqlc-dev/sqlc`)**: Generates type-safe Go code from raw SQL queries. This approach provides the benefits of type safety and compile-time checking while giving developers full control over SQL, making it ideal for performance-critical queries or teams with strong SQL expertise.
    * **SQLX (`github.com/jmoiron/sqlx`)**: An extension to `database/sql` that provides utilities for working with SQL databases, including mapping query results to Go structs. It offers a good balance between the control of raw SQL and the convenience of an ORM, without the full abstraction overhead.

* **Redis Client:**
    * **`github.com/redis/go-redis/v9`**: The de facto standard and highly recommended Redis client for Go. It provides a comprehensive API, built-in connection pooling, support for transactions and pipelines, and integrates well with observability tools.

* **Configuration:**
    * **`github.com/spf13/viper`**: A complete Go configuration solution. Viper can read configurations from JSON, TOML, YAML, HCL, INI, environment variables, and remote configuration systems. It's highly flexible and widely used for managing application settings.

* **Validation:**
    * **`github.com/go-playground/validator/v10`**: A powerful and extensible Go package for struct and field validation. It allows defining validation rules using struct tags, ensuring that incoming data conforms to business requirements before processing.

* **Logging:**
    * **`github.com/sirupsen/logrus`**: A popular and flexible logging library for Go, providing structured logging capabilities.
    * **`go.uber.org/zap`**: A highly performant and structured logging library, ideal for high-throughput services where logging overhead needs to be minimized. It's a strong choice for production environments.

* **Password Hashing:**
    * **`golang.org/x/crypto/bcrypt`**: A robust and secure library for hashing passwords. It's computationally intensive, making brute-force attacks difficult, and includes a cost factor that can be adjusted to keep pace with increasing computing power.

* **JWT (JSON Web Tokens):**
    * **`github.com/golang-jwt/jwt/v5`**: A well-maintained library for working with JWTs, essential for stateless authentication and authorization in RESTful APIs. It supports various signing methods and token validation.

* **Database Migrations:**
    * **`github.com/golang-migrate/migrate/v4`**: A versatile database migration tool that supports various databases and allows managing migrations using SQL or Go. It's widely adopted for its reliability and CLI support.
    * **`github.com/pressly/goose`**: Another popular migration tool that supports both SQL and Go-based migrations. It's often favored for its simplicity and good integration with existing Go projects.

### Local Development Setup

A [`docker-compose.yaml`](docker-compose.yaml) file will be provided to orchestrate the local development environment, including:

* **`db` service:** A MySQL or PostgreSQL container.
* **`redis` service:** A Redis container.
* **`app` service:** The Go backend application container, configured to connect to `db` and `redis` services.

This setup ensures a consistent and easily reproducible environment for all developers.

### Conclusions & Recommendations

This template provides a robust foundation for developing the Debt Helper application. By adhering to a layered architectural approach, carefully selecting database interaction strategies, and intelligently leveraging Redis, the project can achieve high levels of scalability, maintainability, and reliability.
