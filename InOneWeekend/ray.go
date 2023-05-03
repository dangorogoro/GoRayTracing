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
	if r.hit_sphere(sphere){
		return Vec3{1.0, 0.0, 0.0}
	}
	unit_direction := r.direction.normalize()
	t := 0.5 * (unit_direction.y + 1.0)
	return Vec3{1.0, 1.0, 1.0}.mulScalar(1.0 - t).add(Vec3{0.5, 0.7, 1.0}.mulScalar(t))
}

