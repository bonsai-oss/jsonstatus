package jsonstatus

import (
	"encoding/json"
	"io"
	"net/http"
)

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (s Status) Encode(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(s.Code)
	return json.NewEncoder(w).Encode(s)
}

func Decode(r io.Reader) (*Status, error) {
	var s Status
	decodeError := json.NewDecoder(r).Decode(&s)
	return &s, decodeError
}
