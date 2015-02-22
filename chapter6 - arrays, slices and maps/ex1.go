package main

import "fmt"

/*  How do you access the 4 th element of an array or slice?
*/
func main() {
  array := [4]int { 1, 2, 3, 4}
  
  slice := make([]int, 4)
  slice[3] = 4
  
  fmt.Println(array[3], slice[3])
}
