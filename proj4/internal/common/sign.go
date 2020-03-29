package common

func Sign(x float64) int {
	if x < 0 {
		return -1
	}
	return 1
}
