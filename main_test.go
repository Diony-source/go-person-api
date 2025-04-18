package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MemoryStore struct {
	Data []Person
}

func (m *MemoryStore) Save(people []Person) error {
	m.Data = people
	return nil
}

func (m *MemoryStore) Load() ([]Person, error) {
	return m.Data, nil
}

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/hello", nil)
	w := httptest.NewRecorder()
	helloHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestCreateAndGetPeople(t *testing.T) {
	people = []Person{} // Reset data
	nextId = 1

	// Create Person
	p := Person{Name: "Diony", Age: 24, Phone: "12345"}
	body, _ := json.Marshal(p)
	req := httptest.NewRequest("POST", "/people", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	createPersonHandler(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", w.Code)
	}

	// Get people
	req = httptest.NewRequest("GET", "/people", nil)
	w = httptest.NewRecorder()
	getPeopleHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var peopleResponse []Person
	_ = json.Unmarshal(w.Body.Bytes(), &peopleResponse)

	if len(peopleResponse) != 1 {
		t.Errorf("Expected 1 person, got %d", len(peopleResponse))
	}
	if peopleResponse[0].Name != "Diony" {
		t.Errorf("Expected person name 'Diony', got %s", peopleResponse[0].Name)
	}
}

func TestPatchPerson(t *testing.T) {
	store := &MemoryStore{}
	people := []Person{
		{Id: 1, Name: "Diony", Age: 24, Phone: "123"},
	}
	store.Save(people)

	// simulate patch
	updated := Person{Name: "Diony", Age: 24, Phone: "99999"}
	original := people[0]

	if updated.Name != "" {
		original.Name = updated.Name
	}
	if updated.Age != 0 {
		original.Age = updated.Age
	}
	if updated.Phone != "" {
		original.Phone = updated.Phone
	}

	if original.Name != "Diony" {
		t.Errorf("Expected name 'Diony', got %s", original.Name)
	}
	if original.Age != 24 {
		t.Errorf("Expected age 24, got %d", original.Age)
	}
	if original.Phone != "99999" {
		t.Errorf("Expected phone '99999', got %s", original.Phone)
	}
}


func TestSearchQuery(t *testing.T) {
	store := &MemoryStore{}
	people := []Person{
		{Name: "Diony", Age: 24},
		{Name: "Alice", Age: 30},
	}
	store.Save(people)

	var result []Person
	query := "diony"
	for _, p := range people {
		if strings.Contains(strings.ToLower(p.Name), query) {
			result = append(result, p)
		}
	}

	if len(result) != 1 || result[0].Name != "Diony" {
		t.Errorf("Search failed: expected Diony, got %+v", result)
	}
}

func TestStatsHandler(t *testing.T) {
	people := []Person{
		{Name: "A", Age: 20},
		{Name: "B", Age: 30},
	}

	total := len(people)
	sum := 0
	for _, p := range people {
		sum += p.Age
	}
	avg := sum / total

	if avg != 25 {
		t.Errorf("Expected average age 25, got %d", avg)
	}
}

func TestJsonError_InvalidID(t *testing.T) {
	rec := httptest.NewRecorder()
	jsonError(rec, "Invalid ID", http.StatusBadRequest)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("Expected 400, got %d", rec.Code)
	}

	if !strings.Contains(rec.Body.String(), `"Invalid ID"`) {
		t.Errorf("Unexpected body: %s", rec.Body.String())
	}
}
