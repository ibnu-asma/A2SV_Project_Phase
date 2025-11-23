# Clean Architecture Documentation

## Overview

This Task Management API has been refactored following Clean Architecture principles to achieve better separation of concerns, maintainability, and testability.

## Architecture Layers

### 1. Domain Layer (`Domain/`)
**Purpose:** Core business entities and rules

**Responsibilities:**
- Define core business entities (Task, User)
- Contains no external dependencies
- Independent of frameworks, databases, or UI

**Files:**
- `domain.go`: Core entities (Task, User) and DTOs (LoginRequest, RegisterRequest)

**Key Principle:** Domain layer is the innermost layer and has no dependencies on other layers.

---

### 2. Repository Layer (`Repositories/`)
**Purpose:** Data access abstraction

**Responsibilities:**
- Define interfaces for data operations
- Implement database operations
- Abstract MongoDB specifics from business logic

**Files:**
- `task_repository.go`: Task data access interface and implementation
- `user_repository.go`: User data access interface and implementation

**Key Principle:** Repositories implement interfaces, allowing easy substitution for testing or changing databases.

---

### 3. Use Cases Layer (`Usecases/`)
**Purpose:** Application business logic

**Responsibilities:**
- Orchestrate data flow between layers
- Implement business rules
- Coordinate repository and infrastructure services

**Files:**
- `task_usecases.go`: Task-related business logic
- `user_usecases.go`: User authentication and authorization logic

**Key Principle:** Use cases depend on repository interfaces, not implementations (Dependency Inversion).

---

### 4. Infrastructure Layer (`Infrastructure/`)
**Purpose:** External services and frameworks

**Responsibilities:**
- JWT token generation and validation
- Password hashing and comparison
- Authentication middleware

**Files:**
- `jwt_service.go`: JWT token operations
- `password_service.go`: Password hashing with bcrypt
- `auth_middleWare.go`: Authentication and authorization middleware

**Key Principle:** Infrastructure provides services to use cases without coupling business logic to external libraries.

---

### 5. Delivery Layer (`Delivery/`)
**Purpose:** HTTP request handling and routing

**Responsibilities:**
- Handle HTTP requests and responses
- Route configuration
- Dependency injection and initialization

**Files:**
- `main.go`: Application entry point with dependency injection
- `controllers/controller.go`: HTTP request handlers
- `routers/router.go`: Route configuration

**Key Principle:** Delivery layer depends on use cases through interfaces, making it easy to add other delivery mechanisms (gRPC, CLI, etc.).

---

## Dependency Flow

```
Delivery → Usecases → Repositories → Domain
    ↓
Infrastructure
```

**Rules:**
1. Inner layers know nothing about outer layers
2. Dependencies point inward
3. Interfaces are defined in inner layers, implemented in outer layers

---

## Key Design Decisions

### 1. Interface-Based Design
All repositories and use cases are defined as interfaces, allowing:
- Easy mocking for unit tests
- Swapping implementations without changing business logic
- Loose coupling between layers

### 2. Dependency Injection
All dependencies are injected in `main.go`:
```go
taskRepo := repositories.NewTaskRepository(client, "taskdb", "tasks")
taskUsecase := usecases.NewTaskUsecase(taskRepo)
taskController := controllers.NewTaskController(taskUsecase)
```

### 3. Separation of Concerns
- **Domain**: Pure business entities
- **Repositories**: Data access only
- **Use Cases**: Business logic only
- **Infrastructure**: External services only
- **Delivery**: HTTP handling only

### 4. Single Responsibility Principle
Each layer and component has one clear responsibility:
- `TaskRepository`: Task data operations
- `TaskUsecase`: Task business logic
- `TaskController`: HTTP request handling

---

## Benefits of This Architecture

### 1. Testability
- Mock repositories for use case tests
- Mock use cases for controller tests
- Test business logic without database

### 2. Maintainability
- Clear separation makes code easy to understand
- Changes in one layer don't affect others
- Easy to locate and fix bugs

### 3. Scalability
- Add new features without modifying existing code
- Easy to add new delivery mechanisms (gRPC, GraphQL)
- Simple to swap databases or external services

### 4. Independence
- Business logic independent of frameworks
- Can change Gin to another framework easily
- Can switch from MongoDB to PostgreSQL with minimal changes

---

## Layer Dependencies

### Domain Layer
- **Depends on:** Nothing
- **Used by:** All other layers

### Repository Layer
- **Depends on:** Domain
- **Used by:** Use Cases

### Use Cases Layer
- **Depends on:** Domain, Repository interfaces, Infrastructure interfaces
- **Used by:** Delivery

### Infrastructure Layer
- **Depends on:** Domain (for types)
- **Used by:** Use Cases, Delivery

### Delivery Layer
- **Depends on:** All layers (for initialization)
- **Used by:** External clients (HTTP requests)

---

## How to Add New Features

### Adding a New Entity (e.g., Project)

1. **Domain Layer**: Define `Project` struct in `domain.go`
2. **Repository Layer**: Create `project_repository.go` with interface and implementation
3. **Use Cases Layer**: Create `project_usecases.go` with business logic
4. **Delivery Layer**: Add controller methods and routes
5. **Main**: Wire up dependencies

### Adding a New Use Case

1. Add method to use case interface
2. Implement method in use case struct
3. Add controller method to call use case
4. Add route in router

---

## Testing Strategy

### Unit Tests

**Domain Layer:**
```go
// Test pure business logic
func TestTaskValidation(t *testing.T) { ... }
```

**Use Cases Layer:**
```go
// Mock repository
mockRepo := &MockTaskRepository{}
usecase := NewTaskUsecase(mockRepo)
// Test business logic
```

**Repository Layer:**
```go
// Use test database or mock MongoDB
```

**Controllers:**
```go
// Mock use cases
mockUsecase := &MockTaskUsecase{}
controller := NewTaskController(mockUsecase)
// Test HTTP handling
```

---

## Migration from Old Structure

### Old Structure
```
task_manager/
├── main.go
├── controllers/
├── models/
├── data/
├── middleware/
└── router/
```

### New Structure (Clean Architecture)
```
task_manager/
├── Delivery/
│   ├── main.go
│   ├── controllers/
│   └── routers/
├── Domain/
├── Infrastructure/
├── Repositories/
└── Usecases/
```

### Key Changes
1. **Models → Domain**: Business entities moved to Domain layer
2. **Data → Repositories**: Data access logic with interfaces
3. **Middleware → Infrastructure**: External services separated
4. **Controllers**: Now depend on use cases, not services directly
5. **Main**: Moved to Delivery with dependency injection

---

## Best Practices

### 1. Keep Domain Pure
- No external dependencies in Domain layer
- No database tags in domain entities (use separate models if needed)

### 2. Use Interfaces
- Define interfaces in the layer that uses them
- Implement interfaces in outer layers

### 3. Dependency Injection
- Inject all dependencies through constructors
- Avoid global variables

### 4. Error Handling
- Return errors from inner layers
- Handle and format errors in Delivery layer

### 5. Keep Use Cases Thin
- Orchestrate, don't implement
- Delegate to repositories and services

---

## Running the Application

```bash
cd Delivery
go run main.go
```

The application maintains the same API endpoints and functionality while following Clean Architecture principles.

---

## Future Enhancements

1. **Add Unit Tests**: Test each layer independently
2. **Add Integration Tests**: Test layer interactions
3. **Add Logging**: Structured logging in each layer
4. **Add Metrics**: Monitor use case execution
5. **Add Caching**: Cache layer between use cases and repositories
6. **Add Events**: Event-driven architecture for async operations

---

## Conclusion

This Clean Architecture implementation provides:
- ✅ Clear separation of concerns
- ✅ Testable business logic
- ✅ Independent of frameworks
- ✅ Flexible and maintainable
- ✅ Scalable architecture

The architecture follows SOLID principles and makes the codebase easier to understand, test, and extend.
