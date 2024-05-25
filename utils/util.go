package utils

import (
	"math"
)

// Haversine formula to calculate distance between two points in km
func CalculateDistance(loc1, loc2 [2]float64) float64 {
	const R = 6371 // Earth radius in km
	dLat := (loc2[0] - loc1[0]) * (math.Pi / 180.0)
	dLon := (loc2[1] - loc1[1]) * (math.Pi / 180.0)
	lat1 := loc1[0] * (math.Pi / 180.0)
	lat2 := loc2[0] * (math.Pi / 180.0)

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Sin(dLon/2)*math.Sin(dLon/2)*math.Cos(lat1)*math.Cos(lat2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}
