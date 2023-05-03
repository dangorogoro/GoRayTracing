package main


type ray struct {
	origin, direction vec3
}

func (r ray) at(t float64) vec3 {
	b := r.direction.mulScalar(t)
	a := r.origin
	return a.add(b)
}


func (r ray) ray_color() vec3 {
	unit_direction := r.direction.normalize()
	t := 0.5 * (unit_direction.y + 1.0)
	return vec3{1.0, 1.0, 1.0}.mulScalar(1.0 - t).add(vec3{0.5, 0.7, 1.0}.mulScalar(t))
}

