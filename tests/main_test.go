package tests

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/Diony-source/go-person-api/handlers"
)

func TestHelloHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/hello", nil)
    w := httptest.NewRecorder()
    handlers.HelloHandler(w, req)

    if w.Code != http.StatusOK {
        t.Errorf("Expected status 200, got %d", w.Code)
    }
}
