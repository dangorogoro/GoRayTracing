package main

import (
  "fmt"
  "os"
  p "InOneWeekend/primitives"
)
const (
  ns = 100
)
func main(){
  fmt.Fprintln(os.Stderr, "Program Start.\n")
  // Image
  const aspect_ratio = 16.0 / 9.0
  const image_width = 400

  // Camera
  var sphere = p.Sphere{p.Vec3{0, 0, -1}, 0.5}
  var floor  = p.Sphere{p.Vec3{0, -100.5, -1}, 100}

  var world = p.World{[]p.Hittable{&sphere, &floor}}

  var camera = p.NewCamera(aspect_ratio, image_width)
  // Rendar
  camera.Rendar(&world, ns)

}