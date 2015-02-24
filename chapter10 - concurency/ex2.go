package main

import ("fmt"; "time"; "math/rand")

func Sleep(seconds int32) {
  select {
    case <- time.After(time.Duration(seconds) * time.Second):
      fmt.Println("Sleeped", seconds, "seconds")
  }
}

func main() {
    rand.Seed(time.Now().Unix())
    Sleep(rand.Int31n(10))
}
