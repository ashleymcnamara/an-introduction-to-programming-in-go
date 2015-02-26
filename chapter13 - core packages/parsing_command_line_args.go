package main

import ("fmt" ; "flag" ; "math/rand")

/*
[kubuntu@milkyway chapter13 - core packages]$ go run parsing_command_line_args.go -max=102 test 123
35 [test 123]
*/
func main() {
  // Define flags
  maxp := flag.Int("max", 6, "the max value")
  
  // Parse
  flag.Parse()
  
  var args []string
  args = flag.Args()
  
  // Generate a number between 0 and max
  fmt.Println(rand.Intn(*maxp), args)
}