package common

import "math"

func Log1Py(x float64) float64 {
	var y = 1 + x
	var z = y - 1

	if z == 0 {
		return x
	}
	return x * math.Log(y) / z
}
