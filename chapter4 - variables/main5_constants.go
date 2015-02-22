package main

import "fmt"

func main() {
  const x string = "Hello world"
  
  // ./main5_constants.go:7: cannot assign to x
  x = "some other string"
  
  fmt.Println(x)
}