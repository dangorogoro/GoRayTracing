package primitives

type World struct {
  Elements []Hittable
}

func (w *World) Add(h Hittable) {
	w.Elements = append(w.Elements, h)
}

func (w *World) Hit(r *Ray, tMin float64, tMax float64) (bool, HitRecord) {
  hitAnything := false
  closest := tMax
  record := HitRecord{}

  for _, element := range w.Elements {
    hit, tempRecord := element.Hit(r, tMin, closest)
    if hit {
      hitAnything = true
      closest = tempRecord.T
      record = tempRecord
    }
  }
  return hitAnything, record
}