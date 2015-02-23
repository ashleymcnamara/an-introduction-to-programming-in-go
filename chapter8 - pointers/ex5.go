package main

import "fmt"

func swap(x, y *int) {
  tmp := *x
  *x = *y
  *y = tmp
}

func main() {
  x, y := 1, 2
  fmt.Printf("Before swap: x=%d ; y=%d\n", x, y)
  swap(&x, &y)
  fmt.Printf("After swap: x=%d ; y=%d\n", x, y)
}
