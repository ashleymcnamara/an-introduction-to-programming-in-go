package main

/*
Write a function with one variadic parameter that finds the greatest number in a list of num- bers.
*/

import "fmt"

func max(list ...int) int {
  if len(list) == 0 {
    panic("Empty list!")
  }
  
  max := list[0]
  
  for _, val := range list {
    if val > max {
      max = val
    }
  }
  
  return max
}

func main() {
  fmt.Println(max([]int{1, 2, 33, 4, 5, 6}...))
}