package main

import (
	"fmt"
	"net/http"

	"github.com/Diony-source/go-person-api/handlers"
)

func main() {
	http.HandleFunc("/hello", handlers.HelloHandler)

	http.HandleFunc("/people", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			query := r.URL.Query().Get("query")
			if query != "" {
				handlers.SearchPeopleHandler(w, r)
			} else {
				handlers.GetPeopleHandler(w, r)
			}
		case http.MethodPost:
			handlers.CreatePersonHandler(w, r)
		case http.MethodDelete:
			handlers.DeleteAllPeopleHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/people/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetPersonByIDHandler(w, r)
		case http.MethodPut:
			handlers.UpdatePersonHandler(w, r)
		case http.MethodPatch:
			handlers.PatchPersonHandler(w, r)
		case http.MethodDelete:
			handlers.DeletePersonHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/people/stats", handlers.StatsHandler)

	fmt.Println("🚀 Server is running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("❌ Server error:", err)
	}
}
