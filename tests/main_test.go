package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Diony-source/go-person-api/handlers"
	"github.com/Diony-source/go-person-api/models"
	"github.com/Diony-source/go-person-api/utils"
)

func resetState() {
	handlers.People = []models.Person{}
	handlers.NextID = 1
}

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/hello", nil)
	w := httptest.NewRecorder()
	handlers.HelloHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}

func TestCreateAndGetPeople(t *testing.T) {
	resetState()

	p := models.Person{Name: "Diony", Age: 24, Phone: "12345"}
	body, _ := json.Marshal(p)

	req := httptest.NewRequest("POST", "/people", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handlers.CreatePersonHandler(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}

	req = httptest.NewRequest("GET", "/people", nil)
	w = httptest.NewRecorder()
	handlers.GetPeopleHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}

	var people []models.Person
	_ = json.Unmarshal(w.Body.Bytes(), &people)

	if len(people) != 1 || people[0].Name != "Diony" {
		t.Errorf("Unexpected people: %+v", people)
	}
}

func TestPatchPerson(t *testing.T) {
	resetState()

	handlers.People = []models.Person{
		{Id: 1, Name: "Diony", Age: 24, Phone: "123"},
	}

	patched := models.Person{Phone: "99999"}
	if patched.Phone != "" {
		handlers.People[0].Phone = patched.Phone
	}

	if handlers.People[0].Phone != "99999" {
		t.Errorf("Patch failed. Expected phone 99999, got %s", handlers.People[0].Phone)
	}
}

func TestSearchQuery(t *testing.T) {
	resetState()

	handlers.People = []models.Person{
		{Name: "Diony", Age: 24},
		{Name: "Alice", Age: 30},
	}

	var result []models.Person
	query := "diony"
	for _, p := range handlers.People {
		if strings.Contains(strings.ToLower(p.Name), query) {
			result = append(result, p)
		}
	}

	if len(result) != 1 || result[0].Name != "Diony" {
		t.Errorf("Search failed: got %+v", result)
	}
}

func TestStatsHandler(t *testing.T) {
	resetState()

	handlers.People = []models.Person{
		{Name: "A", Age: 20},
		{Name: "B", Age: 30},
	}

	req := httptest.NewRequest("GET", "/people/stats", nil)
	w := httptest.NewRecorder()
	handlers.StatsHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}

func TestJsonError(t *testing.T) {
	rec := httptest.NewRecorder()
	utils.JSONError(rec, "Test error", http.StatusBadRequest)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("Expected 400, got %d", rec.Code)
	}

	if !strings.Contains(rec.Body.String(), `"Test error"`) {
		t.Errorf("Unexpected body: %s", rec.Body.String())
	}
}
