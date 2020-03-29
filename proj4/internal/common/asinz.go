package common

import "math"

func AsinZ(x float64) float64 {
	tmp := -1.0
	if math.Abs(x) > 1 && x > 1 {
		tmp = 1
	}
	return math.Asin(tmp)
}
