package primitives

import (
	"math"
  "math/rand"
)
type Material interface {
	Scatter(r_in Ray, rec HitRecord) (bool, Ray, Vec3)
}

type Lambertian struct {
	Albedo Vec3
}
type Metal struct {
	Albedo Vec3
	fuzz float64
}
type Dielectric struct {
	refraction_index float64
}

func (l Lambertian) Scatter(r_in Ray, rec HitRecord) (flag bool, scattered Ray, attenuation Vec3) {
	flag = true
	var scatter_direction = rec.Normal.Add(random_unit_vector())
	if scatter_direction.NearZero() {
		scatter_direction = rec.Normal
	}
	scattered = Ray{rec.P, scatter_direction}
	attenuation = l.Albedo
	return
}


/*
func (l Lambertian) Color() Vec3 {
	return l.Albedo
}
*/

func (l Metal) Scatter(r_in Ray, rec HitRecord) (flag bool, scattered Ray, attenuation Vec3) {
	flag = true
	var reflected = reflect(r_in.Direction.unitVector(), rec.Normal);

	scattered = Ray{rec.P, reflected.Add(random_unit_vector().MulScalar(l.fuzz))}
	attenuation = l.Albedo
	return
}

/*
func (l Metal) Color() Vec3 {
	return l.Albedo
}
*/

func NewMetal(v Vec3, f float64) *Metal{
	m := new(Metal)
	m.Albedo = v
	m.fuzz = 1.0
	if f < 1.0 {
		m.fuzz = f
	}
	return m
}

func (l Dielectric) reflectance(cosine float64, refraction_index float64) (r0 float64) {
	r0 = (1 - refraction_index) / (1 + refraction_index)
	r0 = r0 * r0
	return
}

func (l Dielectric) Scatter(r_in Ray, rec HitRecord) (flag bool, scattered Ray, attenuation Vec3) {
	flag = true
	attenuation = Vec3{1.0, 1.0, 1.0}
	var ri = 0.0;

	if rec.front_face {
		ri = 1.0 / l.refraction_index
	}else{
		ri = l.refraction_index
	}


	var unit_direction = r_in.Direction.unitVector()
	var cos_theta = math.Min(unit_direction.Minus().Dot(rec.Normal), 1.0)
	var sin_theta = math.Sqrt(1.0 - cos_theta * cos_theta)

	var cannot_refract = false
	if ri * sin_theta > 1.0 {
		cannot_refract = true
	}

	var direction = Vec3{} 

	if cannot_refract || l.reflectance(cos_theta, ri) > rand.Float64(){
		direction = reflect(unit_direction, rec.Normal)
	}else {
		direction = refract(unit_direction, rec.Normal, ri);
	}

	scattered = Ray{rec.P, direction}
	return
}