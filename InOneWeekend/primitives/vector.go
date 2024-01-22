package primitives

import (
  "math"
  "math/rand"
)

type Vec3 struct {
  X, Y, Z float64
}
type Point3 Vec3
type Color Vec3

func (v1 Vec3) Dot(v2 Vec3) float64 {
  return v1.X * v2.X + v1.Y * v2.Y + v1.Z * v2.Z
}
func (v Vec3) Length() float64 {
  return math.Sqrt(v.X * v.X + v.Y * v.Y + v.Z * v.Z)
}
func (v Vec3) Normalize() Vec3 {
  l := v.Length()
  return Vec3{v.X / l, v.Y / l, v.Z / l}
}

func (v Vec3) NearZero() bool {
  var s = 1e-8
  return (math.Abs(v.X) < s && math.Abs(v.Y) < s && math.Abs(v.Z) < s)
}

func (v Vec3) Add(o Vec3) Vec3 {
  return Vec3{v.X + o.X, v.Y + o.Y, v.Z + o.Z}
}
func (v Vec3) Sub(o Vec3) Vec3 {
  return Vec3{v.X - o.X, v.Y - o.Y, v.Z - o.Z}
}
func (v Vec3) Mul(o Vec3) Vec3 {
  return Vec3{v.X * o.X, v.Y * o.Y, v.Z * o.Z}
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
func (v *Vec3) unitVector() Vec3 {
  return v.DivScalar(v.Length())
}

func random_vec() Vec3 {
  return Vec3{rand.Float64(), rand.Float64(), rand.Float64()}
}

func random_vec_min_max(min, max float64) Vec3 {
  return Vec3{randFloat64(min, max), randFloat64(min, max), randFloat64(min, max)}
  
}

func random_in_unit_sphere() Vec3 {
  for {
    var p = random_vec_min_max(-1, 1)
    if p.Length() < 1 {
      return p
    }
  }
}

func random_unit_vector() Vec3 {
  var p = random_in_unit_sphere()
  return p.unitVector()
}

func random_on_hemisphere(normal Vec3) Vec3 {
  var p = random_unit_vector()
  if normal.Dot(p) > 0 {
    return p
  } else {
    return p.MulScalar(-1)
  }
}

func reflect(v Vec3, n Vec3) Vec3 {
  return v.Sub(n.MulScalar(2 * v.Dot(n)))
}