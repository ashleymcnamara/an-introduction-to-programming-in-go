 package main
 
 import (
  "fmt"
  "os"
 )
 
 func main() {
  file, err := os.Open("test.txt")
  if err != nil {
    return
  }
  
  defer file.Close()
  
  stat, err := file.Stat()
  if err != nil {
    return
  }
  
  fileSize := stat.Size()
  
  fmt.Println("File size:", fileSize, "bytes")
  
  bs := make([]byte, fileSize)
  _, err = file.Read(bs)
  if err != nil {
    return
  }
  
  str := string(bs)
  fmt.Println(str)
 }