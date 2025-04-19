package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/hello", nil)
	rec := httptest.NewRecorder()

	helloHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "Hello") {
		t.Error("Expected greeting message")
	}
}

func TestCreateAndGetPeople(t *testing.T) {
	people = []Person{}
	nextId = 1

	person := Person{Name: "Diony", Age: 24, Phone: "12345"}
	body, _ := json.Marshal(person)

	req := httptest.NewRequest("POST", "/people", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	createPersonHandler(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", rec.Code)
	}

	// Check if added
	req = httptest.NewRequest("GET", "/people", nil)
	rec = httptest.NewRecorder()

	getPeopleHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", rec.Code)
	}

	var data []Person
	_ = json.Unmarshal(rec.Body.Bytes(), &data)
	if len(data) != 1 {
		t.Errorf("Expected 1 person, got %d", len(data))
	}
}

func TestGetPersonByIDHandler(t *testing.T) {
	people = []Person{{Id: 1, Name: "Test", Age: 20, Phone: "1111"}}
	req := httptest.NewRequest("GET", "/people/1", nil)
	rec := httptest.NewRecorder()

	getPersonByIDHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "Test") {
		t.Error("Expected person with name Test")
	}
}

func TestPatchPersonHandler(t *testing.T) {
	people = []Person{{Id: 1, Name: "Old", Age: 20, Phone: "1111"}}
	payload := Person{Phone: "9999"}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest("PATCH", "/people/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	patchPersonHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", rec.Code)
	}

	if !strings.Contains(rec.Body.String(), "9999") {
		t.Error("Phone not updated")
	}
}

func TestDeletePersonHandler(t *testing.T) {
	people = []Person{{Id: 1, Name: "DeleteMe", Age: 20, Phone: "0000"}}
	req := httptest.NewRequest("DELETE", "/people/1", nil)
	rec := httptest.NewRecorder()

	deletePersonHandler(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Errorf("Expected 204, got %d", rec.Code)
	}

	if len(people) != 0 {
		t.Error("Person not deleted")
	}
}

func TestSearchPeopleHandler(t *testing.T) {
	people = []Person{
		{Name: "Diony"},
		{Name: "Alice"},
	}

	req := httptest.NewRequest("GET", "/people?query=diony", nil)
	rec := httptest.NewRecorder()

	searchPeopleHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "Diony") {
		t.Error("Expected to find Diony")
	}
}

func TestStatsHandler(t *testing.T) {
	people = []Person{
		{Name: "A", Age: 20},
		{Name: "B", Age: 30},
	}
	req := httptest.NewRequest("GET", "/people/stats", nil)
	rec := httptest.NewRecorder()

	statsHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), `"average_age":25`) {
		t.Error("Average age incorrect")
	}
}
