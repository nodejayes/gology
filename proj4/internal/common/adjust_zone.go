package common

import "math"

func AdjustZone(zone, lon float64) float64 {
	if zone == math.Inf(-1) {
		zone = math.Floor((AdjustLon(lon)+math.Pi)*30/math.Pi) + 1

		if zone < 0 {
			return 0
		} else if zone > 60 {
			return 60
		}
	}
	return zone
}
