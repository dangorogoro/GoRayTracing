package primitives

import (
  "math"
)

type Ray struct {
  Origin, Direction Vec3
}

func (r Ray) At(t float64) Vec3 {
  b := r.Direction.MulScalar(t)
  a := r.Origin
  return a.Add(b)
}


func Ray_color(r *Ray, depth int32, world Hittable) Vec3 {
  if depth <= 0 {
    return Vec3{0, 0, 0}
  }
  hit, rec := world.Hit(r, 0.001, math.MaxFloat64) 
  if hit {
    flag, scattered, attenuation := rec.Mat.Scatter(*r, rec)
    if flag {
      return Ray_color(&scattered, depth-1, world).Mul(attenuation)
    }
    return Vec3{0, 0, 0}
  }
  unit_direction := r.Direction.unitVector()
  return gradient(unit_direction)
}


func gradient(v Vec3) Vec3 {
  // scale t to be between 0.0 and 1.0
  t := 0.5 * (v.Y + 1.0)
  white := Vec3{1.0, 1.0, 1.0}
  blue  := Vec3{0.5, 0.7, 1.0}
  // linear blend: blended_value = (1 - t) * white + t * blue
  return white.MulScalar(1.0 - t).Add(blue.MulScalar(t))
}