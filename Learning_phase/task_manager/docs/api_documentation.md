# Task Management API Documentation

## Base URL
```
http://localhost:8080
```

## Endpoints

### 1. Get All Tasks
**Endpoint:** `GET /tasks`

**Description:** Retrieves a list of all tasks.

**Request:** No body required

**Response:**
- **Status Code:** 200 OK
- **Body:**
```json
{
  "tasks": [
    {
      "id": "1",
      "title": "Complete project",
      "description": "Finish the task management API",
      "due_date": "2024-12-31T23:59:59Z",
      "status": "In Progress"
    }
  ]
}
```

---

### 2. Get Task by ID
**Endpoint:** `GET /tasks/:id`

**Description:** Retrieves details of a specific task by ID.

**Request:** No body required

**URL Parameters:**
- `id` (string): Task ID

**Response:**
- **Status Code:** 200 OK
- **Body:**
```json
{
  "id": "1",
  "title": "Complete project",
  "description": "Finish the task management API",
  "due_date": "2024-12-31T23:59:59Z",
  "status": "In Progress"
}
```

**Error Response:**
- **Status Code:** 404 Not Found
- **Body:**
```json
{
  "error": "Task not found"
}
```

---

### 3. Create Task
**Endpoint:** `POST /tasks`

**Description:** Creates a new task.

**Request Body:**
```json
{
  "title": "Complete project",
  "description": "Finish the task management API",
  "due_date": "2024-12-31T23:59:59Z",
  "status": "In Progress"
}
```

**Required Fields:**
- `title` (string): Task title
- `description` (string): Task description
- `due_date` (string): Due date in ISO 8601 format
- `status` (string): Task status (e.g., "Pending", "In Progress", "Completed")

**Response:**
- **Status Code:** 201 Created
- **Body:**
```json
{
  "id": "1",
  "title": "Complete project",
  "description": "Finish the task management API",
  "due_date": "2024-12-31T23:59:59Z",
  "status": "In Progress"
}
```

**Error Response:**
- **Status Code:** 400 Bad Request
- **Body:**
```json
{
  "error": "validation error message"
}
```

---

### 4. Update Task
**Endpoint:** `PUT /tasks/:id`

**Description:** Updates an existing task.

**URL Parameters:**
- `id` (string): Task ID

**Request Body:**
```json
{
  "title": "Updated title",
  "description": "Updated description",
  "due_date": "2024-12-31T23:59:59Z",
  "status": "Completed"
}
```

**Required Fields:**
- `title` (string): Task title
- `description` (string): Task description
- `due_date` (string): Due date in ISO 8601 format
- `status` (string): Task status

**Response:**
- **Status Code:** 200 OK
- **Body:**
```json
{
  "id": "1",
  "title": "Updated title",
  "description": "Updated description",
  "due_date": "2024-12-31T23:59:59Z",
  "status": "Completed"
}
```

**Error Responses:**
- **Status Code:** 400 Bad Request
```json
{
  "error": "validation error message"
}
```
- **Status Code:** 404 Not Found
```json
{
  "error": "Task not found"
}
```

---

### 5. Delete Task
**Endpoint:** `DELETE /tasks/:id`

**Description:** Deletes a specific task.

**URL Parameters:**
- `id` (string): Task ID

**Request:** No body required

**Response:**
- **Status Code:** 200 OK
- **Body:**
```json
{
  "message": "Task deleted successfully"
}
```

**Error Response:**
- **Status Code:** 404 Not Found
- **Body:**
```json
{
  "error": "Task not found"
}
```

---

## Status Codes Summary

| Status Code | Description |
|-------------|-------------|
| 200 OK | Request successful |
| 201 Created | Resource created successfully |
| 400 Bad Request | Invalid request payload |
| 404 Not Found | Resource not found |

---

## Testing with Postman

### Setup
1. Import the API endpoints into Postman
2. Set base URL: `http://localhost:8080`
3. Ensure the server is running before testing

### Test Scenarios

#### 1. Create a Task
- Method: POST
- URL: `http://localhost:8080/tasks`
- Headers: `Content-Type: application/json`
- Body (raw JSON):
```json
{
  "title": "Learn Go",
  "description": "Complete Go tutorial",
  "due_date": "2024-12-31T23:59:59Z",
  "status": "Pending"
}
```

#### 2. Get All Tasks
- Method: GET
- URL: `http://localhost:8080/tasks`

#### 3. Get Task by ID
- Method: GET
- URL: `http://localhost:8080/tasks/1`

#### 4. Update Task
- Method: PUT
- URL: `http://localhost:8080/tasks/1`
- Headers: `Content-Type: application/json`
- Body (raw JSON):
```json
{
  "title": "Learn Go - Updated",
  "description": "Complete advanced Go tutorial",
  "due_date": "2024-12-31T23:59:59Z",
  "status": "In Progress"
}
```

#### 5. Delete Task
- Method: DELETE
- URL: `http://localhost:8080/tasks/1`

---

## Running the Application

1. Install dependencies:
```bash
go mod tidy
```

2. Run the application:
```bash
go run main.go
```

3. The server will start on `http://localhost:8080`

---

## Notes

- All dates should be in ISO 8601 format (e.g., `2024-12-31T23:59:59Z`)
- Task IDs are auto-generated
- Data is stored in-memory and will be lost when the server restarts
- All endpoints return JSON responses
