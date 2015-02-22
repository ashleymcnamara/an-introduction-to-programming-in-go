package main

import (
    "fmt"
    "os"
)

// this is a comment

/*
 This is another comment
 Multi line, yay!
*/

func OnExit() {
  fmt.Println("On Exit called!")
}

/*
godoc fmt Println
func Println(a ...interface{}) (n int, err error)
    Println formats using the default formats for its operands and writes to
    standard output. Spaces are always added between operands and a newline
    is appended. It returns the number of bytes written and any write error
    encountered.
*/

func main() {
  defer OnExit()  
  
  fmt.Println("Hello world")
  
  fmt.Println(`
            multi 
            line
            string
            yay!
  `)
  
  fmt.Println("Hello, my name is", "Alex")
    
  os.Exit(0)    
}
