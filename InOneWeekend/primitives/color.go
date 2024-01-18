package primitives

import (
  "fmt"
)

func Write_color(c Vec3, samples_per_pixel int32) {
  var avg_color = c.DivScalar(float64(samples_per_pixel))
  fmt.Printf("%d %d %d\n", uint8(avg_color.X * 255.999), uint8(avg_color.Y * 255.999), uint8(avg_color.Z * 255.999))
}