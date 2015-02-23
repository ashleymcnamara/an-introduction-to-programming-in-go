package main

/*
 sum is a function which takes a slice of numbers
and adds them together. What would its function signature look like in Go?
*/

import (
  "fmt"
  //"reflect"
)

func sum(slice []int) int {
  total := 0
  
  for _, val := range slice {
    total += val
  }
  
  return total
}

func main() {
  fmt.Println(sum([]int {1, 2, 3}))
}
