package main

import (
  "fmt"
  "time"
  "math/rand"
)

func f(n int) {
  for i := 0; i < 10; i++ {
    fmt.Println(n, ":", i)
    amt := time.Duration(rand.Intn(250))
    time.Sleep(time.Millisecond * amt)
  }
}

func main() {
    go f(0)
   
   for i := 0; i < 100; i++ {
      go f(i)
   }
   
    var input string
    fmt.Scanln(&input)
}
