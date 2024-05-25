package controller

import (
	//"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"

	"github.com/vashu992/dating-app/api"
	"github.com/vashu992/dating-app/model"
	"github.com/vashu992/dating-app/utils"
)

// Discover fetches profiles of potential matches
func Discover(w http.ResponseWriter, r *http.Request) {
	currentUser := api.GetCurrentUser(r) // Assume this function retrieves the currently authenticated user

	// Get query parameters for filtering
	ageMin, _ := strconv.Atoi(r.URL.Query().Get("ageMin"))
	ageMax, _ := strconv.Atoi(r.URL.Query().Get("ageMax"))
	gender := r.URL.Query().Get("gender")

	// Get all users from the database
	users := api.GetPotentialMatches()

	// Filter users by age and gender
	var filteredUsers []model.User
	for _, user := range users {
		if (ageMin == 0 || user.Age() >= ageMin) &&
			(ageMax == 0 || user.Age() <= ageMax) &&
			(gender == "" || user.Gender == gender) &&
			user.ID != currentUser.ID { // Exclude current user
			filteredUsers = append(filteredUsers, user)
		}
	}

	// Sort users by distance if current user location is available
	if currentUser.Location != "" {
		currentUserLoc := parseLocation(currentUser.Location)
		for i := range filteredUsers {
			filteredUsers[i].DistanceFromMe = utils.CalculateDistance(
				currentUserLoc,
				parseLocation(filteredUsers[i].Location),
			)
		}
		// Sort by distance
		sort.Slice(filteredUsers, func(i, j int) bool {
			return filteredUsers[i].DistanceFromMe < filteredUsers[j].DistanceFromMe
		})
	}

	jsonResponse(w, http.StatusOK, map[string]interface{}{
		"results": filteredUsers,
	})
}

// Utility function to parse location string "lat,lng" to [2]float64
func parseLocation(loc string) [2]float64 {
	var location [2]float64
	fmt.Sscanf(loc, "%f,%f", &location[0], &location[1])
	return location
}

// Utility function to send JSON response
// func jsonResponse(w http.ResponseWriter, status int, payload interface{}) {
// 	response, _ := json.Marshal(payload)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(status)
// 	w.Write(response)
// }