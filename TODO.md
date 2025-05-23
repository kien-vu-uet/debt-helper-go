## Project Phases & Timeline Overview

| Phase                                   | Estimated Duration | Milestone                                                        |
|------------------------------------------|-------------------|------------------------------------------------------------------|
| [Phase 1: Foundation & User Management](#phase-1-foundation--user-management-estimated-3-4-weeks)    | 3-4 Weeks         | Basic User Authentication & Profile Management API Ready         |
| [Phase 2: Core Group & Transaction Logic](#phase-2-core-group--transaction-logic-estimated-4-5-weeks)  | 4-5 Weeks         | Group Creation, Member Management, and Transaction Logging API Ready |
| [Phase 3: Debt Calculation & Optimization](#phase-3-debt-calculation--optimization-estimated-3-4-weeks) | 3-4 Weeks         | Automated Debt Calculation & Settlement Path API Ready           |
| [Phase 4: Real-time & Advanced Features](#phase-4-real-time--advanced-features-estimated-3-4-weeks)   | 3-4 Weeks         | Real-time Updates & Enhanced Group Features Implemented          |
| [Phase 5: Polish, Testing & Deployment](#phase-5-polish-testing--deployment-estimated-2-3-weeks)    | 2-3 Weeks         | Production-Ready Backend Deployed & Monitored                    |

# Debt Helper Project: Detailed TODO List, Timeline, and Milestones
This document outlines a comprehensive TODO list for the "Debt Helper" backend project, breaking down the development into manageable phases with estimated timelines and clear milestones. This plan is designed to be agile and can be adapted based on team size, resources, and evolving requirements.

## Project Phases & Timeline Overview

| Phase                                   | Estimated Duration | Milestone                                                        |
|------------------------------------------|-------------------|------------------------------------------------------------------|
| [Phase 1: Foundation & User Management](#phase-1-foundation--user-management-estimated-3-4-weeks)    | 3-4 Weeks         | Basic User Authentication & Profile Management API Ready         |
| [Phase 2: Core Group & Transaction Logic](#phase-2-core-group--transaction-logic-estimated-4-5-weeks)  | 4-5 Weeks         | Group Creation, Member Management, and Transaction Logging API Ready |
| [Phase 3: Debt Calculation & Optimization](#phase-3-debt-calculation--optimization-estimated-3-4-weeks) | 3-4 Weeks         | Automated Debt Calculation & Settlement Path API Ready           |
| [Phase 4: Real-time & Advanced Features](#phase-4-real-time--advanced-features-estimated-3-4-weeks)   | 3-4 Weeks         | Real-time Updates & Enhanced Group Features Implemented          |
| [Phase 5: Polish, Testing & Deployment](#phase-5-polish-testing--deployment-estimated-2-3-weeks)    | 2-3 Weeks         | Production-Ready Backend Deployed & Monitored                    |

## Detailed TODO List

### Phase 1: Foundation & User Management (Estimated: 3-4 Weeks)
**Goal**: Establish the basic project structure, configure core services, and implement fundamental user authentication and profile management.

**Features/Tasks**:

- [x] Project Setup:
    - [x] Initialize Go module (go mod init debt-helper).
    - [x] Set up recommended directory structure.
    - [x] Create Dockerfile and docker-compose.yaml for local database (MySQL/Postgres) and Redis.
    - [x] Configure internal/config with viper for environment variable loading.
    - [x] Set up internal/infrastructure/logger (e.g., Zap/Logrus).
- [ ] Database Setup & Migrations:
    - [x] Implement database connection pooling (internal/infrastructure/database/mysql.go or postgres.go).
    - [x] Create initial migration scripts (db/migrations/) for users table.
    - [x] Integrate golang-migrate or goose for migration management.
- [ ] User Domain & Repository:
    - [x] Define User entity in internal/domain/user.go.
    - [x] Define UserRepository interface in internal/domain/user.go.
    - [x] Implement MySQLUserRepository (or PostgresUserRepository) in internal/adapter/repository/mysql/user_repository.go.
- [ ] User Use Cases:
    - [ ] Implement RegisterUser use case in internal/usecase/user/service.go (including password hashing with bcrypt).
    - [ ] Implement AuthenticateUser use case in internal/usecase/user/service.go.
    - [ ] Implement GetUserByID use case.
- [ ] Authentication & Authorization:
    - [ ] Implement JWT token generation and validation (internal/adapter/auth/jwt_manager.go).
    - [ ] Create authentication middleware (internal/adapter/auth/middleware.go) for protected routes.
- [ ] HTTP API & Handlers:
    - [ ] Set up Gin/Fiber framework (cmd/api/main.go).
    - [ ] Define basic HTTP routes (internal/adapter/http/routes.go).
    - [ ] Implement UserHandler (internal/adapter/http/user_handler.go) for registration, login, and profile retrieval.
    - [ ] Integrate go-playground/validator for input validation.
- [ ] Initial Testing:
    - [ ] Write unit tests for user domain, usecase, and repository layers.
    - [ ] Write integration tests for user registration and login API endpoints.

**Milestone 1**: Basic User Authentication & Profile Management API Ready. Users can register, log in, and retrieve their own profile information. (Achieved on 2025-05-23)

### Phase 2: Core Group & Transaction Logic (Estimated: 4-5 Weeks)
**Goal**: Enable users to create and manage groups, invite members, and log transactions within these groups.

**Features/Tasks**:

- [ ] Database Migrations:
    - [ ] Create migration scripts for groups, group_members, transactions, and transaction_participants tables.
- [ ] Group Domain & Repository:
    - [ ] Define Group, GroupMember, Transaction, TransactionParticipant entities in internal/domain/.
    - [ ] Define corresponding repository interfaces (GroupRepository, TransactionRepository).
    - [ ] Implement MySQLGroupRepository, MySQLTransactionRepository (or Postgres equivalents).
- [ ] Group Use Cases:
    - [ ] Implement CreateGroup use case (assigning admin role).
    - [ ] Implement AddMemberToGroup use case (invitation/joining logic).
    - [ ] Implement GetGroupByID and ListUserGroups use cases.
- [ ] Transaction Use Cases:
    - [ ] Implement LogTransaction use case (supporting one-to-one and one-to-many splits).
    - [ ] Implement GetTransactionsByGroup use case.
- [ ] HTTP API & Handlers:
    - [ ] Extend GroupHandler for group creation, member management.
    - [ ] Implement TransactionHandler for logging transactions.
    - [ ] Add authorization checks for group actions (e.g., only group members can log transactions in a group).
- [ ] Redis Integration (Basic):
    - [ ] Initialize Redis client (internal/infrastructure/redis/client.go).
    - [ ] Implement basic caching for frequently accessed group data (e.g., group details) in internal/adapter/repository/redis/cache_repository.go.
- [ ] Testing:
    - [ ] Expand unit and integration tests for new domain, use case, and adapter components.

**Milestone 2**: Group Creation, Member Management, and Transaction Logging API Ready. Users can create groups, add members, and log various types of transactions.

### Phase 3: Debt Calculation & Optimization (Estimated: 3-4 Weeks)
**Goal**: Implement the core debt calculation logic and provide endpoints to view current debts and optimize settlement.

**Features/Tasks**:

- [ ] Database Migrations:
    - [ ] Ensure debts table migration is complete.
- [ ] Debt Domain & Repository:
    - [ ] Define Debt entity in internal/domain/debt.go.
    - [ ] Define DebtRepository interface.
    - [ ] Implement MySQLDebtRepository (or Postgres equivalent).
- [ ] Debt Calculation Logic:
    - [ ] Implement CalculateGroupDebts service (internal/usecase/debt_calculator/service.go). This is the complex algorithm to determine net debts and potentially optimize settlement paths (e.g., using a graph algorithm like minimum cost flow or a simplified approach to reduce transactions).
    - [ ] Consider making this an idempotent operation that can be triggered manually or via a background job.
- [ ] Debt Use Cases:
    - [ ] Implement GetGroupDebts use case.
    - [ ] Implement SettleGroupDebts use case (admin-only, updates group.status and debt.status).
- [ ] HTTP API & Handlers:
    - [ ] Extend GroupHandler to expose endpoints for GetGroupDebts and SettleGroupDebts.
- [ ] Redis Integration (Advanced):
    - [ ] Utilize Redis for caching the results of CalculateGroupDebts to improve performance. Invalidate cache on new transactions or group member changes.
    - [ ] (Optional) Implement a simple background worker (cmd/worker/main.go) that periodically recalculates debts for active groups or is triggered by a message queue.
- [ ] Testing:
    - [ ] Thorough unit tests for the debt_calculator service (critical component).
    - [ ] Integration tests for debt calculation and settlement flows.

**Milestone 3**: Automated Debt Calculation & Settlement Path API Ready. The system can accurately calculate debts within a group and allow admins to settle them.

### Phase 4: Real-time & Advanced Features (Estimated: 3-4 Weeks)
**Goal**: Add real-time capabilities and enhance the user experience with additional features.

**Features/Tasks**:

- [ ] Real-time Notifications:
    - [ ] Implement Redis Pub/Sub for notifications (internal/adapter/repository/redis/pubsub_adapter.go).
    - [ ] Integrate Pub/Sub into Transaction and Group use cases (e.g., publish "transaction_added" or "group_settled" events).
    - [ ] (Requires frontend integration, but backend should provide the mechanism, e.g., WebSockets endpoint).
- [ ] Currency Support (if not done in Phase 2):
    - [ ] Implement currency handling (e.g., storing currency code with amounts, potential for conversion if needed).
- [ ] User Interface Enhancements (Backend Support):
    - [ ] Endpoints for listing group members with their individual balances (sum of transactions, not just net debt).
    - [ ] Endpoint for viewing transaction history for a specific member within a group.
- [ ] Rate Limiting:
    - [ ] Implement API rate limiting using Redis counters for critical endpoints (e.g., login, transaction logging).
- [ ] Error Handling & Observability:
    - [ ] Refine custom error handling (internal/domain/errors.go).
    - [ ] Add tracing/metrics instrumentation (e.g., OpenTelemetry if desired).
- [ ] Security Enhancements:
    - [ ] Implement CSRF protection for relevant endpoints.
    - [ ] Review and harden JWT security (e.g., refresh tokens, blacklisting).

**Milestone 4**: Real-time Updates & Enhanced Group Features Implemented. Users receive real-time notifications, and the API supports more detailed financial insights.

### Phase 5: Polish, Testing & Deployment (Estimated: 2-3 Weeks)
**Goal**: Prepare the backend for production, ensuring robustness, performance, and easy deployment.

**Features/Tasks**:

- [ ] Comprehensive Testing:
    - [ ] Execute full end-to-end tests covering all critical user flows.
    - [ ] Perform performance testing (load testing) to identify bottlenecks.
    - [ ] Conduct security audits (e.g., penetration testing, vulnerability scanning).
- [ ] Code Review & Refactoring:
    - [ ] Thorough code review across all layers.
    - [ ] Refactor any identified technical debt or areas for improvement.
    - [ ] Ensure consistent code style (e.g., gofmt, golangci-lint).
- [ ] Documentation:
    - [ ] Update README.md with detailed setup, build, and run instructions.
    - [ ] Generate OpenAPI specification (api/openapi.yaml) and ensure it's up-to-date.
    - [ ] Add inline code comments where necessary.
- [ ] Deployment Preparation:
    - [ ] Optimize Dockerfile for smaller image size and faster builds.
    - [ ] Finalize docker-compose.yaml for production-like local testing.
    - [ ] Prepare deployment scripts (e.g., for Kubernetes, cloud VMs, etc. - beyond the scope of this template's direct code, but a necessary step).
- [ ] Monitoring & Alerting:
    - [ ] Integrate with a monitoring system (e.g., Prometheus, Grafana) for application metrics.
    - [ ] Set up logging aggregation (e.g., ELK stack, cloud logging services).
    - [ ] Configure alerts for critical errors or performance degradation.

**Milestone 5**: Production-Ready Backend Deployed & Monitored. The Debt Helper backend is stable, performant, secure, and ready for user traffic.
