# Task Management API with JWT Authentication

## Base URL
```
http://localhost:8080
```

## Authentication

This API uses JWT (JSON Web Tokens) for authentication. Protected endpoints require a valid JWT token in the Authorization header.

### Token Format
```
Authorization: Bearer <your_jwt_token>
```

---

## User Roles

- **Admin**: Can create, update, delete tasks, and promote users
- **User**: Can view all tasks and individual task details

### First User Rule
The first registered user automatically becomes an admin.

---

## Public Endpoints (No Authentication Required)

### 1. Register User
**Endpoint:** `POST /register`

**Description:** Create a new user account. First user becomes admin automatically.

**Request Body:**
```json
{
  "username": "john_doe",
  "password": "securePassword123"
}
```

**Response (201 Created):**
```json
{
  "id": "507f1f77bcf86cd799439011",
  "username": "john_doe",
  "role": "admin"
}
```

**Error Responses:**
- **400 Bad Request:** Username already exists or validation error

---

### 2. Login
**Endpoint:** `POST /login`

**Description:** Authenticate user and receive JWT token.

**Request Body:**
```json
{
  "username": "john_doe",
  "password": "securePassword123"
}
```

**Response (200 OK):**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": "507f1f77bcf86cd799439011",
    "username": "john_doe",
    "role": "admin"
  }
}
```

**Error Responses:**
- **401 Unauthorized:** Invalid credentials

---

## Protected Endpoints (Authentication Required)

### 3. Get All Tasks
**Endpoint:** `GET /tasks`

**Description:** Retrieve all tasks (accessible by all authenticated users).

**Headers:**
```
Authorization: Bearer <your_jwt_token>
```

**Response (200 OK):**
```json
{
  "tasks": [
    {
      "id": "507f1f77bcf86cd799439011",
      "title": "Complete project",
      "description": "Finish the task management API",
      "due_date": "2024-12-31T23:59:59Z",
      "status": "In Progress"
    }
  ]
}
```

**Error Responses:**
- **401 Unauthorized:** Missing or invalid token

---

### 4. Get Task by ID
**Endpoint:** `GET /tasks/:id`

**Description:** Retrieve specific task details (accessible by all authenticated users).

**Headers:**
```
Authorization: Bearer <your_jwt_token>
```

**Response (200 OK):**
```json
{
  "id": "507f1f77bcf86cd799439011",
  "title": "Complete project",
  "description": "Finish the task management API",
  "due_date": "2024-12-31T23:59:59Z",
  "status": "In Progress"
}
```

**Error Responses:**
- **401 Unauthorized:** Missing or invalid token
- **404 Not Found:** Task not found

---

## Admin-Only Endpoints

### 5. Create Task
**Endpoint:** `POST /tasks`

**Description:** Create a new task (admin only).

**Headers:**
```
Authorization: Bearer <admin_jwt_token>
```

**Request Body:**
```json
{
  "title": "New Task",
  "description": "Task description",
  "due_date": "2024-12-31T23:59:59Z",
  "status": "Pending"
}
```

**Response (201 Created):**
```json
{
  "id": "507f1f77bcf86cd799439011",
  "title": "New Task",
  "description": "Task description",
  "due_date": "2024-12-31T23:59:59Z",
  "status": "Pending"
}
```

**Error Responses:**
- **401 Unauthorized:** Missing or invalid token
- **403 Forbidden:** User is not an admin

---

### 6. Update Task
**Endpoint:** `PUT /tasks/:id`

**Description:** Update an existing task (admin only).

**Headers:**
```
Authorization: Bearer <admin_jwt_token>
```

**Request Body:**
```json
{
  "title": "Updated Task",
  "description": "Updated description",
  "due_date": "2024-12-31T23:59:59Z",
  "status": "Completed"
}
```

**Response (200 OK):**
```json
{
  "id": "507f1f77bcf86cd799439011",
  "title": "Updated Task",
  "description": "Updated description",
  "due_date": "2024-12-31T23:59:59Z",
  "status": "Completed"
}
```

**Error Responses:**
- **401 Unauthorized:** Missing or invalid token
- **403 Forbidden:** User is not an admin
- **404 Not Found:** Task not found

---

### 7. Delete Task
**Endpoint:** `DELETE /tasks/:id`

**Description:** Delete a task (admin only).

**Headers:**
```
Authorization: Bearer <admin_jwt_token>
```

**Response (200 OK):**
```json
{
  "message": "Task deleted successfully"
}
```

**Error Responses:**
- **401 Unauthorized:** Missing or invalid token
- **403 Forbidden:** User is not an admin
- **404 Not Found:** Task not found

---

### 8. Promote User to Admin
**Endpoint:** `PUT /promote/:username`

**Description:** Promote a user to admin role (admin only).

**Headers:**
```
Authorization: Bearer <admin_jwt_token>
```

**URL Parameters:**
- `username`: Username of the user to promote

**Example:** `PUT /promote/jane_doe`

**Response (200 OK):**
```json
{
  "message": "User promoted to admin successfully"
}
```

**Error Responses:**
- **401 Unauthorized:** Missing or invalid token
- **403 Forbidden:** User is not an admin
- **404 Not Found:** User not found

---

## Status Codes Summary

| Status Code | Description |
|-------------|-------------|
| 200 OK | Request successful |
| 201 Created | Resource created successfully |
| 400 Bad Request | Invalid request payload |
| 401 Unauthorized | Missing or invalid authentication token |
| 403 Forbidden | Insufficient permissions |
| 404 Not Found | Resource not found |
| 500 Internal Server Error | Server error |

---

## Security Features

### Password Hashing
- Passwords are hashed using bcrypt with default cost (10)
- Plain text passwords are never stored in the database

### JWT Token
- Tokens expire after 24 hours
- Tokens contain user ID, username, and role
- Signed with HS256 algorithm

### Authorization
- Middleware validates JWT tokens on protected routes
- Admin middleware checks user role for admin-only endpoints

---

## Testing with Postman

### Step 1: Register First User (Admin)
```
POST http://localhost:8080/register
Body (JSON):
{
  "username": "admin",
  "password": "admin123"
}
```
**Note:** First user becomes admin automatically.

### Step 2: Login
```
POST http://localhost:8080/login
Body (JSON):
{
  "username": "admin",
  "password": "admin123"
}
```
**Copy the token from response.**

### Step 3: Create Task (Admin)
```
POST http://localhost:8080/tasks
Headers:
  Authorization: Bearer <your_token>
Body (JSON):
{
  "title": "Test Task",
  "description": "Testing authentication",
  "due_date": "2024-12-31T23:59:59Z",
  "status": "Pending"
}
```

### Step 4: Get All Tasks (Any User)
```
GET http://localhost:8080/tasks
Headers:
  Authorization: Bearer <your_token>
```

### Step 5: Register Regular User
```
POST http://localhost:8080/register
Body (JSON):
{
  "username": "user1",
  "password": "user123"
}
```
**Note:** This user will have "user" role.

### Step 6: Promote User (Admin Only)
```
PUT http://localhost:8080/promote/user1
Headers:
  Authorization: Bearer <admin_token>
```

---

## Testing Scenarios

### Scenario 1: Unauthorized Access
Try accessing protected endpoint without token:
```
GET http://localhost:8080/tasks
(No Authorization header)
```
**Expected:** 401 Unauthorized

### Scenario 2: Regular User Creating Task
Login as regular user and try to create task:
```
POST http://localhost:8080/tasks
Headers:
  Authorization: Bearer <user_token>
```
**Expected:** 403 Forbidden

### Scenario 3: Token Expiration
Use a token after 24 hours:
**Expected:** 401 Unauthorized

---

## MongoDB Collections

### Users Collection
```json
{
  "_id": ObjectId,
  "username": "string",
  "password": "hashed_password",
  "role": "admin|user"
}
```

### Tasks Collection
```json
{
  "_id": ObjectId,
  "title": "string",
  "description": "string",
  "due_date": ISODate,
  "status": "string"
}
```

---

## Environment Configuration

### JWT Secret
Update in `middleware/auth_middleware.go`:
```go
var jwtSecret = []byte("your-secret-key-change-in-production")
```

**Important:** Change this in production!

---

## Running the Application

1. Start MongoDB:
```bash
mongod
```

2. Run the application:
```bash
go run main.go
```

3. Server starts on: `http://localhost:8080`

---

## Common Errors

### "Authorization header required"
- Add `Authorization: Bearer <token>` header

### "Invalid token format"
- Ensure format is: `Bearer <token>` (with space)

### "Admin access required"
- Endpoint requires admin role
- Login with admin account or get promoted

### "username already exists"
- Choose a different username

### "invalid credentials"
- Check username and password

---

## Best Practices

1. **Store tokens securely** on client side
2. **Never share** JWT secret key
3. **Use HTTPS** in production
4. **Implement token refresh** for better UX
5. **Log out** by deleting token on client side
6. **Change default JWT secret** before deployment

---

## Notes

- Passwords are hashed with bcrypt (cost 10)
- JWT tokens expire in 24 hours
- First registered user is automatically admin
- Admin can promote any user to admin
- Regular users can only view tasks
- Admins can perform all CRUD operations
