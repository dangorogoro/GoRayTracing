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

var ray_rec = HitRecord{}

func Ray_color(r *Ray, depth int32, world Hittable) (vec Vec3) {
  vec = Vec3{0, 0, 0}
  if depth <= 0 {
    return
  }
  var hit = world.Hit(r, 0.001, math.MaxFloat64, &ray_rec) 
  if hit {
    flag, scattered, attenuation := ray_rec.Mat.Scatter(*r, ray_rec)
    if flag {
      vec = Ray_color(&scattered, depth-1, world).Mul(attenuation)
      return
    }
    return
  }
  var unit_direction = r.Direction.unitVector()
  vec = gradient(unit_direction)
  return
}


func gradient(v Vec3) Vec3 {
  // scale t to be between 0.0 and 1.0
  t := 0.5 * (v.Y + 1.0)
  white := Vec3{1.0, 1.0, 1.0}
  blue  := Vec3{0.5, 0.7, 1.0}
  // linear blend: blended_value = (1 - t) * white + t * blue
  return white.MulScalar(1.0 - t).Add(blue.MulScalar(t))
}