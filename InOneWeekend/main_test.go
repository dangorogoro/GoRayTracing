package main

import (
	"fmt"
	"os"
  p "InOneWeekend/primitives"
	"testing"
)
func Benchmark_Ray(b *testing.B) {
  fmt.Fprintln(os.Stderr, "Program Start.")
  // Image
  const aspect_ratio = 16.0 / 9.0
  const image_width = 400

  // Camera
  var material_ground = p.Lambertian{p.Vec3{0.8, 0.3, 0.0}}
  var material_center = p.Lambertian{p.Vec3{0.7, 0.3, 0.3}}
  var material_left   = p.NewMetal(p.Vec3{0.8, 0.8, 0.8}, 0.3)
  var material_right  = p.NewMetal(p.Vec3{0.8, 0.6, 0.2}, 1.0)


  var world = p.World{}
  world.Add(&p.Sphere{p.Vec3{ 0.0, -100.5, -1.0}, 100.0, material_ground})
  world.Add(&p.Sphere{p.Vec3{ 0.0,    0.0, -1.0},   0.5, material_center})
  world.Add(&p.Sphere{p.Vec3{-1.0,    0.0, -1.0},   0.5, material_left})
  world.Add(&p.Sphere{p.Vec3{ 1.0,    0.0, -1.0},   0.5, material_right})

  var camera = p.NewCamera(aspect_ratio, image_width)
  // Rendar
  camera.Rendar(&world, ns)

}