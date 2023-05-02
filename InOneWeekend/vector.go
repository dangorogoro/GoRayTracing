package main

import (
	"math"
)

type vec3 struct {
	x, y, z float64
}

func (v1 vec3) dot(v2 vec3) float64 {
	return v1.x * v2.x + v1.y * v2.y + v1.z * v2.z
}
func (v vec3) length() float64 {
	return math.Sqrt(v.x * v.x + v.y * v.y + v.z * v.z)
}
func (v vec3) normalize() vec3 {
	l := v.length()
	return vec3{v.x / l, v.y / l, v.z / l}
}

type point3 vec3
type color vec3