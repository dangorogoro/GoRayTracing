package primitives

import (
	"math"
)
type Sphere struct {
	Center Vec3
	Radius float64
}

func (r Ray) hit_sphere(s Sphere) float64 {
	oc := r.Origin.Sub(s.center)  // A - C
	a := r.Direction.Dot(r.Direction) // (tb)^2
	b := 2.0 * oc.Dot(r.Direction)  // 2t(A-c)
	c := oc.Dot(oc) - s.radius*s.radius // (A-C)^2 - r^2
	ans := b * b - 4 * a * c
	if ans < 0 {
		return -1.0
	}else {
		return (-b - math.Sqrt(ans)) / (2.0 * a)
	}
}