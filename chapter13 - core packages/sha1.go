package main

import (
  "fmt"
  "crypto/sha1"
)

// sha1 computes a 160 bit hash aka a slice of 20 bytes
func main() {
  h := sha1.New()
  h.Write([]byte("test"))
  bs := h.Sum([]byte{})
  fmt.Println(bs)
}