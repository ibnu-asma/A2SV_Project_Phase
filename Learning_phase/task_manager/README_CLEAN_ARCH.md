# Task Management API - Clean Architecture

A Task Management REST API built with Go, following Clean Architecture principles for maintainability, testability, and scalability.

## Features

- ✅ Clean Architecture implementation
- ✅ JWT authentication and authorization
- ✅ Role-based access control
- ✅ MongoDB persistent storage
- ✅ Interface-based design
- ✅ Dependency injection
- ✅ Separation of concerns

## Architecture

### Layers

```
Delivery/          → HTTP handlers, routing, main entry point
├── main.go
├── controllers/
└── routers/

Domain/            → Core business entities (Task, User)
└── domain.go

Infrastructure/    → External services (JWT, Password, Auth)
├── jwt_service.go
├── password_service.go
└── auth_middleWare.go

Repositories/      → Data access abstraction
├── task_repository.go
└── user_repository.go

Usecases/          → Business logic orchestration
├── task_usecases.go
└── user_usecases.go
```

### Dependency Flow
```
Delivery → Usecases → Repositories → Domain
    ↓
Infrastructure
```

## Installation

```bash
cd task_manager
go mod tidy
```

## Running the Application

```bash
cd Delivery
go run main.go
```

Server starts on: `http://localhost:8080`

## API Endpoints

### Public
- `POST /register` - Register new user
- `POST /login` - Login and get JWT token

### Protected (All Users)
- `GET /tasks` - Get all tasks
- `GET /tasks/:id` - Get task by ID

### Admin Only
- `POST /tasks` - Create task
- `PUT /tasks/:id` - Update task
- `DELETE /tasks/:id` - Delete task
- `PUT /promote/:username` - Promote user to admin

## Quick Start

### 1. Start MongoDB
```bash
mongod
```

### 2. Register Admin User
```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"admin123"}'
```

### 3. Login
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"admin123"}'
```

### 4. Create Task (with token)
```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -d '{
    "title":"Clean Architecture Task",
    "description":"Refactor to Clean Architecture",
    "due_date":"2024-12-31T23:59:59Z",
    "status":"In Progress"
  }'
```

## Clean Architecture Benefits

### 1. Testability
- Mock repositories for use case tests
- Mock use cases for controller tests
- Test business logic without database

### 2. Maintainability
- Clear separation of concerns
- Easy to locate and modify code
- Changes in one layer don't affect others

### 3. Scalability
- Add new features without breaking existing code
- Easy to add new delivery mechanisms (gRPC, CLI)
- Simple to swap databases

### 4. Independence
- Business logic independent of frameworks
- Can change web framework easily
- Can switch databases with minimal changes

## Layer Responsibilities

### Domain
- Core business entities
- No external dependencies
- Pure Go structs

### Repositories
- Data access interfaces
- MongoDB operations
- Abstract database details

### Use Cases
- Business logic orchestration
- Coordinate repositories and services
- Enforce business rules

### Infrastructure
- JWT token management
- Password hashing
- Authentication middleware

### Delivery
- HTTP request handling
- Routing configuration
- Dependency injection

## Design Principles

### SOLID Principles
- **S**ingle Responsibility: Each layer has one purpose
- **O**pen/Closed: Open for extension, closed for modification
- **L**iskov Substitution: Interfaces can be substituted
- **I**nterface Segregation: Small, focused interfaces
- **D**ependency Inversion: Depend on abstractions, not concretions

### Clean Architecture Rules
1. Dependencies point inward
2. Inner layers know nothing about outer layers
3. Interfaces defined in inner layers
4. Implementations in outer layers

## Testing

### Unit Tests Example

```go
// Test use case with mock repository
func TestCreateTask(t *testing.T) {
    mockRepo := &MockTaskRepository{}
    usecase := NewTaskUsecase(mockRepo)
    
    task := domain.Task{Title: "Test"}
    result, err := usecase.CreateTask(task)
    
    assert.NoError(t, err)
    assert.Equal(t, "Test", result.Title)
}
```

## Adding New Features

### Example: Add Project Entity

1. **Domain**: Add `Project` struct
2. **Repository**: Create `project_repository.go`
3. **Use Case**: Create `project_usecases.go`
4. **Controller**: Add project handlers
5. **Router**: Add project routes
6. **Main**: Wire dependencies

## Migration from Old Structure

The application has been refactored from:
```
controllers/ → Delivery/controllers/
models/      → Domain/
data/        → Repositories/ + Usecases/
middleware/  → Infrastructure/
router/      → Delivery/routers/
main.go      → Delivery/main.go
```

## Documentation

- [Clean Architecture Guide](CLEAN_ARCHITECTURE.md) - Detailed architecture documentation
- [API Documentation](docs/api_documentation.md) - API endpoints and usage
- [Testing Guide](TESTING_GUIDE.md) - Testing instructions

## Technologies

- **Go** - Programming language
- **Gin** - Web framework
- **MongoDB** - Database
- **JWT** - Authentication
- **bcrypt** - Password hashing

## Project Structure Comparison

### Before (Layered Architecture)
```
task_manager/
├── main.go
├── controllers/
├── models/
├── data/
├── middleware/
└── router/
```

### After (Clean Architecture)
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

## Key Improvements

1. **Separation of Concerns**: Each layer has clear responsibility
2. **Testability**: Easy to mock and test each layer
3. **Flexibility**: Easy to swap implementations
4. **Maintainability**: Clear structure, easy to navigate
5. **Scalability**: Add features without breaking existing code

## Environment

- MongoDB URI: `mongodb://localhost:27017`
- Database: `taskdb`
- Collections: `tasks`, `users`
- Port: `8080`

## License

MIT
