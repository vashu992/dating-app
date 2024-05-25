package controller

import (
	"encoding/json"
	"net/http"

	"github.com/vashu992/dating-app/api"
)

// Login authenticates a user and returns a token
func Login(w http.ResponseWriter, r *http.Request) {
	var credentials map[string]string
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		jsonResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}

	token, err := api.AuthenticateUser(credentials["email"], credentials["password"])
	if err != nil {
		jsonResponse(w, http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
		return
	}

	jsonResponse(w, http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
