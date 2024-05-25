package model

import (
	"time"
)

type User struct {
	ID            int       `json:"id"`
	Email         string    `json:"email"`
	Password      string    `json:"password"`
	Name          string    `json:"name"`
	Gender        string    `json:"gender"`
	BirthDate     time.Time `json:"birth_date"`
	Location      string    `json:"location"` // New field
	DistanceFromMe float64  `json:"distance_from_me"` // New field
}

// Calculate age from birth date
func (u User) Age() int {
	now := time.Now()
	years := now.Year() - u.BirthDate.Year()
	if now.YearDay() < u.BirthDate.YearDay() {
		years--
	}
	return years
}
