package main

import "fmt"

func main() {
  var x [] float64
  x = append(x, 44)
  fmt.Println(x)
  
  y := make([] float64, 5, 10)
  arr := [] float64 { 1, 2, 3, 4, 5 }
  q := arr[0:5]
  fmt.Println(y, q)
  
  slice1 := [] int { 1, 2, 3 }
  slice2 := append(slice1, 4, 5)
  fmt.Println(slice1, slice2)
  
  slice11 := [] int { 1, 2, 3 }
  slice22 := make([]int, 2)
  copy(slice22, slice11)
  fmt.Println(slice11, slice22)
}
