package main 

import (
  "encoding/gob" // go serialization
  "fmt"
  "net"
)

func server() {
  // listen on a port
  listener, err := net.Listen("tcp", ":9999")
  if err != nil {
    fmt.Println(err)
    return
  }
  
  for {
      // accept a connection      
      connection, err := listener.Accept()
      if err != nil {
        fmt.Println(err)
        continue
      }
      
      // handle the connection
      go handleServerConnection(connection)
  }
}

func handleServerConnection(connection net.Conn) {
  // receive the message
  var msg string
  err := gob.NewDecoder(connection).Decode(&msg)
  
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("Received", msg)
  }
  
  connection.Close()
}

func client() {
  // connect to the server
  connection, err := net.Dial("tcp", "127.0.0.1:9999")
  if err != nil {
    fmt.Println(err)
    return
  }
  
  // send the message
  msg := "Hello World"
  fmt.Println("Sending", msg)
  
  err = gob.NewEncoder(connection).Encode(msg)
  if err != nil {
    fmt.Println(err)
  }
  
  connection.Close()
}

func main() {
  go server()
  go client()
  
  var input string
  fmt.Scanln(&input)
}
