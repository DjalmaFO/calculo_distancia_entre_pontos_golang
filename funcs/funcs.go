package funcs

import (
	"math"
)

func Sng(f float64) float64 {
	if f < 0 {
		return -1
	}

	return 1
}

func Atan2(y, x float64) (atan2 float64) {
	if x > 0 {
		return math.Atan(y / x)
	}

	if x < 0 {
		return math.Atan(y/x) + math.Pi
	}

	return math.Pi / 2 * Sng(y)
}
