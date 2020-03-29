package common

import "math"

func AsinHY(x float64) float64 {
	var y = math.Abs(x)
	y = Log1Py(y * (1 + y/(HyPot(1, y)+1)))
	if x < 0 {
		return -y
	}
	return y
}
