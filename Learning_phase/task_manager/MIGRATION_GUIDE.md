# Migration Guide: From Layered to Clean Architecture

## Overview

This guide explains how the Task Management API was refactored from a traditional layered architecture to Clean Architecture.

## Before and After

### Old Structure (Layered Architecture)
```
task_manager/
├── main.go                    # Entry point
├── controllers/
│   └── controller.go          # HTTP handlers
├── models/
│   ├── task.go               # Data models
│   └── user.go
├── data/
│   ├── task_service.go       # Business logic + data access
│   └── user_service.go
├── middleware/
│   └── auth_middleware.go    # Authentication
└── router/
    └── router.go             # Routes
```

### New Structure (Clean Architecture)
```
task_manager/
├── Delivery/
│   ├── main.go               # Entry point with DI
│   ├── controllers/
│   │   └── controller.go     # HTTP handlers only
│   └── routers/
│       └── router.go         # Routes
├── Domain/
│   └── domain.go             # Core entities
├── Infrastructure/
│   ├── auth_middleWare.go    # Auth middleware
│   ├── jwt_service.go        # JWT operations
│   └── password_service.go   # Password hashing
├── Repositories/
│   ├── task_repository.go    # Data access interface
│   └── user_repository.go
└── Usecases/
    ├── task_usecases.go      # Business logic
    └── user_usecases.go
```

## Key Changes

### 1. Models → Domain

**Before:**
```go
// models/task.go
package models

type Task struct {
    ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
    Title       string             `json:"title" binding:"required"`
    // ...
}
```

**After:**
```go
// Domain/domain.go
package domain

type Task struct {
    ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
    Title       string             `json:"title" bson:"title"`
    // No binding tags - pure domain entity
}
```

**Changes:**
- Moved to `Domain/` package
- Removed framework-specific tags (binding)
- Pure business entities

---

### 2. Data Services → Repositories + Use Cases

**Before:**
```go
// data/task_service.go
type TaskService struct {
    collection *mongo.Collection
}

func (s *TaskService) GetAllTasks() ([]models.Task, error) {
    // MongoDB operations + business logic mixed
}
```

**After:**
```go
// Repositories/task_repository.go
type TaskRepository interface {
    GetAll() ([]domain.Task, error)
    // Interface definition
}

type taskRepository struct {
    collection *mongo.Collection
}

func (r *taskRepository) GetAll() ([]domain.Task, error) {
    // Only MongoDB operations
}

// Usecases/task_usecases.go
type TaskUsecase interface {
    GetAllTasks() ([]domain.Task, error)
}

type taskUsecase struct {
    taskRepo repositories.TaskRepository
}

func (u *taskUsecase) GetAllTasks() ([]domain.Task, error) {
    return u.taskRepo.GetAll()
}
```

**Changes:**
- Split into Repository (data access) and Use Case (business logic)
- Added interfaces for testability
- Clear separation of concerns

---

### 3. Middleware → Infrastructure

**Before:**
```go
// middleware/auth_middleware.go
package middleware

func AuthMiddleware() gin.HandlerFunc {
    // JWT validation logic
}

func GenerateToken(...) (string, error) {
    // Token generation
}
```

**After:**
```go
// Infrastructure/jwt_service.go
type JWTService struct{}

func (js *JWTService) GenerateToken(...) (string, error) {
    // Token generation
}

func (js *JWTService) ValidateToken(...) (*Claims, error) {
    // Token validation
}

// Infrastructure/auth_middleWare.go
type AuthMiddleware struct {
    jwtService *JWTService
}

func (am *AuthMiddleware) AuthRequired() gin.HandlerFunc {
    // Uses jwtService
}
```

**Changes:**
- Separated JWT logic into service
- Middleware uses JWT service
- Better testability and reusability

---

### 4. Controllers → Delivery/Controllers

**Before:**
```go
// controllers/controller.go
type TaskController struct {
    service *data.TaskService  // Direct dependency
}

func (tc *TaskController) GetTasks(c *gin.Context) {
    tasks := tc.service.GetAllTasks()
    // ...
}
```

**After:**
```go
// Delivery/controllers/controller.go
type TaskController struct {
    taskUsecase usecases.TaskUsecase  // Interface dependency
}

func (tc *TaskController) GetTasks(c *gin.Context) {
    tasks, err := tc.taskUsecase.GetAllTasks()
    // ...
}
```

**Changes:**
- Depends on use case interface, not implementation
- Moved to Delivery layer
- Easier to test with mocks

---

### 5. Main → Delivery/Main with Dependency Injection

**Before:**
```go
// main.go
func main() {
    client, _ := mongo.Connect(...)
    
    service := data.NewTaskService(client, "taskdb", "tasks")
    controller := controllers.NewTaskController(service)
    
    r := router.SetupRouter(client)
    r.Run(":8080")
}
```

**After:**
```go
// Delivery/main.go
func main() {
    client, _ := mongo.Connect(...)
    
    // Repository layer
    taskRepo := repositories.NewTaskRepository(client, "taskdb", "tasks")
    userRepo := repositories.NewUserRepository(client, "taskdb", "users")
    
    // Infrastructure layer
    passwordService := infrastructure.NewPasswordService()
    jwtService := infrastructure.NewJWTService()
    
    // Use case layer
    taskUsecase := usecases.NewTaskUsecase(taskRepo)
    userUsecase := usecases.NewUserUsecase(userRepo, passwordService, jwtService)
    
    // Delivery layer
    taskController := controllers.NewTaskController(taskUsecase)
    userController := controllers.NewUserController(userUsecase)
    authMiddleware := infrastructure.NewAuthMiddleware(jwtService)
    
    r := routers.SetupRouter(taskController, userController, authMiddleware)
    r.Run(":8080")
}
```

**Changes:**
- Explicit dependency injection
- Clear initialization order
- All dependencies visible

---

## Benefits of Migration

### 1. Testability

**Before:**
```go
// Hard to test - requires MongoDB
func TestGetTasks(t *testing.T) {
    service := data.NewTaskService(mongoClient, "testdb", "tasks")
    // Need real database
}
```

**After:**
```go
// Easy to test - use mock
type MockTaskRepository struct{}

func (m *MockTaskRepository) GetAll() ([]domain.Task, error) {
    return []domain.Task{{Title: "Test"}}, nil
}

func TestGetTasks(t *testing.T) {
    mockRepo := &MockTaskRepository{}
    usecase := usecases.NewTaskUsecase(mockRepo)
    // No database needed
}
```

### 2. Flexibility

**Before:**
- Changing database requires modifying service layer
- Business logic mixed with data access

**After:**
- Swap repository implementation without touching use cases
- Business logic independent of database

### 3. Maintainability

**Before:**
- Hard to find where business logic lives
- Changes ripple across layers

**After:**
- Clear location for each concern
- Changes isolated to specific layers

---

## Running Both Versions

### Old Version
```bash
go run main.go
```

### New Version (Clean Architecture)
```bash
cd Delivery
go run main.go
```

Both versions provide the same API functionality!

---

## API Compatibility

✅ All endpoints remain the same
✅ Request/response formats unchanged
✅ Authentication works identically
✅ Database schema unchanged

The refactoring is **100% backward compatible** from the API perspective.

---

## What Stayed the Same

1. **API Endpoints**: All routes unchanged
2. **Authentication**: JWT still works the same way
3. **Database**: MongoDB operations identical
4. **Functionality**: All features preserved

## What Changed

1. **Code Organization**: Better separation of concerns
2. **Dependencies**: Interface-based, easier to test
3. **Structure**: Follows Clean Architecture layers
4. **Testability**: Much easier to write unit tests

---

## Gradual Migration Strategy

If migrating a larger application:

1. **Start with Domain**: Extract entities first
2. **Add Repositories**: Create interfaces, move data access
3. **Create Use Cases**: Extract business logic
4. **Move Infrastructure**: Separate external services
5. **Update Delivery**: Modify controllers and main
6. **Test**: Ensure everything works
7. **Remove Old Code**: Delete old structure

---

## Conclusion

The migration to Clean Architecture provides:
- ✅ Better code organization
- ✅ Improved testability
- ✅ Greater flexibility
- ✅ Easier maintenance
- ✅ Same functionality

The application is now ready for growth and easier to maintain!
