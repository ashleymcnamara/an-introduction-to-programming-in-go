package main

/*
  The Fibonacci sequence is defined as: 
  fib(0) = 0, 
  fib(1) = 1, 
  fib(n) = fib(n-1) + fib(n-2)
  
  Write a recursive function which can find fib(n).
*/

import "fmt"

func Fibonacci(n uint) uint {
  switch n {
    case 0: return 0
    case 1: return 1
    default: return Fibonacci(n - 1) + Fibonacci(n - 2)
  }
}

func main() {
  fmt.Println(Fibonacci(8))
}