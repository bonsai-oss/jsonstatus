package jsonstatus

import (
	"encoding/json"
	"net/http"
)

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (s Status) Encode(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(s)
}
