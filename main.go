package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type person struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Phone string `json:"phone"`
}

var people []person
var nextId = 1

func findPersonByID(id int) (*person, int) {
	for i, p := range people {
		if p.Id == id {
			return &people[i], i
		}
	}
	return nil, -1
}

func jsonError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{
		"error": message,
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Hello, Diony! Welcome to your Go API 🎉"}
	json.NewEncoder(w).Encode(response)
}

func getPeopleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

func createPersonHandler(w http.ResponseWriter, r *http.Request) {
	var newPerson person
	if err := json.NewDecoder(r.Body).Decode(&newPerson); err != nil {
		jsonError(w, "Invalid input", http.StatusBadRequest)
		return
	}
	newPerson.Id = nextId
	nextId++
	people = append(people, newPerson)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPerson)
}

func getPersonByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/people/"):]
	var id int
	_, err := fmt.Sscanf(idStr, "%d", &id)
	if err != nil {
		jsonError(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	p, _ := findPersonByID(id)
	if p == nil {
		jsonError(w, "Person not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func updatePersonHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/people/"):]
	var id int
	_, err := fmt.Sscanf(idStr, "%d", &id)
	if err != nil {
		jsonError(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	p, _ := findPersonByID(id)
	if p == nil {
		jsonError(w, "Person not found", http.StatusNotFound)
		return
	}

	var updated person
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		jsonError(w, "Invalid input", http.StatusBadRequest)
		return
	}

	p.Name = updated.Name
	p.Age = updated.Age
	p.Phone = updated.Phone

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func deletePersonHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/people/"):]
	var id int
	_, err := fmt.Sscanf(idStr, "%d", &id)
	if err != nil {
		jsonError(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	_, index := findPersonByID(id)
	if index == -1 {
		jsonError(w, "Person not found", http.StatusNotFound)
		return
	}

	people = append(people[:index], people[index+1:]...)
	w.WriteHeader(http.StatusNoContent)
}

func searchPeopleHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		jsonError(w, "Query parameter is required", http.StatusBadRequest)
		return
	}

	var results []person
	for _, p := range people {
		if strings.Contains(strings.ToLower(p.Name), strings.ToLower(query)) {
			results = append(results, p)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func deleteAllPeopleHandler(w http.ResponseWriter, r *http.Request) {
	people = []person{}
	w.WriteHeader(http.StatusNoContent)
}

func patchPersonHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/people/"):]
	var id int
	_, err := fmt.Sscanf(idStr, "%d", &id)
	if err != nil {
		jsonError(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	p, _ := findPersonByID(id)
	if p == nil {
		jsonError(w, "Person not found", http.StatusNotFound)
		return
	}

	var updated person
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		jsonError(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if updated.Name != "" {
		p.Name = updated.Name
	}
	if updated.Age != 0 {
		p.Age = updated.Age
	}
	if updated.Phone != "" {
		p.Phone = updated.Phone
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	total := len(people)
	sum := 0
	for _, p := range people {
		sum += p.Age
	}

	avg := 0
	if total > 0 {
		avg = sum / total
	}

	stats := map[string]interface{}{
		"total":         total,
		"average_age":   avg,
		"person_sample": people,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

func main() {
	http.HandleFunc("/hello", helloHandler)

	http.HandleFunc("/people", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			query := r.URL.Query().Get("query")
			if query != "" {
				searchPeopleHandler(w, r)
			} else {
				getPeopleHandler(w, r)
			}
		case http.MethodPost:
			createPersonHandler(w, r)
		case http.MethodDelete:
			deleteAllPeopleHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/people/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getPersonByIDHandler(w, r)
		case http.MethodPut:
			updatePersonHandler(w, r)
		case http.MethodPatch:
			patchPersonHandler(w, r)
		case http.MethodDelete:
			deletePersonHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/people/stats", statsHandler)

	fmt.Println("🚀 Server is running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("❌ Server error:", err)
	}
}
