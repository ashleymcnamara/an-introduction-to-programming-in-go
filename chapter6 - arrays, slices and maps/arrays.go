package main

import "fmt"

func main() {
  var x [5]int
  x[4] = 100
  fmt.Println(x, x[4])
  
  var y [5] float64
  y[0] = 98
  y[1] = 93
  y[2] = 77
  y[3] = 82
  y[4] = 83
  
  var total float64 = 0
  leny := len(y)
  
  for i := 0; i < leny; i++ {
    total += y[i]
  }
  
  fmt.Printf("%v - %T\n", leny, leny)
  fmt.Println(total / float64(leny))
  
  var t float64 = 0
  for _, value := range y {
    t += value
  }
  
  fmt.Println(t / float64(leny))
  
  w := [6] float64 { 98, 93, 77, 82, 83, 12.22 }
  fmt.Println(w)
  
  u := [5] float64 {
    98,
    93,
    77,
    82,
    83,
  }
  
  fmt.Println(u)
}
