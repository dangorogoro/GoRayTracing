package main

import (
	"math"
)

type Vec3 struct {
	x, y, z float64
}

func (v1 Vec3) dot(v2 Vec3) float64 {
	return v1.x * v2.x + v1.y * v2.y + v1.z * v2.z
}
func (v Vec3) length() float64 {
	return math.Sqrt(v.x * v.x + v.y * v.y + v.z * v.z)
}
func (v Vec3) normalize() Vec3 {
	l := v.length()
	return Vec3{v.x / l, v.y / l, v.z / l}
}

func (v Vec3) add(o Vec3) Vec3 {
	return Vec3{v.x + o.x, v.y + o.y, v.z + o.z}
}
func (v Vec3) sub(o Vec3) Vec3 {
	return Vec3{v.x - o.x, v.y - o.y, v.z - o.z}
}
func (v Vec3) addScalar(t float64) Vec3 {
	return Vec3{v.x + t, v.y + t, v.z + t}
}
func (v Vec3) subScalar(t float64) Vec3 {
	return Vec3{v.x - t, v.y - t, v.z - t}
}
func (v Vec3) mulScalar(t float64) Vec3 {
	return Vec3{v.x * t, v.y * t, v.z * t}
}
func (v Vec3) divScalar(t float64) Vec3 {
	return Vec3{v.x / t, v.y / t, v.z / t}
}
type point3 Vec3
type color Vec3