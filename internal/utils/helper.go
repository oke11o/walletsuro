package utils

import "math"

func FloatWithFraction(in float64, fraction int) int64 {
	return int64(math.Round(in * math.Pow10(fraction)))
}
