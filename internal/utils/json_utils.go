package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func EncodeJSON[T any](w http.ResponseWriter, r *http.Request, status int, data T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		return fmt.Errorf("failed to encode json: %w", err)
	}

	return nil
}
