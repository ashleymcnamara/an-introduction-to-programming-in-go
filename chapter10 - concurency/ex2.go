package main

import ("fmt"; "time"; "math/rand")

func Sleep(seconds int32) {
  fmt.Println("Start time", time.Now())
  select {
    case msg := <- time.After(time.Duration(seconds) * time.Second):
      fmt.Println("Sleeped", seconds, "seconds", "\nEnd time", msg)
  }
}

func main() {
    rand.Seed(time.Now().Unix())
    Sleep(rand.Int31n(10))
}
