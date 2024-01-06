package primitives

import (
  "fmt"
)

func Write_color(c Vec3) {
  fmt.Printf("%d %d %d\n", uint8(c.X * 255.999), uint8(c.Y * 255.999), uint8(c.Z * 255.999))
}