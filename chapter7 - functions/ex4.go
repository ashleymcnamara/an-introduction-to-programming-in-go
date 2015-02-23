package main

/*
 Using makeEvenGenerator as an example, write a makeOddGenerator function that generates odd numbers.
*/

import "fmt"

func makeOddGenerator() func() uint {
  i := uint(1)
  
  return func() (ret uint) {
    ret = i 
    i += 2
    return
  }
}

func main() {
  oddGen := makeOddGenerator()
  
  for i := 100; i >= 0; i-- {
    fmt.Println(oddGen())
  }
}