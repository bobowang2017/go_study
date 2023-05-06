package common

import "math/rand"

func RandomInt(min, max int) int {
	if min < 0 || max < 0 {
		return 0
	}
	if min > max {
		return 0
	}
	if min == max {
		return min
	}
	return rand.Int()%max + min
}
