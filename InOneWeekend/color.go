package main

import (
	"fmt"
)

func write_color(c vec3) {
	fmt.Printf("%d %d %d\n", uint8(c.x * 255.999), uint8(c.y * 255.999), uint8(c.z * 255.999))
}