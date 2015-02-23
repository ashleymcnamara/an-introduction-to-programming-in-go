package main

import ("fmt"; "math")

// Distance between two points P1(x1, y1) and P2(x2, y2)
func distance(x1, y1, x2, y2 float64) float64 {
  a := x2 - x1
  b := y2 - y1
  return math.Sqrt(a*a + b*b)
}

func rectangleArea(x1, y1, x2, y2 float64) float64 {
  l := distance(x1, y1, x1, y2)
  w := distance(x1, y1, x2, y1)
  return l * w
}

func circleArea(x, y, r float64) float64 {
  return math.Pi * math.Pow(r, 2)
}

type Circle struct {
  x, y, r float64 
}

func circleAreaC(c Circle) float64 {
  return math.Pi * math.Pow(c.r, 2)
}

func circleAreaCP(c *Circle) float64 {
  return math.Pi * math.Pow(c.r, 2)
}

func (c *Circle) area() float64 {
  return math.Pi * math.Pow(c.r, 2)
}

type Rectangle struct {
  x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
  l := distance(r.x1, r.y1, r.x1, r.y2)
  w := distance(r.x1, r.y1, r.x2, r.y1)
  return l * w
}

type Person struct {
  Name string
}   

func (p *Person) Talk() {
  fmt.Println("Hi, my name is", p.Name)
}

// has-a relationship
type Android1 struct {
  Person Person
  Model string
}


// is-a relationship
type Android2 struct {
  Person
  Model string
}

type Shape interface {
  area() float64
}

func totalArea(shapes ...Shape) float64 {
  var area float64
  for _, s := range shapes {
    area += s.area()
  }
  return area
}

type MultiShape struct {
  shapes [] Shape
}

func (m *MultiShape) area() float64 {
  var areas float64
  
  for _, shape := range m.shapes {
    areas += shape.area()
  }
  
  return areas
}

func main() {
  var rx1, ry1 float64 = 0, 0
  var rx2, ry2 float64 = 10, 10
  var cx, cy, cr float64 = 0, 0, 5
  
  fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
  fmt.Println(circleArea(cx, cy, cr))
  
  var c1 Circle
  c1.x = 10
  c1.y = 5
  c1.r = 10
  fmt.Println(c1.x, c1.y, c1.r)
  
  c2 := new(Circle)  
  c2.x = 5
  c2.y = 1
  c2.r = 44
  //c3 := Circle{x: 0, y: 0, r:5}
  c4 := Circle{0, 0, 5}
  fmt.Println(circleAreaC(c4), circleAreaCP(c2), c4.area())
  
  r := Rectangle{0, 0, 10, 10}
  fmt.Println(r.area())
  
  a := new(Android2)
  a.Name = "Robot!"
  a.Talk()
  
  andr := Android2{Person: Person{"me"}, Model: "android"}
  andr.Talk()
  
  ms := MultiShape{shapes: []Shape{&c1, &r, c2}}
  fmt.Println(totalArea(&c1, &r), ms.area())
}
