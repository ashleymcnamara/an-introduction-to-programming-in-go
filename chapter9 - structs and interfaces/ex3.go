package main

import ("fmt"; "math")

type Shape interface {
  area() float64
  perimeter() float64
}

type Circle struct {
  x, y, r float64 
}

func (c *Circle) area() float64 {
  return math.Pi * math.Pow(c.r, 2)
}

func (c *Circle) perimeter() float64 {
  return 2 * math.Pi * c.r
}

type Rectangle struct {
  x1, y1, x2, y2 float64
}

// Distance between two points P1(x1, y1) and P2(x2, y2)
func distance(x1, y1, x2, y2 float64) float64 {
  a := x2 - x1
  b := y2 - y1
  return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) area() float64 {
  l := distance(r.x1, r.y1, r.x1, r.y2)
  w := distance(r.x1, r.y1, r.x2, r.y1)
  return l * w
}

func (r *Rectangle) perimeter() float64 {
  l := distance(r.x1, r.y1, r.x1, r.y2)
  w := distance(r.x1, r.y1, r.x2, r.y1)
  return 2 * (l + w)
}

func main() {
  c := Circle{5, 10, 50}
  fmt.Println(c.area(), c.perimeter())
  
  r := Rectangle{10, 5, 20, 45}
  fmt.Println(r.area(), r.perimeter())
}
