package primitives

import (
  "fmt"
  "os"
  "math/rand"
)
type Camera struct {
  lowerLeft, horizontal, vertical, origin, pixel00_loc, pixel_delta_u, pixel_delta_v, center Vec3
  image_width, image_height uint32
  aspect_ratio float64
  samples_per_pixel int32
}

func NewCamera(aspect_ratio float64, image_width uint32) *Camera {
  c := new(Camera)
  c.lowerLeft = Vec3{-2.0, -1.0, -1.0}
  c.horizontal = Vec3{4.0, 0.0, 0.0}
  c.vertical = Vec3{0.0, 2.0, 0.0}
  c.origin = Vec3{0.0, 0.0, 0.0}
  c.image_width = image_width
  c.aspect_ratio = aspect_ratio
  c.image_height = uint32(float64(image_width) / c.aspect_ratio)
  return c
}

func (c *Camera) rayAt(u float64, v float64) Ray {
  position := c.position(u, v)
  direction := c.getDirection(position)
  return Ray{c.origin, direction}
}
func (c *Camera) position(u float64, v float64) Vec3 {
  horizontal := c.horizontal.MulScalar(u)
  vertical := c.vertical.MulScalar(v)

  return horizontal.Add(vertical)
}

func (c *Camera) getDirection(position Vec3) Vec3 {
  return c.lowerLeft.Add(position)
}
func (c *Camera) Rendar(world Hittable, alias_sampling int32) {
  fmt.Fprintln(os.Stderr, "\rRendar start")
  c.initialize(alias_sampling)
  fmt.Printf("P3\n%d %d\n255\n", c.image_width, c.image_height)
  for j := int32(c.image_height - 1); j >= 0; j-- {
    fmt.Fprintln(os.Stderr, "\rScanlines remaining:", j)
    for i := int32(0); i < int32(c.image_width); i++ {
      var pixel_color = Vec3{0, 0, 0}
      for sample := int32(0); sample < c.samples_per_pixel; sample++ {
        var r = c.getRay(i, j)
        pixel_color = pixel_color.Add(Ray_color(&r, world))
      }
      Write_color(pixel_color, c.samples_per_pixel)
    }
  }
  fmt.Fprintln(os.Stderr, "\nDone.\n")

}
func (c *Camera) initialize(alias_sampling int32) {
  if c.image_height < 1 {
    c.image_height = 1
  }
  c.samples_per_pixel = alias_sampling
  var center = Vec3{0, 0, 0}
  // Determine viewport dimensions
  var focal_length = 1.0
  var viewport_height = 2.0
  var viewport_width = viewport_height * c.aspect_ratio
  
  // Calculate the vectors across the horizontal and down the vertical viewport edges.
  var viewport_u = Vec3{viewport_width, 0, 0}
  var viewport_v = Vec3{0, viewport_height, 0} // original is minus

  // Calculate the horizontal and vertical delta vectors from pixel to pixel.
  c.pixel_delta_u = viewport_u.DivScalar(float64(c.image_width))
  c.pixel_delta_v = viewport_v.DivScalar(float64(c.image_height))

  // Calculate the location of the upper left pixel.
  var viewport_upper_left = center.Sub(Vec3{0, 0, focal_length}).Sub(viewport_u.DivScalar(2)).Sub(viewport_v.DivScalar(2))
  c.pixel00_loc = viewport_upper_left.Add((c.pixel_delta_u.Add(c.pixel_delta_v)).MulScalar(2))

}

func (c *Camera) rayColor(r *Ray, world Hittable) Vec3 {
  return Ray_color(r, world)
}

func (c *Camera) getRay(i int32, j int32) Ray {
  var pixel_center = c.pixel00_loc.Add(c.pixel_delta_u.MulScalar(float64(i))).Add(c.pixel_delta_v.MulScalar(float64(j)))
  var pixel_sample = pixel_center.Add(c.pixelSampleSquare())
  var ray_direction = pixel_sample.Sub(c.center)
  return Ray{c.center, ray_direction}
}

func (c *Camera) pixelSampleSquare() Vec3 {
  var px = -0.5 + rand.Float64()
  var py = -0.5 + rand.Float64()
  return c.pixel_delta_u.MulScalar(px).Add(c.pixel_delta_v.MulScalar(py))
}