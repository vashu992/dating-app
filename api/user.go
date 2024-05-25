package api

import (
	"github.com/vashu992/dating-app/store"
	"math/rand"
	"strconv"
	"time"

	"github.com/vashu992/dating-app/model"
)

func GenerateRandomUser() model.User {
	id := len(store.GetUsers()) + 1
	email := "user" + strconv.Itoa(id) + "@example.com"
	password := "password123"
	name := "User " + strconv.Itoa(id)
	gender := "Male"
	birthDate := time.Date(1990, time.Month(rand.Intn(12)+1), rand.Intn(28)+1, 0, 0, 0, 0, time.UTC)

	return model.User{
		ID:        id,
		Email:     email,
		Password:  password,
		Name:      name,
		Gender:    gender,
		BirthDate: birthDate,
	}
}
