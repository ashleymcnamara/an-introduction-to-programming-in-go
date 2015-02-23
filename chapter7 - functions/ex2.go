package main

/*
 Write a function which takes an integer and halves it and returns:
 - true if it was even
 - false if it was odd
 
 For example half(1) should return (0, false) and half(2) should return (1, true).
*/

import "fmt"

func half(i int) (int, bool) {
  x := i / 2
  
  //fmt.Println(i)
  
  if i % 2 == 0 {
    return x, true
  } else {
    return x, false
  }
  
}

func main() {
  x, y := half(1)
  z, w := half(2)
  fmt.Println(x, y, z, w)
}
