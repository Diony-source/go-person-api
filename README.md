# 👤 Go Person API

A fully functional RESTful API built with Go (Golang).  
This project allows you to manage people records using standard CRUD operations, with in-memory storage and clean handler architecture.

---

## 🚀 Features

- Add, list, update, delete, search and patch people
- Get statistics: total people, average age
- Query people by name (case-insensitive search)
- Proper error handling using helper functions
- `jsonError`, `writeJSON`, `getIDFromPath` helpers for cleaner handlers
- Modular and readable design
- 🔬 Includes full test coverage using `net/http/httptest`

---

## 📦 Installation

> Make sure [Go](https://golang.org/dl/) is installed.

```bash
git clone https://github.com/YOUR_USERNAME/go-person-api.git
cd go-person-api
go mod tidy
go run main.go
```

Server will run on:

```
http://localhost:8080
```

---

## 🔧 API Endpoints

| Method | Endpoint             | Description                      |
|--------|----------------------|----------------------------------|
| GET    | `/hello`             | Returns welcome message          |
| GET    | `/people`            | List all people                  |
| GET    | `/people/{id}`       | Get person by ID                 |
| GET    | `/people?query=Ali`  | Search people by name            |
| GET    | `/people/stats`      | Get statistics                   |
| POST   | `/people`            | Add a new person                 |
| PUT    | `/people/{id}`       | Full update of a person          |
| PATCH  | `/people/{id}`       | Partial update of a person       |
| DELETE | `/people/{id}`       | Delete person by ID              |
| DELETE | `/people`            | Delete all people                |

---

## 📄 JSON Structure

```json
{
  "id": 1,
  "name": "Diony",
  "age": 24,
  "phone": "123456789"
}
```

---

## 🧪 Running Tests

This project includes full test coverage for all endpoints and helper functions.

```bash
go test -v
```

You should see output similar to:

```bash
=== RUN   TestHelloHandler
--- PASS: TestHelloHandler (0.00s)
=== RUN   TestCreateAndGetPeople
--- PASS: TestCreateAndGetPeople (0.00s)
...
PASS
```

---

## 🧠 Tech Stack

- Go (Golang)
- `net/http` standard library
- `httptest` for unit testing
- In-memory slice-based data store
- Designed to evolve into PostgreSQL or other DB

---

## 🧊 License

MIT — use freely, modify safely.

---

## ✨ Author

Made with ❤️ by [Diony](https://github.com/Diony-source)