# 👤 Go Person API

A simple RESTful API built with Go to manage people in memory with full CRUD operations and basic test coverage.

---

## 🌐 API Endpoints

| Method | URL              | Description               |
|--------|------------------|---------------------------|
| GET    | `/hello`         | Test server is alive      |
| GET    | `/people`        | List all people           |
| POST   | `/people`        | Add a new person          |
| GET    | `/people/{id}`   | Retrieve person by ID     |
| PUT    | `/people/{id}`   | Update person by ID       |
| DELETE | `/people/{id}`   | Delete person by ID       |

---

## 🚀 Run Locally

### Requirements:
- Go installed (v1.20+)
- Postman or curl for API testing

```bash
git clone https://github.com/YOUR_USERNAME/go-person-api.git
cd go-person-api
go run main.go
