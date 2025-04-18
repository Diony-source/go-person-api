package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/hello", nil)
	w := httptest.NewRecorder()
	helloHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestCreateAndGetPeople(t *testing.T) {
	people = []person{} // Reset data
	nextId = 1

	// Create person
	p := person{Name: "Diony", Age: 24, Phone: "12345"}
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

	var peopleResponse []person
	_ = json.Unmarshal(w.Body.Bytes(), &peopleResponse)

	if len(peopleResponse) != 1 {
		t.Errorf("Expected 1 person, got %d", len(peopleResponse))
	}
	if peopleResponse[0].Name != "Diony" {
		t.Errorf("Expected person name 'Diony', got %s", peopleResponse[0].Name)
	}
}
