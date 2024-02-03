package primitives

import (
  "fmt"
  "os"
)
const (
  ns = 100
)
func App(){
  fmt.Fprintln(os.Stderr, "Program Start.\n")
  // Image
  const aspect_ratio = 16.0 / 9.0
  const image_width = 400

  // Camera
  var material_ground = Lambertian{Vec3{0.8, 0.8, 0.0}}
  //var material_center = Lambertian{Vec3{0.7, 0.3, 0.3}}
  //var material_left   = NewMetal(Vec3{0.8, 0.8, 0.8}, 0.3)
  var material_center = Dielectric{1.5}
  var material_left   = Dielectric{1.5}
  var material_right  = NewMetal(Vec3{0.8, 0.6, 0.2}, 1.0)


  var world = World{}
  world.Add(&Sphere{Vec3{ 0.0, -100.5, -1.0}, 100.0, material_ground})
  world.Add(&Sphere{Vec3{ 0.0,    0.0, -1.0},   0.5, material_center})
  world.Add(&Sphere{Vec3{-1.0,    0.0, -1.0},   0.5, material_left})
  world.Add(&Sphere{Vec3{ 1.0,    0.0, -1.0},   0.5, material_right})

  var camera = NewCamera(aspect_ratio, image_width)
  // Rendar
  camera.Rendar(&world, ns)

}