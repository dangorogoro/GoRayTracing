package primitives

import (
  "math"
)
type Sphere struct {
  Center Vec3
  Radius float64
  Mat Material
}


func (s *Sphere) Hit(r *Ray, tMin float64, tMax float64) (flag bool, rec HitRecord) {
  oc := r.Origin.Sub(s.Center)  // A - C
  a := r.Direction.length_squared() // (tb)^2
  b := 2.0 * oc.Dot(r.Direction)  // 2t(A-c)
  c := oc.length_squared() - s.Radius*s.Radius // (A-C)^2 - r^2
  ans := b * b - 4 * a * c
  if ans > 0 {
    t := (-b - math.Sqrt(ans)) / (2.0 * a)
    if t >= tMax || t <= tMin {
      t = (-b + math.Sqrt(ans)) / (2.0 * a)
      if t >= tMax || t <= tMin {
        flag = false
        return
      }
    }
    rec.T = t
    rec.P = r.At(t)
    rec.Mat = s.Mat
    //rec.Normal = (rec.P.Sub(s.Center)).DivScalar(s.Radius)

    var outward_normal = (rec.P.Sub(s.Center)).DivScalar(s.Radius)
    rec.set_face_normal(r, &outward_normal)
    flag = true
    return
  }
  flag = false
  return
}
