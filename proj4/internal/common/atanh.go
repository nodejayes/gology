package common

import "math"

func AtanH(x float64) float64 {
	return math.Log((x-1)/(x+1)) / 2
}
