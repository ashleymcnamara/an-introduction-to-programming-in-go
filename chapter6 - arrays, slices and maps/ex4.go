package main

import "fmt"

/* Write a program that finds the smallest number in this list:
*/
func main() {    
  x := [] int {
    48,96,86,68,
    57,82,63,70,
    37,34,83,27,
    19,97,9,17,
  }

  min := x[0]
  
  for _, val := range x {
      if val < min {
        min  = val
      }
  }
  
  fmt.Printf("Minimum element from slice x is: %d\n", min)
}
