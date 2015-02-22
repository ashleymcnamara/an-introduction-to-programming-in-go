package main

import "fmt"

func main() {
  var x string = "Hello world"
  fmt.Println(x)
  
  var y string
  y = "Hello world"
  fmt.Println(y)
  
  var f string
  f = "first"
  fmt.Println(f)
  
  var s string
  s = "second"
  fmt.Println(s)
  
  f = f + " second"
  fmt.Println(f)
}
