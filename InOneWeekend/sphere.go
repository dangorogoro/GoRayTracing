package main

type Sphere struct {
	center Vec3
	radius float64
}

func (r Ray) hit_sphere(s Sphere) bool {
	oc := r.origin.sub(s.center)  // A - C
	a := r.direction.dot(r.direction) // (tb)^2
	b := 2.0 * oc.dot(r.direction)  // 2t(A-c)
	c := oc.dot(oc) - s.radius*s.radius // (A-C)^2 - r^2
	ans := b * b - 4 * a * c
	return ans > 0
}