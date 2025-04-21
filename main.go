package main

import (
    "fmt"
    "net/http"

    "github.com/Diony-source/go-person-api/handlers"
)

func main() {
    http.HandleFunc("/hello", handlers.HelloHandler)
    http.HandleFunc("/people", handlers.PeopleHandler)
    http.HandleFunc("/people/", handlers.PersonByIDHandler)
    http.HandleFunc("/people/stats", handlers.StatsHandler)

    fmt.Println("🚀 Server is running at http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("❌ Server error:", err)
    }
}
