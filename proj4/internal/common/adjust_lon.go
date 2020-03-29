package common

import (
	"github.com/nodejayes/gology/proj4/internal/constants"
	"math"
)

func AdjustLon(x float64) float64 {
	if math.Abs(x) <= constants.SPI {
		return x
	}
	return x - (float64(Sign(x)) * constants.TWO_PI)
}
