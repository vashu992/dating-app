package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/vashu992/dating-app/api"
)

var swipeRecords = map[int]map[int]bool{} // userId -> swipedUserId -> preference (true for YES, false for NO)
var swipeStats = map[int]map[bool]int{}   // userId -> preference (true for YES, false for NO) -> count

// Swipe responds to a profile
func Swipe(w http.ResponseWriter, r *http.Request) {
	var requestData map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		jsonResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}

	currentUser := api.GetCurrentUser(r)
	userID := currentUser.ID
	swipedUserID, _ := strconv.Atoi(requestData["swipedUserID"].(string))
	preference := requestData["preference"].(string) == "YES"

	// Record the swipe
	if _, ok := swipeRecords[userID]; !ok {
		swipeRecords[userID] = map[int]bool{}
	}
	swipeRecords[userID][swipedUserID] = preference

	// Update swipe statistics
	if _, ok := swipeStats[swipedUserID]; !ok {
		swipeStats[swipedUserID] = map[bool]int{}
	}
	swipeStats[swipedUserID][preference]++

	// Check if there's a match
	matched := false
	if otherPreference, ok := swipeRecords[swipedUserID][userID]; ok && otherPreference && preference {
		matched = true
	}

	response := map[string]interface{}{
		"matched": matched,
	}
	if matched {
		response["matchID"] = swipedUserID
	}

	jsonResponse(w, http.StatusOK, map[string]interface{}{
		"results": response,
	})
}
