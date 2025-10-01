package utils

import "math"

// CalculateTotalPages returns the number of pages based on total items and page size
func CalculateTotalPages(totalItems, pageSize int) int {
	if pageSize <= 0 {
		return 0
	}
	return int(math.Ceil(float64(totalItems) / float64(pageSize)))
}
