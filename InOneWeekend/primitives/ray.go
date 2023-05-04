package primitives


type Ray struct {
	Origin, Direction Vec3
}

func (r Ray) at(t float64) Vec3 {
	b := r.Direction.MulScalar(t)
	a := r.Origin
	return a.Add(b)
}


func (r Ray) Ray_color() Vec3 {
	sphere := Sphere{Vec3{0,0, -1}, 0.5}
	t := r.hit_sphere(sphere) 
	if t > 0.0 {
		N := unit_vector(r.at(t).Sub(Vec3{0, 0, -1}))
		return N.AddScalar(1).MulScalar(0.5)
	}
	unit_Direction := unit_vector(r.Direction)
	t = 0.5 * (unit_Direction.Y + 1.0)
	return Vec3{1.0, 1.0, 1.0}.MulScalar(1.0 - t).Add(Vec3{0.5, 0.7, 1.0}.MulScalar(t))
}

