package main

import ( 
  "strings"
  "fmt"
)

func main() {
  fmt.Println(
    // true
    strings.Contains("test", "es"),
    
    // 2
    strings.Count("test", "t"),
    
    // true
    strings.HasPrefix("test", "te"),
    
    // true
    strings.HasSuffix("test", "st"),
    
    // 1
    strings.Index("test", "e"),
    
    // "a-b"
    strings.Join([]string{"a", "b"}, "-"),
    
    // == "aaaa"
    strings.Repeat("a", 5),
    
    // bbaa
    strings.Replace("aaaa", "a", "b", 2),
    
    // []string{"a", "b", "c", "d", "e"}
    strings.Split("a-b-c-d-e", "-"),
    
    // "test"
    strings.ToLower("TEST"),
    
    // "TEST"
    strings.ToUpper("test"),
  )
  
  // convert a string to a slice of bytes
  arr := []byte("test")
  
  // convert a slice of bytes to a string
  str := string([]byte{'t', 'e', 's', 't'})
  
  fmt.Println(arr, str)
}