package main

import (
  "fmt"
  "time"
)

func pinger(c chan <- string) {
  for i := 0; ; i++ {
    /* 
      Using a channel likes this synchronized the two goroutines
      when pinger attempts to send a message on the channel it will wait until printer is ready to receive the message
      (this is known as blocking)
    */
     c <- "ping"
  }
}

func printer(c <- chan string) {
  for {
    msg := <- c
    fmt.Println(msg)
    time.Sleep(time.Second * 1)
  }
}

func ponger(c chan string) {
  for i := 0; ; i++ {
    c <- "pong"
  }
}

func main() {
  var c chan string = make(chan string)
  
  go pinger(c)
  go ponger(c)
  go printer(c)
  
  var input string
  fmt.Scanln(&input)
  //fmt.Println(<-c)
}
