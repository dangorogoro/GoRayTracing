package main

import "fmt"

func main(){
	// Image
	const image_width = 256
	const image_height = 256
	fmt.Printf("P3\n%d %d\n255\n", image_width, image_height)
	for j := image_height - 1; j >= 0; j-- {
		for i := 0; i < image_width; i++ {
			var r float64 = float64(i) / (image_width - 1)
			var g float64 = float64(j) / (image_height - 1)
			var b float64 = 0.25
			var ir uint8 = uint8(r * 255.999)
			var ig uint8 = uint8(g * 255.999)
			var ib uint8 = uint8(b * 255.999)
			fmt.Printf("%d %d %d\n", ir, ig, ib)
		}
	}
}