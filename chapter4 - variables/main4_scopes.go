package main

import "fmt"

var x string = "Hello world2"

func main() {
  var y string = "Hello world1"
  fmt.Println(x)
}

// undefined: y error
func f() {
  fmt.Println(x)
  fmt.Println(y)
}