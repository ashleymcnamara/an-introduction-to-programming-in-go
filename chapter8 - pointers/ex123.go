package main

import "fmt"

func main() {
  var yPtr *float64 = new (float64)
  *yPtr = 3.14159
  fmt.Println(yPtr, &yPtr, *yPtr, &(*yPtr))
}
