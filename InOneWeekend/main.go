package main

import (
	"fmt"
	"os"
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

	var origin = vec3{0, 0, 0}
	var horizontal = vec3{viewport_width, 0, 0}
	var vertical = vec3{0, viewport_height, 0}
	var lower_left_corner = origin.sub(horizontal.divScalar(2)).sub(vertical.divScalar(2)).sub(vec3{0, 0, focal_length})

	// Rendar

	fmt.Printf("P3\n%d %d\n255\n", image_width, uint32(image_height))
	for j := image_height - 1; j >= 0; j-- {
		fmt.Fprintln(os.Stderr, "\rScanlines remaining:", j)
		for i := 0; i < image_width; i++ {
			var u = float64(i) / (image_width - 1)
			var v = float64(j) / (image_height - 1)
			var r = ray{origin, lower_left_corner.add(horizontal.mulScalar(u)).add(vertical.mulScalar(v)).sub(origin)}
			pixel_color := r.ray_color()
			write_color(pixel_color)
		}
	}
	fmt.Fprintln(os.Stderr, "\nDone.\n")
}