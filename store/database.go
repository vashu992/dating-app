package store

import "github.com/vashu992/dating-app/model"

// In-memory database
var users []model.User

// CreateUser adds a new user to the database
func CreateUser(user model.User) {
	users = append(users, user)
}

// GetUsers returns all users in the database
func GetUsers() []model.User {
	return users
}
