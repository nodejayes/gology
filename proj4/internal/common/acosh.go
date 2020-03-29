package common

import "math"

func Acosh(x float64) float64 {
	return 2 * math.Log(math.Sqrt((x+1)/2)+math.Sqrt((x-1)/2))
}
