package main

import "fmt"

func main() {
  var x string = "hello"
  var y string = "world"
  fmt.Println( x == y )
  
  var f string = "hello"
  var s string = "hello"
  fmt.Println( f == s )
  
  // type inference
  w := "hello world"
  q := 5
  fmt.Println(w, q)
}
