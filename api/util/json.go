package util

import (
	"encoding/json"
	"net/http"
)

func JsonParsingError(w http.ResponseWriter) error {
	return UserInputError(w, "error parsing request body. please try again")
}

func UserInputError(w http.ResponseWriter, message string) error {
	w.WriteHeader(http.StatusBadRequest)
	err := json.NewEncoder(w).Encode(map[string]string{
		"message": message,
		"success": "false",
	})
	if err != nil {
		return err
	}
	return nil
}
