# Task Management API Documentation

## Base URL
```
http://localhost:8080
```

## MongoDB Configuration

### Prerequisites
- MongoDB installed locally or access to MongoDB Atlas
- Default connection: `mongodb://localhost:27017`
- Database: `taskdb`
- Collection: `tasks`

### Environment Setup
1. Install MongoDB locally or use MongoDB Atlas
2. Start MongoDB service: `mongod` (for local installation)
3. Update connection string in `main.go` if using custom configuration

---

## Endpoints

### 1. Get All Tasks
**Endpoint:** `GET /tasks`

**Description:** Retrieves a list of all tasks from MongoDB.

**Request:** No body required

**Response:**
- **Status Code:** 200 OK
- **Body:**
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

**Error Response:**
- **Status Code:** 500 Internal Server Error
- **Body:**
```json
{
  "error": "database error message"
}
```

---

### 2. Get Task by ID
**Endpoint:** `GET /tasks/:id`

**Description:** Retrieves details of a specific task by MongoDB ObjectID.

**Request:** No body required

**URL Parameters:**
- `id` (string): MongoDB ObjectID (24-character hex string)

**Example:** `GET /tasks/507f1f77bcf86cd799439011`

**Response:**
- **Status Code:** 200 OK
- **Body:**
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
- **Status Code:** 400 Bad Request (Invalid ID format)
```json
{
  "error": "invalid task ID"
}
```
- **Status Code:** 404 Not Found
```json
{
  "error": "Task not found"
}
```

---

### 3. Create Task
**Endpoint:** `POST /tasks`

**Description:** Creates a new task in MongoDB.

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
  "id": "507f1f77bcf86cd799439011",
  "title": "Complete project",
  "description": "Finish the task management API",
  "due_date": "2024-12-31T23:59:59Z",
  "status": "In Progress"
}
```

**Error Responses:**
- **Status Code:** 400 Bad Request
```json
{
  "error": "validation error message"
}
```
- **Status Code:** 500 Internal Server Error
```json
{
  "error": "database error message"
}
```

---

### 4. Update Task
**Endpoint:** `PUT /tasks/:id`

**Description:** Updates an existing task in MongoDB.

**URL Parameters:**
- `id` (string): MongoDB ObjectID

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
  "id": "507f1f77bcf86cd799439011",
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
- **Status Code:** 500 Internal Server Error
```json
{
  "error": "database error message"
}
```

---

### 5. Delete Task
**Endpoint:** `DELETE /tasks/:id`

**Description:** Deletes a specific task from MongoDB.

**URL Parameters:**
- `id` (string): MongoDB ObjectID

**Request:** No body required

**Response:**
- **Status Code:** 200 OK
- **Body:**
```json
{
  "message": "Task deleted successfully"
}
```

**Error Responses:**
- **Status Code:** 400 Bad Request
```json
{
  "error": "invalid task ID"
}
```
- **Status Code:** 404 Not Found
```json
{
  "error": "Task not found"
}
```
- **Status Code:** 500 Internal Server Error
```json
{
  "error": "database error message"
}
```

---

## Status Codes Summary

| Status Code | Description |
|-------------|-------------|
| 200 OK | Request successful |
| 201 Created | Resource created successfully |
| 400 Bad Request | Invalid request payload or ID format |
| 404 Not Found | Resource not found |
| 500 Internal Server Error | Database or server error |

---

## MongoDB Integration Details

### Connection Configuration
- **Default URI:** `mongodb://localhost:27017`
- **Database Name:** `taskdb`
- **Collection Name:** `tasks`

### Data Persistence
- All tasks are stored in MongoDB with persistent storage
- Task IDs are MongoDB ObjectIDs (24-character hex strings)
- Data survives application restarts

### MongoDB Operations
- **Create:** `InsertOne` operation
- **Read:** `Find` and `FindOne` operations
- **Update:** `UpdateOne` operation with `$set`
- **Delete:** `DeleteOne` operation

---

## Testing with Postman

### Setup
1. Ensure MongoDB is running
2. Start the application: `go run main.go`
3. Import endpoints into Postman
4. Set base URL: `http://localhost:8080`

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
- URL: `http://localhost:8080/tasks/507f1f77bcf86cd799439011`
- Note: Replace with actual ObjectID from created task

#### 4. Update Task
- Method: PUT
- URL: `http://localhost:8080/tasks/507f1f77bcf86cd799439011`
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
- URL: `http://localhost:8080/tasks/507f1f77bcf86cd799439011`

---

## Verifying Data in MongoDB

### Using MongoDB Shell
```bash
mongosh
use taskdb
db.tasks.find().pretty()
```

### Using MongoDB Compass
1. Connect to `mongodb://localhost:27017`
2. Navigate to `taskdb` database
3. Open `tasks` collection
4. View and verify stored documents

---

## Running the Application

1. Install MongoDB:
   - **Local:** Download from [mongodb.com](https://www.mongodb.com/try/download/community)
   - **Cloud:** Use [MongoDB Atlas](https://www.mongodb.com/cloud/atlas)

2. Start MongoDB service:
```bash
mongod
```

3. Install dependencies:
```bash
go mod tidy
```

4. Run the application:
```bash
go run main.go
```

5. The server will start on `http://localhost:8080`

---

## Configuration Options

### Custom MongoDB URI
Update `main.go`:
```go
mongoURI := "mongodb://username:password@host:port"
```

### MongoDB Atlas Connection
```go
mongoURI := "mongodb+srv://username:password@cluster.mongodb.net/?retryWrites=true&w=majority"
```

### Custom Database/Collection Names
Update `router/router.go`:
```go
service := data.NewTaskService(client, "your_db_name", "your_collection_name")
```

---

## Notes

- Task IDs are MongoDB ObjectIDs (24-character hex strings)
- All dates must be in ISO 8601 format
- Data persists across application restarts
- Connection timeout is set to 10 seconds
- All endpoints return JSON responses
- Backward compatible with previous API version (same endpoint structure)
