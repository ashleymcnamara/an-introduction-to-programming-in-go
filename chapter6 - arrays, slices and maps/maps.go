package main

import "fmt"

func main() {
  var x map[string]int = make(map[string]int)
  x["key"] = 10
  fmt.Println(x["key"])
  
  y := make(map[int]int)
  y[1] = 10
  fmt.Println(y[1])
  
  delete(y, 1)
  fmt.Println(y)
}
