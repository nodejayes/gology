package common

import (
	"github.com/nodejayes/gology/proj4/internal/constants"
	"math"
)

func AdjustLat(x float64) float64 {
	if math.Abs(x) < constants.HALF_PI {
		return x
	}
	return x - (float64(Sign(x)) * math.Pi)
}
