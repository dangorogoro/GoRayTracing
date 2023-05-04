package primitives

type hitRecord struct {
	T float64
	P, Normal Vec3
}

type Hittable interface {
	Hit(r *Ray, tMin float64, tMax float64) (bool, HitRecord)
}