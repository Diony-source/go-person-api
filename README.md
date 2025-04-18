# 👤 Go Person API

A simple RESTful API built in Go for managing contacts.  
Supports CRUD operations, search functionality, stats summary, and robust testing setup.

---

## 🚀 Features

- Add, update, delete and list people
- Get a person by ID
- Partial update with PATCH
- Delete all people
- Search people by name (case-insensitive)
- Summary stats: total people + average age
- JSON error response helper for cleaner error handling
- Full unit test coverage with mock storage

---

## 📦 Installation

Make sure you have Go installed on your machine.

```bash
git clone https://github.com/YOUR_USERNAME/go-person-api.git
cd go-person-api
go mod tidy
go run main.go
```

Visit the API at: `http://localhost:8080`

---

## 📚 API Endpoints

| Method | Endpoint         | Description                       |
|--------|------------------|-----------------------------------|
| GET    | /people          | Get all people                    |
| GET    | /people?query=x  | Search people by name             |
| POST   | /people          | Add new person                    |
| DELETE | /people          | Delete all people                 |
| GET    | /people/{id}     | Get person by ID                  |
| PUT    | /people/{id}     | Update person completely          |
| PATCH  | /people/{id}     | Update person partially           |
| DELETE | /people/{id}     | Delete person by ID               |
| GET    | /people/stats    | Show total count + avg age        |
| GET    | /hello           | Welcome message                   |

---

## 📊 Example Stats Response

```json
{
  "total": 2,
  "average_age": 26,
  "person_sample": [
    { "id": 1, "name": "Diony", "age": 24, "phone": "12345" },
    { "id": 2, "name": "Eren", "age": 28, "phone": "67890" }
  ]
}
```

---

## ✅ Tests & Coverage

This project includes a suite of unit tests using Go’s built-in `testing` package.  
It uses an in-memory store (`MemoryStore`) to mock storage and isolate tests.

### Run tests with verbose output:

```bash
go test -v
```

---

## 📄 License

This project is licensed under the MIT License.