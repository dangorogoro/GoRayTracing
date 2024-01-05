package main

import (
	"fmt"
	"os"
	p "InOneWeekend/primitives"
)
func main(){
	// Image
	const aspect_ratio = 16.0 / 9.0
	const image_width = 400
	const image_height = image_width / aspect_ratio

	// Camera
	var viewport_height = 2.0;
	var viewport_width = viewport_height * aspect_ratio
	var focal_length = 1.0

	var Origin = p.Vec3{0, 0, 0}
	var horizontal = p.Vec3{viewport_width, 0, 0}
	var vertical = p.Vec3{0, viewport_height, 0}
	var lower_left_corner = Origin.Sub(horizontal.DivScalar(2)).Sub(vertical.DivScalar(2)).Sub(p.Vec3{0, 0, focal_length})


	var sphere = p.Sphere{p.Vec3{0, 0, -1}, 0.5}
	var floor  = p.Sphere{p.Vec3{0, -100.5, -1}, 100}

	var world = p.World{[]p.Hittable{&sphere, &floor}}

	// Rendar

	fmt.Printf("P3\n%d %d\n255\n", image_width, uint32(image_height))
	for j := image_height - 1; j >= 0; j-- {
		fmt.Fprintln(os.Stderr, "\rScanlines remaining:", j)
		for i := 0; i < image_width; i++ {
			var u = float64(i) / (image_width - 1)
			var v = float64(j) / (image_height - 1)
			var r = p.Ray{Origin, lower_left_corner.Add(horizontal.MulScalar(u)).Add(vertical.MulScalar(v)).Sub(Origin)}
			pixel_color := p.Ray_color(&r, &world)
			p.Write_color(pixel_color)
		}
	}
	fmt.Fprintln(os.Stderr, "\nDone.\n")
}