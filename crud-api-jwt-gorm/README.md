# CRUD API with JWT Authentication (Go + GORM + MySQL)

This project is a simple RESTful API written in Go, featuring user authentication using JWT, MySQL database integration via GORM, and protected CRUD operations for personal Todo items.

## Features

- User registration and login
- JWT-based authentication
- GORM + MySQL integration
- Full CRUD support for Todo items
- Middleware for protected routes
- Modular and clean folder structure

---

## API Endpoints

### Authentication
| Method | Endpoint   | Description        |
|--------|------------|--------------------|
| POST   | `/signup`  | Register a new user |
| POST   | `/login`   | Authenticate and get JWT |

### Todos (Requires JWT)
| Method | Endpoint        | Description               |
|--------|------------------|---------------------------|
| GET    | `/todos`         | Get user's todo list      |
| POST   | `/todos`         | Create a new todo         |
| PUT    | `/todos/{id}`    | Update a todo             |
| DELETE | `/todos/{id}`    | Delete a todo             |

> You must include this header in all Todo requests:
```
Authorization: Bearer <your-token>
```

---

## Setup & Run

### 1. Clone the project

```bash
git clone https://github.com/yusufborucu/go-projects.git
cd go-projects/crud-api-jwt-gorm
```

### 2. Create `.env` file

Create a `.env` file in the root:

```env
DB_USER=root
DB_PASS=yourpassword
DB_NAME=crud-api-jwt-gorm
DB_HOST=127.0.0.1
DB_PORT=3306

JWT_SECRET=supersecretkey
```

### 3. Install dependencies

```bash
go mod tidy
```

### 4. Run the app

```bash
go run main.go
```

Make sure your MySQL database is running. Tables will be auto-migrated on first run.

---

## Example Requests (Postman or curl)

### Register:
```http
POST /signup
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "123456"
}
```

### Login:
```http
POST /login
Content-Type: application/json

{
  "email": "john@example.com",
  "password": "123456"
}
```

Response:
```json
{ "token": "eyJhbGciOi..." }
```

---

## Project Structure

```
crud-api-jwt-gorm/
│
├── main.go
├── .env
├── go.mod
│
├── config/         # DB connection
├── controllers/    # HTTP handlers
├── middleware/     # JWT middleware
├── models/         # GORM models
└── routes/         # Routing logic
```

---

## Tech Stack

- [Go](https://go.dev/)
- [GORM](https://gorm.io/)
- [MySQL](https://www.mysql.com/)
- [JWT](https://jwt.io/)
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)

---

## License

MIT