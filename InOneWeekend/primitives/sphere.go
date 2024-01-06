package primitives

import (
  "math"
)
type Sphere struct {
  Center Vec3
  Radius float64
}

func (s *Sphere) Hit(r *Ray, tMin float64, tMax float64) (bool, HitRecord) {
  oc := r.Origin.Sub(s.Center)  // A - C
  a := r.Direction.Dot(r.Direction) // (tb)^2
  b := 2.0 * oc.Dot(r.Direction)  // 2t(A-c)
  c := oc.Dot(oc) - s.Radius*s.Radius // (A-C)^2 - r^2
  ans := b * b - 4 * a * c
  rec := HitRecord{}
  if ans > 0 {
    t := (-b - math.Sqrt(ans)) / (2.0 * a)
    if t < tMax && t > tMin {
      rec.T = t
      rec.P = r.At(t)
      rec.Normal = (rec.P.Sub(s.Center)).DivScalar(s.Radius)
      return true, rec
    }
    t = (-b + math.Sqrt(ans)) / (2.0 * a)
    if t < tMax && t > tMin {
      rec.T = t
      rec.P = r.At(t)
      rec.Normal = (rec.P.Sub(s.Center)).DivScalar(s.Radius)
      return true, rec
    }
  }
  return false, rec
}