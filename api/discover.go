package api

import (
	"github.com/vashu992/dating-app/store"
	"net/http"
	"sort"

	"github.com/vashu992/dating-app/model"
)

// GetCurrentUser is a placeholder to get the currently authenticated user
func GetCurrentUser(r *http.Request) model.User {
	// For simplicity, return the first user
	return store.GetUsers()[0]
}

// GetPotentialMatches fetches all users except the current user
func GetPotentialMatches() []model.User {
	users := store.GetUsers()
	currentUser := GetCurrentUser(nil) // Assume a way to get the current user

	var potentialMatches []model.User
	for _, user := range users {
		if user.ID != currentUser.ID {
			potentialMatches = append(potentialMatches, user)
		}
	}

	// Optionally sort by attractiveness
	sort.Slice(potentialMatches, func(i, j int) bool {
		return calculateAttractiveness(potentialMatches[i]) > calculateAttractiveness(potentialMatches[j])
	})

	return potentialMatches
}

// Calculate attractiveness based on swipe statistics (placeholder implementation)
func calculateAttractiveness(user model.User) float64 {
	// Placeholder: number of positive swipes - negative swipes
	return float64(user.ID % 10) // Just a dummy calculation
}
