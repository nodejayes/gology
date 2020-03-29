package common

import "math"

func AsinH(x float64) float64 {
	s := -1
	if x >= 0 {
		s = 1
	}
	return float64(s) * (math.Log(math.Abs(x) + math.Sqrt(x*x+1)))
}
