package main

/*
   What are defer, panic and recover? How do you recover from a run-time panic?
*/

import "fmt"

func inc(n *uint) {
  defer func() {
    str := recover()
    fmt.Println(str)
  }()
  
  *n++
  if *n == 9 {
    panic("9!!!")
  } 
}

func main() {  
  defer func() {
    str := recover()
    fmt.Println(str)
  }()
  
  var i uint = 7
  inc(&i)
  inc(&i)  
    
  fmt.Println(i)
}