package primitives

import (
  "fmt"
  "math"
)

func Write_color(c Vec3, samples_per_pixel int32) {
  var avg_color = c.DivScalar(float64(samples_per_pixel))
  // Apply the linear to gamma transform.
  var gamma_filtered_r = math.Sqrt(avg_color.X)
  var gamma_filtered_g = math.Sqrt(avg_color.Y)
  var gamma_filtered_b = math.Sqrt(avg_color.Z)
  fmt.Printf("%d %d %d\n", uint8(gamma_filtered_r * 255.999), uint8(gamma_filtered_g * 255.999), uint8(gamma_filtered_b * 255.999))
}