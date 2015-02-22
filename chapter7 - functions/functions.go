package main

import "fmt"

// functions can name the return type
func f2() (r int) {
  r = 1
  return
}

func average(xs []float64) float64 {
  //panic("Not implemented")
  total := 0.0
  for _, v := range xs {
    total += v    
  }
  
  return total / float64(len(xs))
}

// function returning multiple values
func f() (int, int) {
  return 5, 6
}

// variadic function
func add(args ...int) int {
  total := 0
  for _, v := range args {
    total += v
  }
  return total
}

// Returns a function which generates even numbers
// Each time it's called it adds 2 to the local i variable
// which unline normal local variables persists between calls
func makeEvenGenerator() func() uint {
  i := uint(0)
  return func() (ret uint) {
    ret = i 
    i += 2
    return
  }
}

func factorial(x uint) uint {
  if x == 0 {
    return 1
  }
  
  return x * factorial(x - 1)
}

func main() {
  someOtherName := []float64 {98, 93, 77, 82, 83}
  fmt.Println(average(someOtherName), f2)  
  x, y := f()
  fmt.Println(x, y)
  xs := []int {1, 2, 3, 4, 5}
  fmt.Println(add(1, 2, 3), add(xs...))
  
  // closure
  addc := func(x, y int) int {
        return x + y
  }

  fmt.Println(addc(1, 1))
  
  u := 0
  increment := func() int {
    u++
    return u
  }
  fmt.Println(increment())
  fmt.Println(increment(), u)
  
  nextEven := makeEvenGenerator()
  fmt.Println(nextEven())
  fmt.Println(nextEven())
  fmt.Println(nextEven())  
  
  fmt.Println(factorial(8))
}