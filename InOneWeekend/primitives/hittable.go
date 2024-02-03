package primitives

type HitRecord struct {
  T float64
  P, Normal Vec3
  Mat Material
  front_face bool
}

type Hittable interface {
  Hit(r *Ray, tMin float64, tMax float64) (bool, HitRecord)
}

func (h *HitRecord) set_face_normal(r *Ray, outward_normal *Vec3){
  if outward_normal.Dot(r.Direction) < 0 {
    h.front_face = true
  }else {
    h.front_face = false
  }
  if h.front_face {
    h.Normal = *outward_normal
  }else {
    h.Normal = outward_normal.Minus()
  }
}