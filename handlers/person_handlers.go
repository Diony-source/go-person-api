package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Diony-source/go-person-api/models"
	"github.com/Diony-source/go-person-api/utils"
)

var People []models.Person
var NextID = 1

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "Hello, Diony! Welcome to your Go API 🎉"})
}

func GetPeopleHandler(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, People)
}

func CreatePersonHandler(w http.ResponseWriter, r *http.Request) {
	var newPerson models.Person
	if err := json.NewDecoder(r.Body).Decode(&newPerson); err != nil {
		utils.JSONError(w, "Invalid input", http.StatusBadRequest)
		return
	}
	newPerson.Id = NextID
	NextID++
	People = append(People, newPerson)

	utils.WriteJSON(w, http.StatusCreated, newPerson)
}

func FindPersonByID(id int) (*models.Person, int) {
	for i, p := range People {
		if p.Id == id {
			return &People[i], i
		}
	}
	return nil, -1
}

func GetPersonByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIDFromPath(r, "/people/")
	if err != nil {
		utils.JSONError(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	p, _ := FindPersonByID(id)
	if p == nil {
		utils.JSONError(w, "Person not found", http.StatusNotFound)
		return
	}

	utils.WriteJSON(w, http.StatusOK, p)
}

func UpdatePersonHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIDFromPath(r, "/people/")
	if err != nil {
		utils.JSONError(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	p, _ := FindPersonByID(id)
	if p == nil {
		utils.JSONError(w, "Person not found", http.StatusNotFound)
		return
	}

	var updated models.Person
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		utils.JSONError(w, "Invalid input", http.StatusBadRequest)
		return
	}

	p.Name = updated.Name
	p.Age = updated.Age
	p.Phone = updated.Phone

	utils.WriteJSON(w, http.StatusOK, p)
}

func PatchPersonHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIDFromPath(r, "/people/")
	if err != nil {
		utils.JSONError(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	p, _ := FindPersonByID(id)
	if p == nil {
		utils.JSONError(w, "Person not found", http.StatusNotFound)
		return
	}

	var updated models.Person
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
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
}

func DeletePersonHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIDFromPath(r, "/people/")
	if err != nil {
		utils.JSONError(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	_, index := FindPersonByID(id)
	if index == -1 {
		utils.JSONError(w, "Person not found", http.StatusNotFound)
		return
	}

	People = append(People[:index], People[index+1:]...)
	w.WriteHeader(http.StatusNoContent)
}

func DeleteAllPeopleHandler(w http.ResponseWriter, r *http.Request) {
	People = []models.Person{}
	w.WriteHeader(http.StatusNoContent)
}

func SearchPeopleHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		utils.JSONError(w, "Query parameter is required", http.StatusBadRequest)
		return
	}

	var results []models.Person
	for _, p := range People {
		if strings.Contains(strings.ToLower(p.Name), strings.ToLower(query)) {
			results = append(results, p)
		}
	}

	utils.WriteJSON(w, http.StatusOK, results)
}

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	total := len(People)
	sum := 0
	for _, p := range People {
		sum += p.Age
	}

	avg := 0
	if total > 0 {
		avg = sum / total
	}

	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"total":         total,
		"average_age":   avg,
		"person_sample": People,
	})
}
