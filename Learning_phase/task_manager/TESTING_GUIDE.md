# JWT Authentication Testing Guide

## Prerequisites
1. MongoDB running on `localhost:27017`
2. Application running on `localhost:8080`
3. Postman installed

## Test Sequence

### Test 1: Register First User (Admin)

**Request:**
```
POST http://localhost:8080/register
Content-Type: application/json

{
  "username": "admin",
  "password": "admin123"
}
```

**Expected Response (201):**
```json
{
  "id": "...",
  "username": "admin",
  "role": "admin"
}
```

✅ **Verify:** First user has "admin" role

---

### Test 2: Login as Admin

**Request:**
```
POST http://localhost:8080/login
Content-Type: application/json

{
  "username": "admin",
  "password": "admin123"
}
```

**Expected Response (200):**
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

✅ **Save the token** for next requests

---

### Test 3: Create Task (Admin)

**Request:**
```
POST http://localhost:8080/tasks
Authorization: Bearer <admin_token>
Content-Type: application/json

{
  "title": "Test Task",
  "description": "Testing authentication",
  "due_date": "2024-12-31T23:59:59Z",
  "status": "Pending"
}
```

**Expected Response (201):**
```json
{
  "id": "...",
  "title": "Test Task",
  "description": "Testing authentication",
  "due_date": "2024-12-31T23:59:59Z",
  "status": "Pending"
}
```

✅ **Verify:** Admin can create tasks

---

### Test 4: Get All Tasks (Admin)

**Request:**
```
GET http://localhost:8080/tasks
Authorization: Bearer <admin_token>
```

**Expected Response (200):**
```json
{
  "tasks": [...]
}
```

✅ **Verify:** Admin can view tasks

---

### Test 5: Register Regular User

**Request:**
```
POST http://localhost:8080/register
Content-Type: application/json

{
  "username": "user1",
  "password": "user123"
}
```

**Expected Response (201):**
```json
{
  "id": "...",
  "username": "user1",
  "role": "user"
}
```

✅ **Verify:** Second user has "user" role

---

### Test 6: Login as Regular User

**Request:**
```
POST http://localhost:8080/login
Content-Type: application/json

{
  "username": "user1",
  "password": "user123"
}
```

**Expected Response (200):**
```json
{
  "token": "...",
  "user": {
    "id": "...",
    "username": "user1",
    "role": "user"
  }
}
```

✅ **Save the user token**

---

### Test 7: Try to Create Task as Regular User (Should Fail)

**Request:**
```
POST http://localhost:8080/tasks
Authorization: Bearer <user_token>
Content-Type: application/json

{
  "title": "Unauthorized Task",
  "description": "This should fail",
  "due_date": "2024-12-31T23:59:59Z",
  "status": "Pending"
}
```

**Expected Response (403):**
```json
{
  "error": "Admin access required"
}
```

✅ **Verify:** Regular users cannot create tasks

---

### Test 8: Get Tasks as Regular User (Should Succeed)

**Request:**
```
GET http://localhost:8080/tasks
Authorization: Bearer <user_token>
```

**Expected Response (200):**
```json
{
  "tasks": [...]
}
```

✅ **Verify:** Regular users can view tasks

---

### Test 9: Promote User to Admin

**Request:**
```
PUT http://localhost:8080/promote/user1
Authorization: Bearer <admin_token>
```

**Expected Response (200):**
```json
{
  "message": "User promoted to admin successfully"
}
```

✅ **Verify:** Admin can promote users

---

### Test 10: Login as Promoted User

**Request:**
```
POST http://localhost:8080/login
Content-Type: application/json

{
  "username": "user1",
  "password": "user123"
}
```

**Expected Response (200):**
```json
{
  "token": "...",
  "user": {
    "id": "...",
    "username": "user1",
    "role": "admin"
  }
}
```

✅ **Verify:** User now has "admin" role

---

### Test 11: Create Task as Promoted User (Should Succeed)

**Request:**
```
POST http://localhost:8080/tasks
Authorization: Bearer <new_admin_token>
Content-Type: application/json

{
  "title": "Promoted User Task",
  "description": "Created by promoted admin",
  "due_date": "2024-12-31T23:59:59Z",
  "status": "Pending"
}
```

**Expected Response (201):**
```json
{
  "id": "...",
  "title": "Promoted User Task",
  ...
}
```

✅ **Verify:** Promoted user can now create tasks

---

### Test 12: Access Without Token (Should Fail)

**Request:**
```
GET http://localhost:8080/tasks
(No Authorization header)
```

**Expected Response (401):**
```json
{
  "error": "Authorization header required"
}
```

✅ **Verify:** Protected endpoints require authentication

---

### Test 13: Access With Invalid Token (Should Fail)

**Request:**
```
GET http://localhost:8080/tasks
Authorization: Bearer invalid_token_here
```

**Expected Response (401):**
```json
{
  "error": "Invalid token"
}
```

✅ **Verify:** Invalid tokens are rejected

---

### Test 14: Update Task (Admin Only)

**Request:**
```
PUT http://localhost:8080/tasks/<task_id>
Authorization: Bearer <admin_token>
Content-Type: application/json

{
  "title": "Updated Task",
  "description": "Updated by admin",
  "due_date": "2024-12-31T23:59:59Z",
  "status": "Completed"
}
```

**Expected Response (200):**
```json
{
  "id": "...",
  "title": "Updated Task",
  ...
}
```

✅ **Verify:** Admin can update tasks

---

### Test 15: Delete Task (Admin Only)

**Request:**
```
DELETE http://localhost:8080/tasks/<task_id>
Authorization: Bearer <admin_token>
```

**Expected Response (200):**
```json
{
  "message": "Task deleted successfully"
}
```

✅ **Verify:** Admin can delete tasks

---

## Summary Checklist

- [ ] First user becomes admin automatically
- [ ] Subsequent users have "user" role
- [ ] Login returns JWT token
- [ ] Admin can create, update, delete tasks
- [ ] Regular users can only view tasks
- [ ] Admin can promote users
- [ ] Promoted users gain admin privileges
- [ ] Protected endpoints require valid token
- [ ] Invalid tokens are rejected
- [ ] Missing tokens return 401
- [ ] Non-admin access to admin endpoints returns 403

## Verify in MongoDB

```bash
mongosh
use taskdb
db.users.find().pretty()
db.tasks.find().pretty()
```

Check that:
- Passwords are hashed (not plain text)
- First user has role "admin"
- Other users have role "user"
- Promoted users have role "admin"
