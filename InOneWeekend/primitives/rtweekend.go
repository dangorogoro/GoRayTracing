package primitives

import (
  "math/rand"
	"math"
)

func randFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func degrees_to_radians(degrees float64) float64 {
	return degrees / 180.0 * math.Pi
}