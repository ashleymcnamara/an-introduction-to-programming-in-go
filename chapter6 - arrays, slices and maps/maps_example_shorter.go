package main

import "fmt"

func main() {
  elements := map[string]string {
    "H": "Hydrogen",
    "He": "Helium",
    "Li": "Lithium",
    "Be": "Beryllium",
    "B": "Boron",
    "C": "Carbon",
    "N": "Nitrogen",
    "O": "Oxygen",
    "F": "Fluorine",
    "Ne": "Neon",
    }
  
  fmt.Println(elements["Li"], elements["Un"])
  
  name, ok := elements["Un"]
  fmt.Println(name, ok)
  
  if name, ok := elements["un"]; ok {
    fmt.Println(name, ok)
  }
}
