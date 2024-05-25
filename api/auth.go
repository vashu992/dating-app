package api

import (
	"github.com/vashu992/dating-app/store"
	"errors"
)

func AuthenticateUser(email, password string) (string, error) {
	users := store.GetUsers()
	for _, user := range users {
		if user.Email == email && user.Password == password {
			// Return a dummy token for simplicity
			return "dummy_token", nil
		}
	}
	return "", errors.New("invalid credentials")
}
