package main

import (
  "fmt"
  "hash/crc32"
  "io/ioutil"
)

// crc32 computes a 32 bit hash
func getHash(filename string) (uint32, error) {
  bs, err := ioutil.ReadFile(filename)
  if err != nil {
    return 0, err
  }
  
  h := crc32.NewIEEE()
  h.Write(bs)
  
  return h.Sum32(), nil
}

func main() {
  h1, err := getHash("test1.txt")
  if err != nil {
    return
  }
  
  h2, err := getHash("test2.txt")
  if err != nil {
    return
  }
  
  h3, err := getHash("test3.txt")
  if err != nil {
    return
  }
  
  fmt.Println(h1, h2, h3, h1 == h2, h2 == h3)
}
