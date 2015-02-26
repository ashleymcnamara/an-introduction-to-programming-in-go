package main

import ("fmt"; "sort")

type Person struct {
  Name string
  Age int
}

type ByName []Person

// The sort.Interface requires 3 methods: Len, Less, Swap
func (this ByName) Len() int {
  return len(this)
}

func (this ByName) Less(i, j int) bool {
  return this[i].Name < this[j].Name
}

func (this ByName) Swap(i, j int) {
  this[i], this[j] = this[j], this[i]
}

type ByAge []Person

func (this ByAge) Len() int {
  return len(this)
}

func (this ByAge) Less(i, j int) bool {
  return this[i].Age < this[j].Age
}

func (this ByAge) Swap(i, j int) {
  this[i], this[j] = this[j], this[i]
}


func main() {
  kids1 := []Person {
    {"Jill", 9},
    {"Jack", 10},
  }
  
  sort.Sort(ByName(kids1))
  fmt.Println(kids1)
  
  kids2 := []Person {
    {"Jill", 9},
    {"Jack", 10},
  }
  
  sort.Sort(ByAge(kids2))
  fmt.Println(kids2)
}