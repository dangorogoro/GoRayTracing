package primitives

import (
  "fmt"
  "os"
  "math"
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
  var R = math.Cos(math.Pi/4)
  var material_left   = Lambertian{Vec3{0, 0, 1}}
  var material_right   = Lambertian{Vec3{1, 0, 0}}


  var world = World{}
  world.Add(&Sphere{Vec3{-R,    0.0, -1.0},   R, material_left})
  world.Add(&Sphere{Vec3{ R,    0.0, -1.0},   R, material_right})

  var camera = NewCamera(aspect_ratio, image_width)
  // Rendar
  camera.Rendar(&world, ns)

}