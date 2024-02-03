package primitives

type World struct {
  Elements []Hittable
}

func (w *World) Add(h Hittable) {
	w.Elements = append(w.Elements, h)
}

func (w *World) Hit(r *Ray, tMin float64, tMax float64, record *HitRecord) (hitAnything bool) {
  closest := tMax

  for _, element := range w.Elements {
    hit := element.Hit(r, tMin, closest, record)
    if hit {
      hitAnything = true
      closest = record.T
      record = record
    }
  }
  return
}