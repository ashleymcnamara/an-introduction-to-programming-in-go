package main

import "fmt"

func zero(x int) {
  x = 0
}

func zerop(xPtr *int) {
  *xPtr = 0
}

func one(xPtr *int) {
  *xPtr = 1
}

func main() {
  x := 5
  zero(x)
  fmt.Println(x) // x is still 5
  
  zerop(&x)
  fmt.Println(x) // x is 0
  
  xPtr := new(int)
  one(xPtr)
  fmt.Println(*xPtr) 
}
