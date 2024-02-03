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
func (v Vec3) Minus() Vec3 {
  return Vec3{-v.X, -v.Y, -v.Z}
}
func (v Vec3) length_squared() float64 {
  return v.X * v.X + v.Y * v.Y + v.Z * v.Z
}
func (v Vec3) Length() float64 {
  return math.Sqrt(v.length_squared())
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
    if p.length_squared() < 1 {
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

func refract(uv Vec3, n Vec3, etai_over_etat float64) Vec3 {
  var cos_theta = math.Min(uv.Minus().Dot(n), 1.0)
  var r_out_perp = uv.Add(n.MulScalar(cos_theta)).MulScalar(etai_over_etat)
  var r_out_parallel = n.Minus().MulScalar(math.Sqrt(math.Abs(1.0 - r_out_perp.length_squared())))
  return r_out_parallel.Add(r_out_perp)
}