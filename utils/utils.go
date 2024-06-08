package utils

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var Validate = validator.New()

func ParseJSON(r *http.Request, payload interface{}) error {
	if r.Body == nil {
		return errors.New("please send a request body")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	_ = WriteJSON(w, status, map[string]string{"error": err.Error()})
}
