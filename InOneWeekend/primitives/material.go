package primitives

type Material interface {
	Scatter(r_in Ray, rec HitRecord) (bool, Ray, Vec3)
	Color() Vec3
}

type Lambertian struct {
	Albedo Vec3
}
type Metal struct {
	Albedo Vec3
}
func (l Lambertian) Scatter(r_in Ray, rec HitRecord) (bool, Ray, Vec3) {
	var scatter_direction = rec.Normal.Add(random_unit_vector())
	if scatter_direction.NearZero() {
		scatter_direction = rec.Normal
	}
	var scattered = Ray{rec.P, scatter_direction}
	var attenuation = l.Albedo
	return true, scattered, attenuation
}

func (l Lambertian) Color() Vec3 {
	return l.Albedo
}

func (l Metal) Scatter(r_in Ray, rec HitRecord) (bool, Ray, Vec3) {
	var reflected = reflect(r_in.Direction.unitVector(), rec.Normal);

	var scattered = Ray{rec.P, reflected}
	var attenuation = l.Albedo
	return true, scattered, attenuation
}

func (l Metal) Color() Vec3 {
	return l.Albedo
}