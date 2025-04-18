package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type person struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Phone string `json:"phone"`
}

var people []person
var nextId = 1

// === Yardımcı Fonksiyon ===
func findPersonByID(id int) (*person, int) {
	for i, p := range people {
		if p.Id == id {
			return &people[i], i
		}
	}
	return nil, -1
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
		http.Error(w, "Invalid input", http.StatusBadRequest)
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
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	p, _ := findPersonByID(id)
	if p == nil {
		http.Error(w, "Person not found", http.StatusNotFound)
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
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	p, _ := findPersonByID(id)
	if p == nil {
		http.Error(w, "Person not found", http.StatusNotFound)
		return
	}

	var updated person
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
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
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	_, index := findPersonByID(id)
	if index == -1 {
		http.Error(w, "Person not found", http.StatusNotFound)
		return
	}

	people = append(people[:index], people[index+1:]...)
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	http.HandleFunc("/hello", helloHandler)

	http.HandleFunc("/people", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getPeopleHandler(w, r)
			return
		}
		if r.Method == http.MethodPost {
			createPersonHandler(w, r)
			return
		}
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	http.HandleFunc("/people/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getPersonByIDHandler(w, r)
		case http.MethodPut:
			updatePersonHandler(w, r)
		case http.MethodDelete:
			deletePersonHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("🚀 Server is running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("❌ Server error:", err)
	}
}
