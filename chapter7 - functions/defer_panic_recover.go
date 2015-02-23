package main

import "fmt"
import "os"

func first() {
  fmt.Println("1st")
}

func second() {
  fmt.Println("2nd")
}

func test_panic1() {
  panic("PANIC")
  str := recover()
  fmt.Println(str)
}

func test_panic2() {
  defer func() {
    str := recover()
    fmt.Println(str)
  }()
  
  panic("Panic")
}

func main() {  
  defer second()
  second()
  first()
  
  filename := "test.txt"
  f, _ := os.Open(filename)
  defer f.Close()
  
  test_panic2()
  test_panic1()
}