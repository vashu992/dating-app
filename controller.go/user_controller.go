package controller

import (
	"encoding/json"
	"net/http"

	"github.com/vashu992/dating-app/api"

	"github.com/vashu992/dating-app/store"
)

// CreateUser generates and stores a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	newUser := api.GenerateRandomUser()
	store.CreateUser(newUser)

	jsonResponse(w, http.StatusOK, map[string]interface{}{
		"result": newUser,
	})
}

// Utility function to send JSON response
func jsonResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
