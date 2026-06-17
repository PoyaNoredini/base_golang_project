# BaseProject

A Go REST API backend with **authentication** (JWT) and **authorization** (role-based access control). Built with [Gin](https://github.com/gin-gonic/gin) and [GORM](https://gorm.io), using MySQL as the database.

## Features

- **Authentication**
  - Register with phone, national ID, and password
  - Login with password
  - Login with OTP (one-time password)
  - JWT tokens (24-hour expiry)
- **Authorization**
  - Roles and permissions (RBAC)
  - Permission middleware on protected routes
  - Seeded roles: Administrator, User, Supporter
- **User management**
  - List, create, update, and delete users (permission-gated)
  - Authenticated profile endpoint
- **File upload**
  - Protected file upload with public storage

## Tech Stack

| Layer        | Technology                          |
| ------------ | ----------------------------------- |
| Language     | Go 1.25+                            |
| Web framework| Gin                                 |
| ORM          | GORM                                |
| Database     | MySQL                               |
| Auth         | JWT (`golang-jwt/jwt/v5`)           |
| Config       | Environment variables via `godotenv`|

## Project Structure

```
BaseProject/
├── api/
│   ├── controllers/v01/   # HTTP handlers (auth, user, upload)
│   ├── error/             # Application error types
│   ├── helper/            # JWT, password hashing, OTP helpers
│   ├── middleware/        # Auth & permission middleware
│   ├── resource/          # API response transformers
│   ├── service/           # Business services (file management)
│   └── validations/       # Request validation structs
├── config/                # Database connection & migrations
├── database/seeders/      # Role, permission, and user seeders
├── models/                # GORM models
├── routes/                # Route registration
├── storage/               # Uploaded files (served at /storage)
├── main.go
└── .env                   # Environment configuration (not committed)
```

## Prerequisites

- Go 1.25 or later
- MySQL 5.7+ or 8.x
- Git

## Getting Started

### 1. Clone the repository

```bash
git clone <repository-url>
cd BaseProject
```

### 2. Install dependencies

```bash
go mod download
```

### 3. Configure environment

Create a `.env` file in the project root:

```env
APP_PORT=8080
APP_URL=http://localhost:8080

DB_HOST=127.0.0.1
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=base_project

JWT_SECRET=your-secret-key-change-in-production
```

### 4. Create the database

```sql
CREATE DATABASE base_project CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### 5. Run migrations and seed data

On startup, GORM auto-migrates all models. To seed roles, permissions, and a default admin user:

```bash
go run main.go --seed
```

**Default admin user** (created by seeder):

| Field    | Value              |
| -------- | ------------------ |
| Phone    | `09372718990`      |
| Password | `password`         |
| Role     | Administrator      |

### 6. Start the server

```bash
go run main.go
```

The API will be available at `http://localhost:8080` (or your configured `APP_PORT`).

## Authentication

Protected routes require a JWT in the `Authorization` header:

```
Authorization: Bearer <your-jwt-token>
```

### Auth flow

1. **Register** — `POST /api/v1/auth/register` → returns JWT
2. **Login (password)** — `POST /api/v1/auth/login-with-password` → returns JWT
3. **Login (OTP)**
   - `POST /api/v1/auth/send-otp-code` — request an OTP for a phone number
   - `POST /api/v1/auth/login-with-otp` — verify OTP and receive JWT

Tokens are signed with `JWT_SECRET` and expire after 24 hours.

## Authorization

Access to sensitive endpoints is controlled by **permissions** assigned to **roles**, and roles assigned to **users**.

```
users → user_roles → roles → role_permissions → permissions
```

The `RequirePermission` middleware checks whether the authenticated user has the required permission before allowing the request.

### Seeded permissions (Administrator)

- `create-user`, `view-user`, `update-user`, `delete-user`, `deactive-user`
- `create-role`, `view-role`, `update-role`, `delete-role`
- `create-permission`, `view-permission`, `update-permission`, `delete-permission`

## API Endpoints

### Public (no token)

| Method | Endpoint                          | Description              |
| ------ | --------------------------------- | ------------------------ |
| POST   | `/api/v1/auth/register`           | Register a new user      |
| POST   | `/api/v1/auth/login-with-password`| Login with phone/password|
| POST   | `/api/v1/auth/send-otp-code`      | Send OTP to phone        |
| POST   | `/api/v1/auth/login-with-otp`     | Login with OTP           |

### Protected (JWT required)

| Method | Endpoint                        | Permission     | Description              |
| ------ | ------------------------------- | -------------- | ------------------------ |
| GET    | `/api/v1/users/profile`         | —              | Current user profile     |
| GET    | `/api/v1/users`                 | `view-user`    | List all users           |
| POST   | `/api/v1/users`                 | `create-user`  | Create a user            |
| PUT    | `/api/v1/users/update`          | `update-user`  | Update own profile       |
| PUT    | `/api/v1/users/admin/update/:id`| `update-user`  | Admin update user by ID  |
| DELETE | `/api/v1/users/delete/:id`      | `delete-user`  | Delete user by ID        |
| POST   | `/api/v1/upload-file`           | —              | Upload a file            |

Static files are served at `/storage`.

## Request & Response Format

### Success response

```json
{
  "status": true,
  "message": "success",
  "data": { }
}
```

### Error response

```json
{
  "status": false,
  "message": "Error description"
}
```

Auth middleware errors use:

```json
{
  "code": 401,
  "message": "Authorization token is required"
}
```

### Example: Login with password

**Request**

```bash
curl -X POST http://localhost:8080/api/v1/auth/login-with-password \
  -H "Content-Type: application/json" \
  -d '{"phone": "09372718990", "password": "password"}'
```

**Response**

```json
{
  "status": true,
  "message": "success",
  "data": {
    "message": "Login successful",
    "token": "eyJhbGciOiJIUzI1NiIs...",
    "user": {
      "id": 1,
      "first_name": "Admin",
      "last_name": "Admin",
      "phone_number": "",
      "email": "admin@example.com",
      "roles": [{ "id": 1, "title": "Administrator" }]
    }
  }
}
```

### Example: Protected request

```bash
curl http://localhost:8080/api/v1/users/profile \
  -H "Authorization: Bearer <token>"
```

## Development Notes

- Phone numbers must be **11 characters** (validated on auth endpoints).
- Passwords are hashed with bcrypt before storage.
- OTP codes expire after **2 minutes** and are invalidated after use.
- Uploaded files are stored under `storage/` and exposed via `/storage`.

## License

This project is provided as a learning base. Add your own license as needed.
