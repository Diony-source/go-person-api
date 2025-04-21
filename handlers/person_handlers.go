package handlers

import (
    "net/http"
    "strings"

    "github.com/Diony-source/go-person-api/models"
    "github.com/Diony-source/go-person-api/utils"
)

var people []models.Person
var nextId = 1

func HelloHandler(w http.ResponseWriter, _ *http.Request) {
    utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "Hello, Diony! Welcome to your Go API 🎉"})
}

func PeopleHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        query := r.URL.Query().Get("query")
        if query != "" {
            handleSearch(w, query)
        } else {
            utils.WriteJSON(w, http.StatusOK, people)
        }
    case http.MethodPost:
        var newPerson models.Person
        if err := utils.ParseJSONBody(r, &newPerson); err != nil {
            utils.JSONError(w, "Invalid input", http.StatusBadRequest)
            return
        }
        newPerson.Id = nextId
        nextId++
        people = append(people, newPerson)
        utils.WriteJSON(w, http.StatusCreated, newPerson)
    case http.MethodDelete:
        people = []models.Person{}
        w.WriteHeader(http.StatusNoContent)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func PersonByIDHandler(w http.ResponseWriter, r *http.Request) {
    id, err := utils.GetIDFromPath(r, "/people/")
    if err != nil {
        utils.JSONError(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    p, index := findPersonByID(id)
    if p == nil {
        utils.JSONError(w, "Person not found", http.StatusNotFound)
        return
    }

    switch r.Method {
    case http.MethodGet:
        utils.WriteJSON(w, http.StatusOK, p)
    case http.MethodPut:
        var updated models.Person
        if err := utils.ParseJSONBody(r, &updated); err != nil {
            utils.JSONError(w, "Invalid input", http.StatusBadRequest)
            return
        }
        p.Name = updated.Name
        p.Age = updated.Age
        p.Phone = updated.Phone
        utils.WriteJSON(w, http.StatusOK, p)
    case http.MethodPatch:
        var updated models.Person
        if err := utils.ParseJSONBody(r, &updated); err != nil {
            utils.JSONError(w, "Invalid input", http.StatusBadRequest)
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
        utils.WriteJSON(w, http.StatusOK, p)
    case http.MethodDelete:
        people = append(people[:index], people[index+1:]...)
        w.WriteHeader(http.StatusNoContent)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func StatsHandler(w http.ResponseWriter, _ *http.Request) {
    total := len(people)
    sum := 0
    for _, p := range people {
        sum += p.Age
    }
    avg := 0
    if total > 0 {
        avg = sum / total
    }
    utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
        "total":         total,
        "average_age":   avg,
        "person_sample": people,
    })
}

func handleSearch(w http.ResponseWriter, query string) {
    var results []models.Person
    for _, p := range people {
        if strings.Contains(strings.ToLower(p.Name), strings.ToLower(query)) {
            results = append(results, p)
        }
    }
    utils.WriteJSON(w, http.StatusOK, results)
}

func findPersonByID(id int) (*models.Person, int) {
    for i, p := range people {
        if p.Id == id {
            return &people[i], i
        }
    }
    return nil, -1
}
