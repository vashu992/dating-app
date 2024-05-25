package api

import (

)

var swipeRecords = map[int]map[int]bool{} // userId -> swipedUserId -> preference (true for YES, false for NO)

func RecordSwipe(userID, swipedUserID int, preference bool) (bool, error) {
	if _, ok := swipeRecords[userID]; !ok {
		swipeRecords[userID] = map[int]bool{}
	}
	swipeRecords[userID][swipedUserID] = preference

	// Check if there's a match
	if otherPreference, ok := swipeRecords[swipedUserID][userID]; ok && otherPreference && preference {
		return true, nil
	}
	return false, nil
}
