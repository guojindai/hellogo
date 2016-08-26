package main

import (
  "fmt"
)

func main() {
  fmt.Print("Go runs on ")
  switch i := 1; i {
  case 1:
    fmt.Println(10)
  case f():
    fmt.Println(100)
  default:
    fmt.Printf("default")
  }
}

func f () int {
  return 1
}