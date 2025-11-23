# Quick Setup Guide

## Step 1: Install MongoDB

### Windows
1. Download MongoDB Community Server from [mongodb.com](https://www.mongodb.com/try/download/community)
2. Run the installer
3. Add MongoDB to PATH: `C:\Program Files\MongoDB\Server\7.0\bin`

### macOS
```bash
brew tap mongodb/brew
brew install mongodb-community
```

### Linux (Ubuntu)
```bash
wget -qO - https://www.mongodb.org/static/pgp/server-7.0.asc | sudo apt-key add -
echo "deb [ arch=amd64,arm64 ] https://repo.mongodb.org/apt/ubuntu focal/mongodb-org/7.0 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-7.0.list
sudo apt-get update
sudo apt-get install -y mongodb-org
```

## Step 2: Start MongoDB

### Windows
```bash
mongod
```

### macOS/Linux
```bash
brew services start mongodb-community
# or
sudo systemctl start mongod
```

## Step 3: Verify MongoDB is Running

```bash
mongosh
```

You should see MongoDB shell prompt.

## Step 4: Install Go Dependencies

```bash
cd task_manager
go mod tidy
```

## Step 5: Run the Application

```bash
go run main.go
```

Expected output:
```
Connected to MongoDB successfully!
[GIN-debug] Listening and serving HTTP on :8080
```

## Step 6: Test the API

### Create a task
```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Test Task","description":"Testing MongoDB","due_date":"2024-12-31T23:59:59Z","status":"Pending"}'
```

### Get all tasks
```bash
curl http://localhost:8080/tasks
```

## Troubleshooting

### Error: "Failed to connect to MongoDB"
- Ensure MongoDB is running: `mongod`
- Check if port 27017 is available
- Verify connection string in `main.go`

### Error: "Failed to ping MongoDB"
- MongoDB service might not be started
- Check firewall settings
- Verify MongoDB is listening on localhost:27017

### Error: "invalid task ID"
- Ensure you're using valid MongoDB ObjectID (24-character hex)
- Copy ID from create response

## Using MongoDB Atlas (Cloud)

1. Create account at [MongoDB Atlas](https://www.mongodb.com/cloud/atlas)
2. Create a free cluster
3. Add your IP to whitelist
4. Create database user
5. Get connection string
6. Update `main.go`:
```go
mongoURI := "mongodb+srv://username:password@cluster.mongodb.net/?retryWrites=true&w=majority"
```

## Next Steps

- Test all endpoints with Postman
- View data in MongoDB Compass
- Read full API documentation in `docs/api_documentation.md`
