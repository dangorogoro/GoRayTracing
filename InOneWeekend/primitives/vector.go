package primitives

import (
	"math"
)

type Vec3 struct {
	X, Y, Z float64
}

func (v1 Vec3) Dot(v2 Vec3) float64 {
	return v1.X * v2.X + v1.Y * v2.Y + v1.Z * v2.Z
}
func (v Vec3) length() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y + v.Z * v.Z)
}
func (v Vec3) normalize() Vec3 {
	l := v.length()
	return Vec3{v.X / l, v.Y / l, v.Z / l}
}

func (v Vec3) Add(o Vec3) Vec3 {
	return Vec3{v.X + o.X, v.Y + o.Y, v.Z + o.Z}
}
func (v Vec3) Sub(o Vec3) Vec3 {
	return Vec3{v.X - o.X, v.Y - o.Y, v.Z - o.Z}
}
func (v Vec3) AddScalar(t float64) Vec3 {
	return Vec3{v.X + t, v.Y + t, v.Z + t}
}
func (v Vec3) SubScalar(t float64) Vec3 {
	return Vec3{v.X - t, v.Y - t, v.Z - t}
}
func (v Vec3) MulScalar(t float64) Vec3 {
	return Vec3{v.X * t, v.Y * t, v.Z * t}
}
func (v Vec3) DivScalar(t float64) Vec3 {
	return Vec3{v.X / t, v.Y / t, v.Z / t}
}
func unit_vector(v Vec3) Vec3 {
	return v.DivScalar(v.length())
}
type point3 Vec3
type color Vec3