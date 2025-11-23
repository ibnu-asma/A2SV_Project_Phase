# Task Management REST API with JWT Authentication

A secure Task Management REST API built with Go, Gin Framework, MongoDB, and JWT authentication.

## Features

- ✅ JWT-based authentication and authorization
- ✅ User registration and login
- ✅ Role-based access control (Admin/User)
- ✅ Password hashing with bcrypt
- ✅ Protected API endpoints
- ✅ MongoDB persistent storage
- ✅ CRUD operations for tasks
- ✅ Admin promotion functionality

## Project Structure

```
task_manager/
├── main.go                      # Application entry point
├── controllers/
│   └── controller.go            # Task and user controllers
├── models/
│   ├── task.go                  # Task model
│   └── user.go                  # User model with auth structures
├── data/
│   ├── task_service.go          # Task business logic
│   └── user_service.go          # User authentication logic
├── middleware/
│   └── auth_middleware.go       # JWT authentication & authorization
├── router/
│   └── router.go                # Route configuration with auth
├── docs/
│   └── api_documentation.md     # Complete API documentation
└── go.mod                       # Dependencies
```

## Prerequisites

- Go 1.16+
- MongoDB 4.0+

## Installation

```bash
cd task_manager
go mod tidy
```

## Running the Application

1. Start MongoDB:
```bash
mongod
```

2. Run the application:
```bash
go run main.go
```

Server starts on: `http://localhost:8080`

## Quick Start Guide

### 1. Register First User (Becomes Admin)
```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"admin123"}'
```

### 2. Login and Get Token
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"admin123"}'
```

Response:
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": "...",
    "username": "admin",
    "role": "admin"
  }
}
```

**Copy the token for subsequent requests!**

### 3. Create Task (Admin Only)
```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -d '{
    "title":"Complete API",
    "description":"Finish authentication",
    "due_date":"2024-12-31T23:59:59Z",
    "status":"In Progress"
  }'
```

### 4. Get All Tasks (Any Authenticated User)
```bash
curl -X GET http://localhost:8080/tasks \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

### 5. Promote User to Admin
```bash
curl -X PUT http://localhost:8080/promote/username \
  -H "Authorization: Bearer ADMIN_TOKEN_HERE"
```

## API Endpoints

### Public Endpoints
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | /register | Register new user |
| POST | /login | Login and get JWT token |

### Protected Endpoints (All Users)
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | /tasks | Get all tasks |
| GET | /tasks/:id | Get task by ID |

### Admin-Only Endpoints
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | /tasks | Create new task |
| PUT | /tasks/:id | Update task |
| DELETE | /tasks/:id | Delete task |
| PUT | /promote/:username | Promote user to admin |

## Authentication

### Token Format
```
Authorization: Bearer <your_jwt_token>
```

### Token Expiration
- JWT tokens expire after 24 hours
- Login again to get a new token

## User Roles

### Admin
- First registered user
- Can create, update, delete tasks
- Can promote other users to admin
- Can view all tasks

### User
- Regular registered users
- Can only view tasks
- Cannot create, update, or delete tasks

## Security Features

- **Password Hashing:** bcrypt with cost 10
- **JWT Tokens:** HS256 algorithm, 24-hour expiration
- **Role-Based Access:** Middleware enforces permissions
- **Secure Storage:** Passwords never stored in plain text

## Testing with Postman

See [API Documentation](docs/api_documentation.md) for detailed Postman testing instructions.

### Quick Test Flow
1. Register admin user
2. Login to get token
3. Create task with admin token
4. Register regular user
5. Login as regular user
6. Try to create task (should fail - 403 Forbidden)
7. View tasks (should succeed)
8. Promote regular user with admin token
9. Login as promoted user
10. Create task (should now succeed)

## MongoDB Configuration

- **URI:** `mongodb://localhost:27017`
- **Database:** `taskdb`
- **Collections:** `tasks`, `users`

## Environment Variables

Update JWT secret in `middleware/auth_middleware.go`:
```go
var jwtSecret = []byte("your-secret-key-change-in-production")
```

**⚠️ Important:** Change the JWT secret before deploying to production!

## Error Responses

| Status | Description |
|--------|-------------|
| 400 | Bad Request - Invalid input |
| 401 | Unauthorized - Missing/invalid token |
| 403 | Forbidden - Insufficient permissions |
| 404 | Not Found - Resource doesn't exist |
| 500 | Internal Server Error |

## Common Issues

### "Authorization header required"
- Add `Authorization: Bearer <token>` header to request

### "Admin access required"
- Endpoint requires admin role
- Login with admin account or get promoted

### "invalid credentials"
- Check username and password
- Ensure user is registered

## Technologies Used

- **Go** - Programming language
- **Gin** - Web framework
- **MongoDB** - Database
- **JWT** - Authentication
- **bcrypt** - Password hashing

## Development

### Install Dependencies
```bash
go get github.com/gin-gonic/gin
go get go.mongodb.org/mongo-driver/mongo
go get github.com/golang-jwt/jwt/v5
go get golang.org/x/crypto/bcrypt
```

### Build
```bash
go build
```

### Run Tests
```bash
go test ./...
```

## License

MIT
