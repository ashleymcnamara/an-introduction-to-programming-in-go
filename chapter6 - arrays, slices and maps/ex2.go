package main

import "fmt"

/*   What is the length of a slice created using:
make([]int, 3, 9) ?
*/
func main() {    
  slice := make([]int, 3, 9)
  fmt.Println(slice, len(slice))
  
  for i := 4; i <= 9; i++ {
    slice = append(slice, i)
  }
  
  slice = append(slice, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)
  fmt.Println(slice, len(slice))
}
