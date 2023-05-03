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

func (v vec3) add(o vec3) vec3 {
	return vec3{v.x + o.x, v.y + o.y, v.z + o.z}
}
func (v vec3) sub(o vec3) vec3 {
	return vec3{v.x - o.x, v.y - o.y, v.z - o.z}
}
func (v vec3) addScalar(t float64) vec3 {
	return vec3{v.x + t, v.y + t, v.z + t}
}
func (v vec3) subScalar(t float64) vec3 {
	return vec3{v.x - t, v.y - t, v.z - t}
}
func (v vec3) mulScalar(t float64) vec3 {
	return vec3{v.x * t, v.y * t, v.z * t}
}
func (v vec3) divScalar(t float64) vec3 {
	return vec3{v.x / t, v.y / t, v.z / t}
}
type point3 vec3
type color vec3