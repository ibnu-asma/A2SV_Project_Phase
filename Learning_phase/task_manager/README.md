# Task Management REST API with MongoDB

A Task Management REST API built with Go, Gin Framework, and MongoDB for persistent data storage.

## Features

- ✅ Create, Read, Update, Delete (CRUD) operations for tasks
- ✅ MongoDB persistent data storage
- ✅ RESTful API design
- ✅ Proper error handling and HTTP status codes
- ✅ JSON request/response format
- ✅ Input validation
- ✅ MongoDB ObjectID support

## Project Structure

```
task_manager/
├── main.go                      # Application entry point with MongoDB connection
├── controllers/
│   └── task_controller.go       # HTTP request handlers
├── models/
│   └── task.go                  # Task data structure with BSON tags
├── data/
│   └── task_service.go          # Business logic with MongoDB operations
├── router/
│   └── router.go                # Route configuration
├── docs/
│   └── api_documentation.md     # Comprehensive API documentation
└── go.mod                       # Go module dependencies
```

## Prerequisites

- Go 1.16 or higher
- MongoDB 4.0 or higher (local or Atlas)

## Installation

1. Clone the repository
2. Navigate to the project directory:
```bash
cd task_manager
```

3. Install dependencies:
```bash
go mod tidy
```

4. Ensure MongoDB is running:
```bash
mongod
```

## MongoDB Setup

### Local MongoDB
1. Install MongoDB from [mongodb.com](https://www.mongodb.com/try/download/community)
2. Start MongoDB service:
```bash
mongod
```

### MongoDB Atlas (Cloud)
1. Create account at [MongoDB Atlas](https://www.mongodb.com/cloud/atlas)
2. Create a cluster
3. Get connection string
4. Update `main.go` with your connection string:
```go
mongoURI := "mongodb+srv://username:password@cluster.mongodb.net/?retryWrites=true&w=majority"
```

## Configuration

Default MongoDB settings in `main.go`:
- **URI:** `mongodb://localhost:27017`
- **Database:** `taskdb`
- **Collection:** `tasks`

To customize, update:
```go
// main.go
mongoURI := "your_connection_string"

// router/router.go
service := data.NewTaskService(client, "your_db", "your_collection")
```

## Running the Application

```bash
go run main.go
```

Expected output:
```
Connected to MongoDB successfully!
[GIN-debug] Listening and serving HTTP on :8080
```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | /tasks | Get all tasks |
| GET | /tasks/:id | Get task by MongoDB ObjectID |
| POST | /tasks | Create new task |
| PUT | /tasks/:id | Update task |
| DELETE | /tasks/:id | Delete task |

## Example Usage

### Create a Task
```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Learn Go",
    "description": "Complete Go tutorial",
    "due_date": "2024-12-31T23:59:59Z",
    "status": "Pending"
  }'
```

Response:
```json
{
  "id": "507f1f77bcf86cd799439011",
  "title": "Learn Go",
  "description": "Complete Go tutorial",
  "due_date": "2024-12-31T23:59:59Z",
  "status": "Pending"
}
```

### Get All Tasks
```bash
curl http://localhost:8080/tasks
```

### Get Task by ID
```bash
curl http://localhost:8080/tasks/507f1f77bcf86cd799439011
```

### Update Task
```bash
curl -X PUT http://localhost:8080/tasks/507f1f77bcf86cd799439011 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Learn Go - Updated",
    "description": "Complete advanced Go tutorial",
    "due_date": "2024-12-31T23:59:59Z",
    "status": "In Progress"
  }'
```

### Delete Task
```bash
curl -X DELETE http://localhost:8080/tasks/507f1f77bcf86cd799439011
```

## Testing with Postman

See [API Documentation](docs/api_documentation.md) for detailed Postman testing instructions.

## Verifying Data in MongoDB

### MongoDB Shell
```bash
mongosh
use taskdb
db.tasks.find().pretty()
```

### MongoDB Compass
1. Connect to `mongodb://localhost:27017`
2. Navigate to `taskdb` → `tasks`
3. View stored documents

## Technologies Used

- **Go** - Programming language
- **Gin Framework** - Web framework
- **MongoDB** - NoSQL database
- **MongoDB Go Driver** - Official MongoDB driver

## Key Changes from In-Memory Version

1. **Persistent Storage:** Data survives application restarts
2. **ObjectID:** MongoDB ObjectIDs instead of string IDs
3. **BSON Tags:** Model fields tagged for MongoDB serialization
4. **Error Handling:** Enhanced error handling for database operations
5. **Context Timeouts:** 10-second timeout for all database operations

## Error Handling

- **400 Bad Request:** Invalid input or ID format
- **404 Not Found:** Task doesn't exist
- **500 Internal Server Error:** Database connection or operation errors

## Notes

- Task IDs are MongoDB ObjectIDs (24-character hex strings)
- All dates must be in ISO 8601 format
- Connection timeout: 10 seconds
- API is backward compatible (same endpoint structure)
- Data persists across restarts

## Troubleshooting

### MongoDB Connection Failed
- Ensure MongoDB is running: `mongod`
- Check connection string in `main.go`
- Verify MongoDB port (default: 27017)

### Invalid Task ID Error
- Ensure ID is valid MongoDB ObjectID (24-char hex)
- Copy ID from create/get response

## License

MIT
