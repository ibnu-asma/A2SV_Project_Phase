# Task Management REST API

A simple Task Management REST API built with Go and Gin Framework.

## Features

- ✅ Create, Read, Update, Delete (CRUD) operations for tasks
- ✅ In-memory data storage
- ✅ RESTful API design
- ✅ Proper error handling and HTTP status codes
- ✅ JSON request/response format
- ✅ Input validation

## Project Structure

```
task_manager/
├── main.go                      # Application entry point
├── controllers/
│   └── task_controller.go       # HTTP request handlers
├── models/
│   └── task.go                  # Task data structure
├── data/
│   └── task_service.go          # Business logic and data operations
├── router/
│   └── router.go                # Route configuration
├── docs/
│   └── api_documentation.md     # API documentation
└── go.mod                       # Go module dependencies
```

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

## Running the Application

```bash
go run main.go
```

The server will start on `http://localhost:8080`

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | /tasks | Get all tasks |
| GET | /tasks/:id | Get task by ID |
| POST | /tasks | Create new task |
| PUT | /tasks/:id | Update task |
| DELETE | /tasks/:id | Delete task |

## Testing with Postman

See [API Documentation](docs/api_documentation.md) for detailed endpoint specifications and Postman testing instructions.

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

### Get All Tasks
```bash
curl http://localhost:8080/tasks
```

### Get Task by ID
```bash
curl http://localhost:8080/tasks/1
```

### Update Task
```bash
curl -X PUT http://localhost:8080/tasks/1 \
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
curl -X DELETE http://localhost:8080/tasks/1
```

## Technologies Used

- Go 1.x
- Gin Framework
- In-memory storage (map)

## Notes

- Data is stored in-memory and will be lost when the server restarts
- Task IDs are auto-generated
- All dates must be in ISO 8601 format
