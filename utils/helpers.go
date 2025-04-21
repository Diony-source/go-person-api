package utils

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strings"
)

func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(data)
}

func JSONError(w http.ResponseWriter, message string, status int) {
    WriteJSON(w, status, map[string]string{"error": message})
}

func ParseJSONBody(r *http.Request, target interface{}) error {
    return json.NewDecoder(r.Body).Decode(target)
}

func GetIDFromPath(r *http.Request, prefix string) (int, error) {
    idStr := strings.TrimPrefix(r.URL.Path, prefix)
    var id int
    _, err := fmt.Sscanf(idStr, "%d", &id)
    return id, err
}
