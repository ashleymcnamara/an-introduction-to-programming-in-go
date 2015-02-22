package main

/*
 Using the example program as a starting point,
write a program that converts from Fahrenheit
into Celsius. (C = (F - 32) * 5/9)
*/

import "fmt"

func fahreinheit_to_celsius(fahrenheit float64) float64 {
  return (fahrenheit - 32) * 5.0/9  
}

func celsius_to_fahrenheit(celsius float64) float64 {
  return (9.0/5 * celsius) + 32
}

func main() {
  c := 42.22
  f := 107.996
  fmt.Printf("%f degrees Celsius equals %f degrees Fahrenheit\n", c, celsius_to_fahrenheit(c)) 
  fmt.Printf("%f degrees Fahrenheit equals %f degrees Celsius\n", f, fahreinheit_to_celsius(f)) 
}
