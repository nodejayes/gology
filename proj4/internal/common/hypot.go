package common

import "math"

func HyPot(x, y float64) float64 {
	x = math.Abs(x)
	y = math.Abs(y)
	a := math.Max(x, y)
	tmp := 1.0
	if a > 0 {
		tmp = a
	}
	var b = math.Min(x, y) / tmp

	return a * math.Sqrt(1+math.Pow(b, 2))
}
