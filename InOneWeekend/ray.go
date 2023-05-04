package main


type Ray struct {
	origin, direction Vec3
}

func (r Ray) at(t float64) Vec3 {
	b := r.direction.mulScalar(t)
	a := r.origin
	return a.add(b)
}


func (r Ray) Ray_color() Vec3 {
	sphere := Sphere{Vec3{0,0, -1}, 0.5}
	t := r.hit_sphere(sphere) 
	if t > 0.0 {
		N := unit_vector(r.at(t).sub(Vec3{0, 0, -1}))
		return N.addScalar(1).mulScalar(0.5)
	}
	unit_direction := unit_vector(r.direction)
	t = 0.5 * (unit_direction.y + 1.0)
	return Vec3{1.0, 1.0, 1.0}.mulScalar(1.0 - t).add(Vec3{0.5, 0.7, 1.0}.mulScalar(t))
}

